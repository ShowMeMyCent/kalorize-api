package utils

type FranchiseRequest struct {
	NamaFranchise      string  `json:"namaFranchise"`
	AlamatFranchise    string  `json:"alamatFranchise"`
	LongitudeFranchise float64 `json:"longitudeFranchise"`
	LatitudeFranchise  float64 `json:"latitudeFranchise"`
	EmailFranchise     string  `json:"emailFranchise"`
	PasswordFranchise  string  `json:"passwordFranchise"`
	NoTeleponFranchise string  `json:"noTeleponFranchise"`
	FotoFranchise      string  `json:"fotoFranchise"`
	LokasiFranchise    string  `json:"lokasiFranchise"`
}
