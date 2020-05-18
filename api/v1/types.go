package apiv1

import "reflect"

// MonthCalc -
type MonthCalc struct {
	Year                 int     `json:"year"`
	Month                int     `json:"month"`
	CapitalBefore        float64 `json:"capital_before,omitempty"`
	MonthlyReplenishment float64 `json:"monthly_replenishment,omitempty"`
	PassiveIncome        float64 `json:"passive_income"`
	LifePassiveIncome    float64 `json:"life_passive_income,omitempty"`
	NetPassiveIncome     float64 `json:"net_passive_income"`
	CapitalAfter         float64 `json:"capital_after,omitempty"`
}

// MonthsCalc -
type MonthsCalc struct {
	Months []MonthCalc `json:"months"`
}

// Sum -
func (mc *MonthsCalc) Sum(field string) float64 {
	total := 0.0
	for _, monthCalc := range mc.Months {
		total += reflect.ValueOf(monthCalc).FieldByName(field).Float()
	}
	return total
}

// YearCalc -
type YearCalc struct {
	Year                  int     `json:"year"`
	CapitalBefore         float64 `json:"capital_before,omitempty"`
	TotalMonthlyIncome    float64 `json:"total_monthly_income,omitempty"`
	TotalPassiveIncome    float64 `json:"total_passive_income"`
	TotalNetPassiveIncome float64 `json:"total_net_passive_income"`
	CapitalAfter          float64 `json:"capital_after,omitempty"`

	// promoted field
	MonthsCalc
}
