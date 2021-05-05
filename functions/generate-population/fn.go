package generate-population

import (
  "encoding/json"
  "fmt"
  "html"
  "net/http"
)

// GeneratePopulation generates the population for the genetic internship planner
func GeneratePopulation(w http.ResponseWriter, r *http.Request) {
  var d struct {
    Message string `json:"message"`
  }
  if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
    fmt.Fprint(w, "Trying generate-population algorithm")
    return
  }
  if d.Message == "" {
    fmt.Fprint(w, "Trying generate-population algorithm")
    return
  }
  fmt.Fprint(w, html.EscapeString(d.Message))
}