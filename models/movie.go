package models

type Movie struct {
  Title string `json:"title"`
  Year int `json:"year"`
  Director string `json:"director",emitempty`
  Description string `json:"description",emitempty`
}
