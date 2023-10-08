run:
	@go run main.go csv-etl business-financial-data-mar-2022-quarter-csv.csv

fast-build:
	go build -gcflags "all=-N -l" -o ./bin/main ./main.go
	docker build -f Dockerfile.fast -t simple-csv-etl:latest .
