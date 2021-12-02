package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/amik3r/neptun-timetable/database"
	"github.com/amik3r/neptun-timetable/models"
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		return ""
	}
	return os.Getenv(key)
}

func main() {

	port, err := strconv.Atoi(GetEnv("DBPORT"))
	if err != nil {
		panic(err)
	}
	db := &database.DBPostgres{
		Dbuser: GetEnv("DBUSER"),
		Dbpass: GetEnv("DBPASS"),
		Dbhost: GetEnv("DBHOST"),
		Dbport: port,
		Dbname: GetEnv("DBNAME"),
	}

	err = db.Connect()
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	var course models.Course

	res := db.Con.First(&course)
	fmt.Println(*res)
	if err != nil {
		panic(err)
	}
}
