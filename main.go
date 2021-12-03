package main

import (
	"fmt"

	"github.com/amik3r/neptun-timetable/database"
	"github.com/amik3r/neptun-timetable/server"
)

/* func CreateManufacturer(name string, db database.DbConnection, wg *sync.WaitGroup) {
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
} */

func main() {
	db := database.DBPostgres{}
	db.Connect()
	s := server.NewServer(&db)
	err := s.StartServer()
	if err != nil {
		fmt.Println(err.Error())
	}
}
