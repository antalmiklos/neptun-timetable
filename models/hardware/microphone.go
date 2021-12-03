package hardware

import (
	"encoding/json"
	"io"

	"github.com/amik3r/neptun-timetable/models"
	connections "github.com/amik3r/neptun-timetable/models/io"
)

type Microphone struct {
	ID          uint `gorm:"primaryKey"`
	Manufaturer models.Manufacturer
	Type        string
	Impedance   int
	SPL         int
	Direction   string
	FreqMin     int
	FreqMax     int
	Weigth      float32
	Width       float32
	Heigth      float32
	Inputs      []connections.Cable
	Outputs     []connections.Cable
}

func (m *Microphone) NewMicrophone() *Microphone {
	return &Microphone{}
}

func (m *Microphone) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}
