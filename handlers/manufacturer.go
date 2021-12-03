package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amik3r/neptun-timetable/database"
	"github.com/amik3r/neptun-timetable/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Manufacturer struct {
	l  log.Logger
	db *gorm.DB
}

func NewManufacturer(l *log.Logger, dc database.DbConnection) *Manufacturer {
	db, err := dc.GetDBInstance()
	if err != nil {
		panic(err)
	}
	return &Manufacturer{*l, db}
}

func (m *Manufacturer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	mm := models.NewManufacturer()
	params := mux.Vars(r)
	id := params["id"]
	m.db.Find(&mm, fmt.Sprintf("id=%s", id))
	err := mm.ToJSON(rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
