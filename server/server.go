package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amik3r/neptun-timetable/database"
	"github.com/amik3r/neptun-timetable/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

//type WebModel interface {
//	Get(m WebModel, key string, value string, db *gorm.DB) error
//	GetMany(m WebModel, key string, value string, db *gorm.DB) error
//	Post(m WebModel, db *gorm.DB) error
//	Put(m WebModel, db *gorm.DB) error
//	Patch(m WebModel, db *gorm.DB) error
//	Exists(m WebModel, db *gorm.DB) error
//}

type Server struct {
	l           *log.Logger
	DB          *gorm.DB
	BindAddress string
	Port        int
	//Sm          *mux.Router
	HttpServer http.Server
}

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		return ""
	}
	return os.Getenv(key)
}

func NewServer(dbc database.DbConnection) *Server {
	l := log.New(os.Stdout, "", log.LstdFlags)
	s := Server{BindAddress: getEnv("HOST"), l: l}

	mh := handlers.NewManufacturer(l, dbc)
	ch := handlers.NewCable(l, dbc)

	r := mux.NewRouter()
	r.Use(s.loggingMW)

	mgetr := r.PathPrefix("/manufacturer").Methods(http.MethodGet).Subrouter()
	cgetr := r.PathPrefix("/cable").Methods(http.MethodGet).Subrouter()
	cpostr := r.PathPrefix("/cable").Methods(http.MethodPost).Subrouter()

	// manufacturer get
	mgetr.HandleFunc("/", mh.GetManufacturers)
	mgetr.HandleFunc("/{id:[0-9]+}", mh.GetManufacturer)
	mgetr.HandleFunc("/{name:[a-zA-Z0-9]+}", mh.GetManufacturer)
	// manufacturer post
	mpostr := r.PathPrefix("/manufacturer").Methods(http.MethodPost).Subrouter()
	mpostr.Use(mh.ValidateManufacturerMW)
	mpostr.HandleFunc("/", mh.PostManufacturer)
	// cables post
	cpostr.Use(ch.ValidateCableMW)
	cpostr.HandleFunc("/", ch.PostCable)
	// cables get
	cgetr.HandleFunc("/", ch.GetCables)
	cgetr.HandleFunc("/{id:[0-9]+}", mh.GetManufacturer)
	cgetr.HandleFunc("/{name:[a-zA-Z0-9]+}", mh.GetManufacturer)

	s.HttpServer = http.Server{
		Addr:     s.BindAddress,
		ErrorLog: l,
		Handler:  r,
	}
	return &s
}

func (s *Server) loggingMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		s.l.Println(r.RemoteAddr, r.RequestURI)
		next.ServeHTTP(rw, r)
	})
}

func (s *Server) StartServer() {
	fmt.Printf("Starting server on %s\n", s.BindAddress)
	err := s.HttpServer.ListenAndServe()
	if err != nil {
		s.l.Println(err.Error())
	}
}

/* func getResource(m gorm.Model, db *gorm.DB) {
	defer r.Body.Close()
	j := json.NewEncoder(rw)
	mm := models.NewManufacturer()
	err := j.Encode(mm.GetManufacturers(m.db))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
} */
