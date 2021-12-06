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
	Name string `gorm:"unique" json:"name"`
}

func NewManufacturer() *Manufacturer {
	return &Manufacturer{}
}

func (m *Manufacturer) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func (m *Manufacturer) GetManufacturers(db *gorm.DB) []Manufacturer {
	var results []Manufacturer
	//var resSet []map[string]interface{}
	res := db.Model(&m).Find(&results)
	count := int(res.RowsAffected)
	resr, _ := res.Rows()
	for i := 0; i < count; i++ {
		mm := NewManufacturer()
		res.Scan(mm)
		mm = &Manufacturer{Name: mm.Name, ID: mm.ID}
		results[i] = *mm
		resr.NextResultSet()
	}
	return results
}

func (m *Manufacturer) GetManufacturer(field, value string, db *gorm.DB) error {
	db.Where(fmt.Sprintf("%s=?", field), value).Find(&m)
	if m.ID == 0 {
		return errors.New("record not found")
	}
	return nil
}

func (m *Manufacturer) AddManufacturer(db *gorm.DB) {
	db.Create(&m)
}
