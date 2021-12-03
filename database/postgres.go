package database

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/amik3r/neptun-timetable/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection interface {
	Migrate(m models.Model) error
	CreateRecord(m models.Model) error
	Close() error
	Connect() error
}

type DBPostgres struct {
	Con    *gorm.DB
	dbpass string
	dbuser string
	dbhost string
	dbport int
	dbname string
}

func (db *DBPostgres) Migrate(m models.Model) error {
	return db.Con.AutoMigrate(&m)
}

func (db *DBPostgres) CreateRecord(m models.Model) error {
	res := db.Con.Create(m)
	if errors.Is(res.Error, gorm.ErrInvalidData) {
		return errors.New(res.Error.Error())
	}
	return nil
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
	fmt.Println("Connection closed")
	return nil
}

func (d *DBPostgres) Connect() error {
	con, err := createConnection()
	if err != nil {
		return err
	}
	d.Con = con
	fmt.Println("Connection created")
	return nil
}

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		return ""
	}
	return os.Getenv(key)
}

func createConnection() (*gorm.DB, error) {

	port, err := strconv.Atoi(getEnv("DBPORT"))
	if err != nil {
		panic(err)
	}
	db := &DBPostgres{
		dbuser: getEnv("DBUSER"),
		dbpass: getEnv("DBPASS"),
		dbhost: getEnv("DBHOST"),
		dbport: port,
		dbname: getEnv("DBNAME"),
	}

	dsn := url.URL{
		User:     url.UserPassword(db.dbuser, db.dbpass),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", db.dbhost, db.dbport),
		Path:     db.dbname,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	dbCon, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return dbCon, nil
}
