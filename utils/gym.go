package utils

type GymRequest struct {
	NamaGym    string  `json:"namaGym"`
	AlamatGym  string  `json:"alamatGym"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	LinkGoogle string  `json:"linkGoogle"`
}
