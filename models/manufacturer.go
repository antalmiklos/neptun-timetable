package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"gorm.io/gorm"
)

type Manufacturer struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique" json:"name" validate:"required"`
}

//regex=[a-zA-Z0-9]

func NewManufacturer() *Manufacturer {
	return &Manufacturer{}
}

func (m *Manufacturer) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (m *Manufacturer) GetManufacturers(db *gorm.DB) []Manufacturer {
	var results []Manufacturer
	res := db.Model(&m).Find(&results)
	resr, _ := res.Rows()
	defer resr.Close()
	for resr.Next() {
		resr.Scan(&m)
	}
	return results
}

func (m *Manufacturer) GetManufacturer(field, value string, db *gorm.DB) error {
	db.Where(fmt.Sprintf("LOWER(%s)=LOWER(?)", field), value).Find(&m)
	if m.ID == 0 {
		return errors.New("record not found")
	}
	return nil
}

func (m *Manufacturer) AddManufacturer(db *gorm.DB, mm *Manufacturer) error {
	err := db.Create(mm).Error
	if err != nil {
		return err
	}
	return nil
}
