package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/amik3r/neptun-timetable/database"
	"github.com/amik3r/neptun-timetable/models/connector"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Cable struct {
	l  log.Logger
	db *gorm.DB
}

func NewCable(l *log.Logger, dc database.DbConnection) *Cable {
	db, err := dc.GetDBInstance()
	if err != nil {
		panic(err)
	}
	return &Cable{*l, db}
}

func (c *Cable) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.GetCables(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		c.PostCable(rw, r)
		return
	}
	http.Error(rw, "Not implemented", http.StatusNotImplemented)
}

func (c *Cable) GetCables(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	j := json.NewEncoder(rw)
	mm := connector.NewCable()
	err := j.Encode(mm.FindAll(c.db))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Cable) PostCable(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	mm := connector.NewCable()
	j := json.NewDecoder(r.Body)
	err := j.Decode(&mm)
	if err != nil {
		c.l.Println(err.Error())
	}
	fmt.Println(mm)
	err = mm.CreateCable(c.db, mm)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
	} else {
		rw.WriteHeader(http.StatusCreated)
		mm.ToJSON(rw)
	}
}

func (c *Cable) GetCable(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	mm := connector.NewCable()
	id, foundid := mux.Vars(r)["id"]
	name, foundname := mux.Vars(r)["name"]
	if foundid {
		err := mm.Find("id", id, c.db)
		if err != nil {
			http.Error(rw, fmt.Sprintf("Cable not found: %s", id), http.StatusNotFound)
			return
		}
	} else if foundname {
		err := mm.Find("name", name, c.db)
		if err != nil {
			http.Error(rw, fmt.Sprintf("Cable not found: %s", name), http.StatusNotFound)
			return
		}
	} else {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
	err := mm.ToJSON(rw)
	if err != nil {
		c.l.Println(err.Error())
	}
}

func (c *Cable) ValidateCableMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		v := validator.New()
		mm := connector.NewCable()
		buf, _ := ioutil.ReadAll(r.Body)
		rdr1 := io.NopCloser(bytes.NewBuffer(buf))
		rdr2 := io.NopCloser(bytes.NewBuffer(buf))
		r.Body = rdr2
		j := json.NewDecoder(rdr1)
		j.Decode(mm)
		if mm.Size == 0 {
			http.Error(rw, "Size can't be 0", http.StatusBadRequest)
			return
		}
		err := v.Struct(mm)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(rw, r)
	})
}

/* func (c *Cable) Get(mm models.Manufacturer, key string, value string, db *gorm.DB) error {
	db.Where(fmt.Sprintf("LOWER(%s)=LOWER(?)", key), value).Find(&mm)
	if mm.ID == 0 {
		return errors.New("record not found")
	}
	return nil
} */

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
