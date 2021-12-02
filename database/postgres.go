package database

import (
	"fmt"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection interface {
	Connect() error
	Close() error
}

type DBPostgres struct {
	Con    *gorm.DB
	Dbpass string
	Dbuser string
	Dbhost string
	Dbport int
	Dbname string
}

func (d *DBPostgres) Close() error {
	sqlDB, err := d.Con.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		return err
	}
	return nil
}

func (d *DBPostgres) Connect() error {
	con, err := createConnection(d.Dbuser, d.Dbpass, d.Dbport, d.Dbhost, d.Dbname)
	if err != nil {
		return err
	}
	d.Con = con
	return nil
}

func createConnection(dbuser string, dbpass string, dbport int, dbhost string, dbname string) (*gorm.DB, error) {
	dsn := url.URL{
		User:     url.UserPassword(dbuser, dbpass),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", dbhost, dbport),
		Path:     dbname,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
