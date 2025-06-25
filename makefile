app := /bin/api_server
test_folder := ./tests/
main := ./main.go

.PHONY:

run: 
	go run $(main)