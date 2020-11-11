package handler

import (
  "fmt"
  "time"
  "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  currentTime := time.Now()
  fmt.Fprintf(w, currentTime.String())
}
