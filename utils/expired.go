package utils

import "time"

func GetExpiredTime() time.Time {
	now := time.Now()
	expiredTime := now.Add(3 * 24 * time.Hour)
	return expiredTime
}

func GetExpiredTimeGym() time.Time {
	now := time.Now()
	expiredTime := now.Add(30 * 24 * time.Hour)
	return expiredTime
}
