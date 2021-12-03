package hardware

import (
	"encoding/json"
	"io"

	"github.com/amik3r/neptun-timetable/models"
	connections "github.com/amik3r/neptun-timetable/models/io"
)

type Amplifier struct {
	ID           uint                `gorm:"primaryKey" json:"id"`
	Manufacturer models.Manufacturer `json:"manufacturer"`
	Model        string              `gorm:"unique" json:"model"`
	Type         string
	Technology   string
	Output       int
	Outputs      []connections.Cable
	Inputs       []connections.Cable
}

func (m *Amplifier) NewManufacturer() *Amplifier {
	return &Amplifier{}
}

func (m *Amplifier) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}
