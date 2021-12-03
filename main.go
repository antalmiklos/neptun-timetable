package main

import (
	"fmt"
	"sync"

	"github.com/amik3r/neptun-timetable/database"
	"github.com/amik3r/neptun-timetable/models"
)

func CreateManufacturer(name string, db database.DbConnection, wg *sync.WaitGroup) {
	var manufacturer models.Manufacturer
	m := manufacturer.NewManufacturer()
	m.Name = name
	err := db.Migrate(m)
	if err != nil {
		fmt.Println(err)
	}
	err = db.CreateRecord(m)
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	db := &database.DBPostgres{}
	db.Connect()
	go CreateManufacturer("Orange", db, &wg)
	go CreateManufacturer("Shure", db, &wg)
	go CreateManufacturer("Bose", db, &wg)
	wg.Wait()
}
