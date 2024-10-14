package utils

import "time"

type HistoryRequest struct {
	IdBreakfast   int
	IdLunch       int
	IdDinner      int
	TotalProtein  int
	TotalKalori   int
	TanggalDibuat time.Time
}
