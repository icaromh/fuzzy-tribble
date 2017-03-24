package handlers

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/dimfeld/httptreemux"
  "github.com/icaromh/fuzzy-tribble/models"
)

type UpsertMovieHandler struct{

}

func (h *UpsertMovieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  movie := &models.Movie{}
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(movie)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  // TODO: insert in db
  fmt.Fprintf(w, "received: %#v \n", movie)

}

type GetMovieHandler struct{}

func (h *GetMovieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
  if params["id"] != "" {
    fmt.Fprintf(w, "Get id: %s!", params["id"])
  } else {
    fmt.Fprintf(w, "Get all")
  }
}
