package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

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
	if r.Method == http.MethodGet {
		m.GetManufacturer(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		m.PostManufacturer(rw, r)
		return
	}
	http.Error(rw, "Not implemented", http.StatusNotImplemented)

}

func (m *Manufacturer) PostManufacturer(rw http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer r.Body.Close()
		mm := models.NewManufacturer()
		j := json.NewDecoder(r.Body)
		err := j.Decode(&mm)
		if err != nil {
			m.l.Fatalln(err.Error())
		}
		mm.AddManufacturer(m.db)
		mm.ToJSON(rw)
	}()
	wg.Wait()
}

func (m *Manufacturer) GetManufacturer(rw http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		mm := models.NewManufacturer()
		id, foundid := mux.Vars(r)["id"]
		name, foundname := mux.Vars(r)["name"]
		if foundid {
			err := mm.GetManufacturer("id", id, m.db)
			if err != nil {
				http.Error(rw, fmt.Sprintf("Manufacturer not found: %s", id), http.StatusNotFound)
				return
			}
		} else if foundname {
			err := mm.GetManufacturer("name", name, m.db)
			if err != nil {
				http.Error(rw, fmt.Sprintf("Manufacturer not found: %s", name), http.StatusNotFound)
				return
			}
		} else {
			http.Error(rw, "Bad request", http.StatusBadRequest)
			return
		}
		err := mm.ToJSON(rw)
		if err != nil {
			m.l.Println(err.Error())
		}
	}()
	wg.Wait()
}
func (m *Manufacturer) GetManufacturers(rw http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		j := json.NewEncoder(rw)
		//var res []*models.Manufacturer
		mm := models.NewManufacturer()
		/* 		ml, count := mm.GetManufacturers(m.db)
		   		for i := 0; i < count; i++ {
		   			ml.Scan(&res)
		   			ml.Next()
		   			res = append(res, mm)
		   		} */
		j.Encode(mm.GetManufacturers(m.db))
	}()
	wg.Wait()
}
