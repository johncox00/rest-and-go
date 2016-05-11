package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "github.com/dgrijalva/jwt-go"
  "github.com/auth0/go-jwt-middleware"
  "time"
)

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
   w.Write([]byte("API is up and running"))
})

var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  payload, _ := json.Marshal(products)
  w.Header().Set("Content-Type", "applicaiton/json")
  w.Write([]byte(payload))
})

var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  var product Product
  vars := mux.Vars(r)
  slug := vars["slug"]

  for _, p := range products {
    if p.Slug == slug {
        product = p
    }
  }

  w.Header().Set("Content-Type", "application/json")
  if product.Slug != "" {
    payload, _ := json.Marshal(product)
    w.Write([]byte(payload))
  } else {
    http.Error(w, "Product not found.", http.StatusNotFound)
  }
})

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  token := jwt.New(jwt.SigningMethodHS256)
  token.Claims["admin"] = true
  token.Claims["name"] = "Ado Kukic"
  token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
  tokenString, _ := token.SignedString(mySigningKey)
  tokenMap := make(map[string]string)
  tokenMap["token"] = tokenString
  payload, _ := json.Marshal(tokenMap)
  w.Write([]byte(payload))
})

var mySigningKey = []byte("thisISmysigningkey")


// Header format: Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNDYzMDc2NjIxLCJuYW1lIjoiQWRvIEt1a2ljIn0.fMRuj12qQ76r8SSsbomZcV4SeJKvU0TgsuJlv5sF6OI
var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
  ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
    return mySigningKey, nil
  },
  SigningMethod: jwt.SigningMethodHS256,
})
