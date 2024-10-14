package models

type UserAdmin struct {
	IdUser       int    `json:"id_user" gorm:"column:id_user;primary_key;autoIncrement"`
	Fullname     string `json:"fullname" gorm:"column:full_name;type:varchar(255);"`
	Email        string `json:"email" gorm:"uniqueIndex"`
	Password     string `json:"password" gorm:"column:password;type:varchar(255);"`
	Role         string `json:"role" gorm:"column:role;type:varchar(20);"`
	JenisKelamin int    `json:"jenis_kelamin" gorm:"column:jenis_kelamin;type:int(2);"`
	Umur         int    `json:"umur" gorm:"column:umur;type:int;"`
	Foto         string `json:"foto" gorm:"column:foto;type:varchar(255);"`
	FotoUrl      string `json:"foto_url" gorm:"column:foto_url;type:varchar(255);"`
	NoTelepon    string `json:"no_telepon" gorm:"column:no_telepon;type:varchar(255);"`
}

func (u *UserAdmin) TableName() string {
	return "user_admin"
}
