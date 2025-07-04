app := ./api_server
test_folder := ./tests/
main := ./main.go
main_folder := ./
DB := ./db/urls.db

.PHONY:

all: format vet tidy build run_and_test clean

format:
	gofmt $(main_folder)

vet:
	go vet $(main_folder)

tidy:
	go mod tidy

build: 
	go build -o $(app) $(main)

run_and_test: build
	./$(app) & \
	PID=$$!; \
	sleep 5; \
	go test -count=1 $(test_folder); \
	kill $$PID

clean:
	rm -f $(app)

clean_db:
	rm -f $(DB) 