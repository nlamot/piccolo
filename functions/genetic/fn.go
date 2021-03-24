package genetic

import (
  "encoding/json"
  "fmt"
  "html"
  "net/http"
)

// Generate prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func Generate(w http.ResponseWriter, r *http.Request) {
  var d struct {
    Message string `json:"message"`
  }
  if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
    fmt.Fprint(w, "Trying genetic algorithm")
    return
  }
  if d.Message == "" {
    fmt.Fprint(w, "Trying genetic algorithm")
    return
  }
  fmt.Fprint(w, html.EscapeString(d.Message))
}