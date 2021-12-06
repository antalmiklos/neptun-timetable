package hardware

import (
	"encoding/json"
	"io"

	"github.com/amik3r/neptun-timetable/models"
	"github.com/amik3r/neptun-timetable/models/connector"
)

type Amplifier struct {
	ID           uint                `gorm:"primaryKey" json:"id"`
	Manufacturer models.Manufacturer `json:"manufacturer"`
	Model        string              `gorm:"unique" json:"model"`
	Type         string
	Technology   string
	Output       int
	Outputs      []connector.Cable
	Inputs       []connector.Cable
}

func (m *Amplifier) NewManufacturer() *Amplifier {
	return &Amplifier{}
}

func (m *Amplifier) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}
