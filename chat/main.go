package main

import (
  "log"
  "net/http"
  "sync"
  "path/filepath"
  "text/template"
  "flag"
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
  t.temp1.Execute(w, r)
}

func main() {
  var addr = flag.String("addr", ":8080", "application address")
  flag.Parse()

  r := newRoom()

  http.Handle("/", &templateHandler{filename: "chat.html"})
  http.Handle("/room", r)

  go r.run()
  log.Println("Starting web server. Port: ", *addr)
  if err := http.ListenAndServe(*addr, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
