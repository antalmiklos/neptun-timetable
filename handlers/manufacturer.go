package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/amik3r/neptun-timetable/database"
	"github.com/amik3r/neptun-timetable/models"
	"github.com/go-playground/validator"
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
	defer r.Body.Close()
	mm := models.NewManufacturer()
	j := json.NewDecoder(r.Body)
	err := j.Decode(&mm)
	if err != nil {
		m.l.Println(err.Error())
	}
	fmt.Println(mm)
	err = mm.AddManufacturer(m.db, mm)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
	} else {
		rw.WriteHeader(http.StatusCreated)
		mm.ToJSON(rw)
	}
}

func (m *Manufacturer) GetManufacturer(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
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
}

func (m *Manufacturer) GetManufacturers(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	j := json.NewEncoder(rw)
	mm := models.NewManufacturer()
	err := j.Encode(mm.GetManufacturers(m.db))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func (m *Manufacturer) ValidateManufacturerMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		v := validator.New()
		mm := models.NewManufacturer()
		buf, _ := ioutil.ReadAll(r.Body)
		rdr1 := io.NopCloser(bytes.NewBuffer(buf))
		rdr2 := io.NopCloser(bytes.NewBuffer(buf))
		r.Body = rdr2
		j := json.NewDecoder(rdr1)
		j.Decode(mm)
		err := v.Struct(mm)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(rw, r)
	})
}

func (m *Manufacturer) Get(mm models.Manufacturer, key string, value string, db *gorm.DB) error {
	db.Where(fmt.Sprintf("LOWER(%s)=LOWER(?)", key), value).Find(&mm)
	if mm.ID == 0 {
		return errors.New("record not found")
	}
	return nil
}

/*
func (m *Manufacturer) GetMany(zm server.WebModel, key string, value string, db *gorm.DB) error {
	return nil
}
func (m *Manufacturer) Post(zm server.WebModel, db *gorm.DB) error {
	return nil
}
func (m *Manufacturer) Put(zm server.WebModel, db *gorm.DB) error {
	return nil
}
func (m *Manufacturer) Patch(wm server.WebModel, db *gorm.DB) error {
	return nil
}
func (m *Manufacturer) Exists(wm server.WebModel, db *gorm.DB) error {
	return nil
} */
