# Readme

### Heroku URL
https://mysterious-stream-97316.herokuapp.com

### Available Endpoints
- *POST /users/create*
	```curl
	curl -X POST https://mysterious-stream-97316.herokuapp.com/users/create 
				-d '{"name": "John Doe", "age": 23, "address": "1 Hampton Drive, CA"}'
	```
- *GET /users/list*
	curl -X GET https://mysterious-stream-97316.herokuapp.com/users/list

### Run server
```shell
$ go mod tidy && go mod vendor
$ source .env
$ go run main.go users.go
```

### Tests 
```shell
$ go test ./...
```

