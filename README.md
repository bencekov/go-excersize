# go-excersize

This is a practice excersize for go projects. The codebase implements a REST API with two endpoints. The server is listening on port 8080.

Sending a POST request with a string in its body to the /message route will return a json containing the same message without vowels. Sending a GET request to the /count route will return a json with the number of messages recieved.

```
curl -X POST -d "aaabcad" localhost:8080/message
'{'message': 'bcd'}'
```

```
curl  localhost:8080/count
'{'count': 1}'
```