MAIN_FILE=main.go

run:
	go run $(MAIN_FILE)

build:
	go build -o money_wizard $(MAIN_FILE)

curl:
	curl "http://localhost:8080/api/v1/calc_years/12?tax_percent=19.5&passive_percent=14.0&capital=10000&monthly_replenishment=5000&retire_year=10&life_percent=45.0" | jq
