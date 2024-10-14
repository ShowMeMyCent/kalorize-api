package formatter

import (
	"kalorize-api/app/models"
	"kalorize-api/utils"
)

type MakananFormat struct {
	ID          int
	Nama        string
	Jenis       string
	Bahan       []string
	CookingStep []string
	Kalori      int
	Protein     int
	Foto        string
}

func FormatterMakananIndo(makanan models.Makanan) MakananFormat {
	var makananFormatted MakananFormat
	makananFormatted.ID = makanan.IdMakanan
	makananFormatted.Nama = makanan.Nama
	makanan.Bahan = utils.CleanAngleBracketsinString(makanan.Bahan)
	makanan.Bahan = utils.CleanSingleQuoteinString(makanan.Bahan)
	makananFormatted.Bahan = utils.ConvertToArrayWithCommaSeparator(makanan.Bahan)
	makananFormatted.Bahan = utils.AddNumbering(makananFormatted.Bahan)
	makanan.CookingStep = utils.CleanAngleBracketsinString(makanan.CookingStep)
	makanan.CookingStep = utils.CleanSingleQuoteinString(makanan.CookingStep)
	makananFormatted.CookingStep = utils.ConvertToArrayWithCommaSeparator(makanan.CookingStep)
	makananFormatted.CookingStep = utils.AddNumbering(makananFormatted.CookingStep)
	makananFormatted.Kalori = makanan.Kalori
	makananFormatted.Protein = makanan.Protein
	makananFormatted.Foto = makanan.Foto
	return makananFormatted
}
