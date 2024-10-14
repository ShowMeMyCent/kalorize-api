package models

type Gym struct {
	IdGym      int        `json:"id_gym" gorm:"column:id_gym;primaryKey;autoIncrement"`
	NamaGym    string     `json:"nama" gorm:"column:nama;type:varchar(255);"`
	AlamatGym  string     `json:"alamat" gorm:"column:alamat;type:varchar(255);"`
	Latitude   float64    `json:"latitude" gorm:"column:latitude;type:double;"`
	Longitude  float64    `json:"longitude" gorm:"column:longitude;type:double;"`
	LinkGoogle string     `json:"link_google" gorm:"column:link_google;type:varchar(255);"`
	PhotoGym   string     `json:"photo_gym" gorm:"column:photo_gym;type:varchar(255);"`
	PhotoUrl   string     `json:"photo_url" gorm:"column:photo_url;type:varchar(255);"`
	UsedCodes  []UsedCode `gorm:"foreignKey:IdGym"`
}

func (Gym) TableName() string {
	return "gyms"
}
