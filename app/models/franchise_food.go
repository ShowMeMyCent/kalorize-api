package models

type FranchiseMakanan struct {
	IdFranchiseMakanan int `json:"id_franchise_makanan" gorm:"column:id_franchise_makanan;primary_key;autoIncrement"`
	IdFranchise        int `json:"id_franchise" gorm:"column:id_franchise;type:char(36);"`
	IdMakanan          int `json:"id_makanan" gorm:"column:id_makanan;type:char(36);"`
}

func (FranchiseMakanan) TableName() string {
	return "franchise_makanans"
}
