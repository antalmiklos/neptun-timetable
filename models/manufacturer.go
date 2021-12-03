package models

import (
	"encoding/json"
	"io"
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
