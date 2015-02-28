package main

import "fmt"
import "log"
import "net/http"
import "github.com/gorilla/mux"

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  //r.HandleFunc("/products", ProductsHandler)
  //r.HandleFunc("/articles", ArticlesHandler)
  http.Handle("/", r)

  log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(w, "Hello, World")
}
