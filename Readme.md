# Readme

### Available Endpoints
- *POST /users/create*
	```curl
	curl -X POST http://localhost:5000/users/create 
				-d '{"name": "John Doe", "age": 23, "address": "1 Hampton Drive, CA"}'
	```
- *GET /users/list*
	curl -X GET http://localhost:5000/users/list

### Run server
```shell
$ source .env
$ go run main.go users.go
```

### Tests 
```shell
$ go test ./...
```

