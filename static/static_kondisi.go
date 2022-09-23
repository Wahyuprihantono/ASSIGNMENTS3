package static

func KondisiWater(waterVal int) string {
	var waterStatus string
	switch {
	case waterVal <= 5:
		waterStatus = "aman"
	case waterVal >= 6 && waterVal <= 8:
		waterStatus = "siaga"
	default:
		waterStatus = "bahaya"
	}
	return waterStatus
}

func KondisiWind(windVal int) string {
	var windStatus string
	switch {
	case windVal <= 6:
		windStatus = "aman"
	case windVal >= 7 && windVal <= 15:
		windStatus = "siaga"
	default:
		windStatus = "bahaya"
	}
	return windStatus
}
