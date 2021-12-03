package connections

import (
	"encoding/json"
	"io"
)

type Cable struct {
	ID             uint `gorm:"primaryKey"`
	ConnectionName string
	Size           string
	Channels       int
}

func (m *Cable) NewIO() *Cable {
	return &Cable{}
}

func (m *Cable) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}
