package main

import (
  "log"
  "net/http"
  "github.com/dimfeld/httptreemux"
  "github.com/icaromh/fuzzy-tribble/handlers"
)

func main(){
  addr := "127.0.0.1:8081"
  router := httptreemux.NewContextMux()

  router.Handler(http.MethodGet,  "/movies"     , &handlers.GetMovieHandler{})
  router.Handler(http.MethodGet,  "/movies/:id" , &handlers.GetMovieHandler{})
  router.Handler(http.MethodPost, "/movies"     , &handlers.UpsertMovieHandler{})
  // router.Handler(http.MethodPut,  "/movies/:id" , &handlers.UpsertMovieHandler{})

  log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
