package goodbye

import (
  "encoding/json"
  "fmt"
  "html"
  "net/http"
)

// GoodBye prints the JSON encoded "message" field in the body
// of the request or "Goodbye!" if there isn't one.
func GoodBye(w http.ResponseWriter, r *http.Request) {
  var d struct {
    Message string `json:"message"`
  }
  if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
    fmt.Fprint(w, "Goodbye!")
    return
  }
  if d.Message == "" {
    fmt.Fprint(w, "Goodbye!")
    return
  }
  fmt.Fprint(w, html.EscapeString(d.Message))
}