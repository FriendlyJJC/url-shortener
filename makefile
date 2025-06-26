app := ./api_server
test_folder := ./tests/
main := ./main.go

.PHONY:

build: 
	go build -o $(app) $(main)

run_and_test:
	./$(app) & \
	PID=$$!; \
	sleep 5; \
	go test -count=1 $(test_folder); \
	kill $$PID

clean:
	rm -f $(app)