package main

import (
  "log"
  "net/http"
  "sync"
  "path/filepath"
  "text/template"
)

type templateHandler struct {
  once sync.Once
  filename string
  temp1 *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  t.once.Do(func () {
    t.temp1 =
      template.Must(template.ParseFiles(filepath.Join("template", t.filename)))
  })
  t.temp1.Execute(w, nil)
}

func main() {
  // root
  http.Handle("/", &templateHandler{filename: "chat.html"})
  // Start web server
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
