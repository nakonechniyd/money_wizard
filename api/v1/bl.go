package apiv1

// GetPercent -
func GetPercent(ammount float64, percent float64) float64 {
	return ammount * (percent / 100)
}
