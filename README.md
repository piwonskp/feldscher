## Run

Use following commands:
```
docker-compose up
```
Migrate and run web worker:
```
docker-compose exec web bash
createSchema
go run main.go
```
Run fetcher worker:
```
docker-compose exec worker bash
go run cmd/worker/main.go
```
