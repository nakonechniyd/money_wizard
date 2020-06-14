MAIN_FILE=main.go
OUT_NAME=money_wizard

run:
	go run $(MAIN_FILE)

build:
	go build -o $(OUT_NAME) $(MAIN_FILE)

docker-build:
	docker build -t $(OUT_NAME) -f Dockerfile .

docker-run:
	docker container run --rm --publish 8080:8080 $(OUT_NAME):latest

curl:
	curl "http://localhost:8080/api/v1/calc_years/12?tax_percent=19.5&passive_percent=14.0&capital=10000&monthly_replenishment=5000&retire_year=10&life_percent=45.0" | jq
