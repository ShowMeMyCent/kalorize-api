package models

type TargetKalori struct {
}

type User struct {
	IdUser       int        `json:"id_user" gorm:"column:id_user;primary_key;autoIncrement"`
	Fullname     string     `json:"fullname" gorm:"column:full_name;type:varchar(255);"`
	Email        string     `json:"email" gorm:"uniqueIndex"`
	Password     string     `json:"password" gorm:"column:password;type:varchar(255);"`
	Role         string     `json:"role" gorm:"column:role;type:varchar(20);"`
	JenisKelamin int        `json:"jenis_kelamin" gorm:"column:jenis_kelamin;type:int(2);"`
	Umur         int        `json:"umur" gorm:"column:umur;type:int;"`
	BeratBadan   int        `json:"berat_badan" gorm:"column:berat_badan;type:int;"`
	TinggiBadan  int        `json:"tinggi_badan" gorm:"column:tinggi_badan;type:int;"`
	FrekuensiGym int        `json:"frekuensi_gym" gorm:"column:frekuensi_gym;type:int;"`
	TargetKalori int        `json:"target_kalori" gorm:"column:target_kalori;type:int;"`
	ReferalCode  string     `json:"referal_code" gorm:"column:referal_code;type:varchar(255);"`
	Foto         string     `json:"foto" gorm:"column:foto;type:varchar(255);"`
	FotoUrl      string     `json:"foto_url" gorm:"column:foto_url;type:varchar(255);"`
	NoTelepon    string     `json:"no_telepon" gorm:"column:no_telepon;type:varchar(255);"`
	UsedCodes    []UsedCode `json:"user_code" gorm:"foreignKey:IdUser"`
	Histories    []History  `json:"history" gorm:"foreignKey:IdUser"`
}

func (u *User) TableName() string {
	return "users"
}
