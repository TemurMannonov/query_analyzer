## API for getting postgres queries info

Swagger documentation (http://localhost:8000/swagger/index.html)

### Init database
Create .env file or add environment variables
 - HTTP_PORT=:8000
 - POSTGRES_HOST=localhost
 - POSTGRES_PORT=5432
 - POSTGRES_USER=postgres
 - POSTGRES_PASSWORD=postgres
 - POSTGRES_DATABASE=postgres

Run these commands before running application:
```bash
docker compose up -d
make migrateup
```

### Usage
Run application
```bash
go run cmd/main.go 
```

Run test
```bash
make test
```
