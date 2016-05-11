package main

import (
    "os"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

func main() {

  r := mux.NewRouter()

  // On the default page we will simply serve our static index page.
  r.Handle("/", http.FileServer(http.Dir("./views/")))
  // We will setup our server so we can serve static assest like images, css from the /static/{file} route
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

  r.Handle("/status", StatusHandler).Methods("GET")
  r.Handle("/products", jwtMiddleware.Handler(ProductsHandler)).Methods("GET")
  r.Handle("/products/{slug}/feedback", jwtMiddleware.Handler(AddFeedbackHandler)).Methods("POST")

  r.Handle("/tokens", GetTokenHandler).Methods("GET")


  // Our application will run on port 3000. Here we declare the port and pass in our router.
  http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Not Implemented"))
})
