package models

type Manufacturer struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

func (m *Manufacturer) NewManufacturer() *Manufacturer {
	return &Manufacturer{}
}

func (m *Manufacturer) ToJSON() (string, error) {
	return "", nil
}
