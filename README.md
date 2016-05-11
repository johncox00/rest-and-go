# REST & Go

This is a sample REST API build int Go with authentication from auth0.

```
$ go build
$ ./rest
```

Then open up POSTMAN and GET `http://localhost:3000/tokens`. Copy the token that is returned. Now try to GET `http://localhost:3000/products`. Your headers should be:

```
Content-Type: application/json
Authorization: bearer {the token you copied}
```
