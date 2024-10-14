package formatter

import (
	"kalorize-api/app/models"
	"strconv"
)

func FormatterMakananToMultiDimentionalArray(makanan []models.Makanan) [][]string {
	var header = []string{"id", "Nama", "Jenis", "Foto", "Bahan", "Cooking Step", "Kalori", "Protein"}

	var result [][]string
	result = append(result, header)

	for i := range makanan {
		var row []string
		row = append(row, strconv.Itoa(makanan[i].IdMakanan))
		row = append(row, makanan[i].Nama)
		row = append(row, makanan[i].Foto)
		row = append(row, makanan[i].Bahan)
		row = append(row, makanan[i].CookingStep)
		row = append(row, intToString(makanan[i].Kalori))
		row = append(row, intToString(makanan[i].Protein))

		result = append(result, row)
	}
	return result
}

func intToString(num int) string {
	return strconv.Itoa(num)
}
