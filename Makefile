migration:
	migrate create -ext sql -dir src/migrations -seq ${NAME}

run:
	go run src/cmd/main.go