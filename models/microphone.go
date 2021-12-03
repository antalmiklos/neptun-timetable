package models

type Microphone struct {
	ID          uint `gorm:"primaryKey"`
	Manufaturer Manufacturer
	Type        string
	Impedance   int
	SPL         int
	Direction   string
	FreqMin     int
	FreqMax     int
	Weigth      float32
	Width       float32
	Heigth      float32
}

func (m *Microphone) NewMicrophone() *Microphone {
	return &Microphone{}
}

func (m *Microphone) ToJSON() (string, error) {
	return "", nil
}
