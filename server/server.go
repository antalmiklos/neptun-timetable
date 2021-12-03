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
	s := Server{}

	l := log.New(os.Stdout, "", log.LstdFlags)
	mh := handlers.NewManufacturer(l, dbc)

	r := mux.NewRouter()
	r.Handle("/manufacturer/{id:[0-9]+}", mh)
	r.Use(loggingMiddleWare)

	s.HttpServer = http.Server{
		Addr:     getEnv("HOST"),
		ErrorLog: l,
		Handler:  r,
	}
	return &s
}

func loggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("lol")
		next.ServeHTTP(rw, r)
	})
}

func (s *Server) StartServer() error {
	return s.HttpServer.ListenAndServe()
}
