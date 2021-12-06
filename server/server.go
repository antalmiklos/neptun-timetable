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

	r := mux.NewRouter()
	r.Use(s.loggingMiddleWare)

	s.addHandlers(r)
	mr := r.PathPrefix("/manufacturer").Subrouter()
	mr.PathPrefix("/").HandlerFunc(mh.GetManufacturers).Methods("GET")
	mr.PathPrefix("/{id:[0-9]+}").HandlerFunc(mh.GetManufacturer).Methods("GET")
	mr.PathPrefix("/name/{name:[a-zA-Z0-9]+}").HandlerFunc(mh.GetManufacturer).Methods("GET")
	mr.PathPrefix("/create").HandlerFunc(mh.PostManufacturer).Methods("POST")

	s.HttpServer = http.Server{
		Addr:     s.BindAddress,
		ErrorLog: l,
		Handler:  r,
	}
	return &s
}

func (s *Server) addHandlers(r *mux.Router) {

}

func (s *Server) loggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		s.l.Println(r.RemoteAddr, r.RequestURI)
		next.ServeHTTP(rw, r)
	})
}

func (s *Server) StartServer() {
	fmt.Printf("Starting server on %s\n", s.BindAddress)
	err := s.HttpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
