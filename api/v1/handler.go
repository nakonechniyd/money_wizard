package apiv1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CalcYears -
func CalcYears(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)

	years, err := strconv.Atoi(query["years"])
	if err != nil {
		ErrorJSON(w, "`years` should be set as int")
		return
	}

	var monthlyIncomeAmmount float64
	monthlyIncomeParam := r.FormValue("monthly_income")
	if monthlyIncomeParam != "" {
		monthlyIncomeAmmount, err = strconv.ParseFloat(monthlyIncomeParam, 64)
		if err != nil {
			ErrorJSON(w, "`monthly_income` should be set as float")
			return
		}
	}

	var capital float64
	capitalParam := r.FormValue("capital")
	if capitalParam != "" {
		capital, err = strconv.ParseFloat(capitalParam, 64)
		if err != nil {
			ErrorJSON(w, "`capital` should be set as float")
			return
		}
	}

	if int(monthlyIncomeAmmount) == 0 && int(capital) == 0 {
		ErrorJSON(w, "`monthly_income` and `capital` can't be 0 simultaneously")
		return
	}

	var retireYear int
	retireYearParam := r.FormValue("retire_year")
	if retireYearParam != "" {
		retireYear, err = strconv.Atoi(retireYearParam)
		if err != nil {
			ErrorJSON(w, "`retire_year` should be set as int")
			return
		}
	}

	var lifePercent float64
	lifePercentParam := r.FormValue("life_percent")
	if lifePercentParam != "" {
		lifePercent, err = strconv.ParseFloat(lifePercentParam, 64)
		if err != nil {
			ErrorJSON(w, "`life_percent` should be set as float")
			return
		}
	}

	taxPercent, err := strconv.ParseFloat(query["tax_percent"], 64)
	if err != nil {
		ErrorJSON(w, "`tax_percent` should be set as float")
		return
	}

	passivePercent, err := strconv.ParseFloat(query["passive_percent"], 64)
	if err != nil {
		ErrorJSON(w, "`passive_percent` should be set as float")
		return
	}

	yearsData := make([]YearCalc, 0)
	for year := 1; year <= years; year++ {

		capitalBefore := capital
		monthsCalc := MonthsCalc{}
		isOnRetire := retireYear != 0 && year > retireYear

		for month := 1; month <= 12; month++ {

			passiveIncome := GetPercent(capital/12, passivePercent)
			tax := GetPercent(passiveIncome, taxPercent)
			netPassiveIncome := passiveIncome - tax

			var monthlyReplenishment float64
			var lifePassiveIncome float64
			if isOnRetire {
				// on retirement:
				// 	- no more monthly replenishments
				//  - part of passive income transforms to dividends
				monthlyReplenishment = 0
				lifePassiveIncome = GetPercent(netPassiveIncome, lifePercent)
				netPassiveIncome = netPassiveIncome - lifePassiveIncome
			} else {
				// working hard:
				// 	- apply next monthly replenishment
				//  - no dividends at all
				monthlyReplenishment = monthlyIncomeAmmount
				lifePassiveIncome = 0
			}

			// monthlyReplenishment - in fact is a net monthly income
			newCapital := capital + netPassiveIncome + monthlyReplenishment

			monthsCalc.Months = append(monthsCalc.Months, MonthCalc{
				Year:                 year,
				Month:                month,
				CapitalBefore:        capital,
				MonthlyReplenishment: monthlyReplenishment,
				PassiveIncome:        passiveIncome,
				LifePassiveIncome:    lifePassiveIncome,
				NetPassiveIncome:     netPassiveIncome,
				CapitalAfter:         newCapital,
			})

			capital = newCapital
		}

		yearCalc := YearCalc{
			Year:                  year,
			CapitalBefore:         capitalBefore,
			TotalMonthlyIncome:    monthsCalc.Sum("MonthlyReplenishment"),
			TotalPassiveIncome:    monthsCalc.Sum("PassiveIncome"),
			TotalNetPassiveIncome: monthsCalc.Sum("NetPassiveIncome"),
			CapitalAfter:          capital,
			MonthsCalc:            monthsCalc,
		}
		yearsData = append(yearsData, yearCalc)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	json.NewEncoder(w).Encode(yearsData)
}
