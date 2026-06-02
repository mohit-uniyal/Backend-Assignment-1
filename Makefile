migration:
	migrate create -ext sql -dir src/migrations -seq ${NAME}

run:
	go run src/cmd/main.go

compose-up:
	sudo DB_USER=${USERNAME} DB_PASSWORD=${PASSWORD} DB_NAME=${DBNAME} docker compose up --build -d

compose-down:
	sudo docker compose down

compose-up-db:
	sudo DB_USER=${USERNAME} DB_PASSWORD=${PASSWORD} DB_NAME=${DBNAME} docker compose up --build -d postgres