build:
		docker-compose build selling
run:
		docker-compose up selling
migrate:	
		migrate -path ./schema -database 'postgres://postgres:54321@localhost:5436/postgres?sslmode=disable' up
migrate-down:	
		migrate -path ./schema -database 'postgres://postgres:54321@localhost:5436/postgres?sslmode=disable' down