package connector

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"gorm.io/gorm"
)

type Cable struct {
	ID             uint    `gorm:"primaryKey" json:"id"`
	ConnectionName string  `validate:"required" json:"name"`
	Size           float32 `validate:"required" json:"size"`
	Channels       int     `validate:"required" json:"channels"`
}

func NewCable() *Cable {
	return &Cable{}
}

func (c *Cable) FindAll(db *gorm.DB) []Cable {
	var results []Cable
	res := db.Model(&c).Find(&results)
	resr, _ := res.Rows()
	defer resr.Close()
	for resr.Next() {
		resr.Scan(&c)
	}
	return results
}
func (c *Cable) Find(field, value string, db *gorm.DB) error {
	db.Where(fmt.Sprintf("%s=?", field), value).Find(&c)
	if c.ID == 0 {
		return errors.New("record not found")
	}
	return nil
}
func (c *Cable) CreateCable(db *gorm.DB, cc *Cable) error {
	err := db.Create(cc).Error
	if err != nil {
		return err
	}
	return nil
}

/* func (c *Cable) GetCables(db *gorm.DB) []Cable {
	var results []Cable
	res := db.Model(&c).Find(&results)
	resr, _ := res.Rows()
	defer resr.Close()
	for resr.Next() {
		resr.Scan(&c)
	}
	return results
} */

func (m *Cable) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}
