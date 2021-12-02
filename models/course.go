package models

import (
	"fmt"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	CourseCode  string `gorm:"unique"`
	SubjectCode string
	Start       string
}

func CreateConnection(dbuser string, dbpass string, dbport int, dbhost string, dbname string) (*gorm.DB, error) {
	con, err := connect(dbuser, dbpass, dbport, dbhost, dbname)
	if err != nil {
		return nil, err
	}
	return con, nil
}

func connect(dbuser string, dbpass string, dbport int, dbhost string, dbname string) (*gorm.DB, error) {
	dsn := url.URL{
		User:     url.UserPassword(dbuser, dbpass),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", dbhost, dbport),
		Path:     dbname,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})

	db.AutoMigrate(&Course{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
