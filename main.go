package main

import (
  "html/template"
  "log"
  "net/http"
  "os"
  "path"
  "strings"
)

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
func ShiftPath(p string) (head, tail string) {
  p = path.Clean("/" + p)
  i := strings.Index(p[1:], "/") + 1
  if i <= 0 {
    return p[1:], "/"
  }
  return p[1:i], p[i:]
}

// Router is the router responsible for serving "/"
type Router struct {
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  var head string
  head, req.URL.Path = ShiftPath(req.URL.Path)

  switch head {
  case "":
    templ, _ := template.ParseFiles("templates/index.html")
    templ.Execute(res, "")
  }
}

func main() {
  port := os.Getenv("PORT")

  if port == "" {
    log.Fatal("$PORT must be set")
  }

  router := &Router{}

  http.ListenAndServe(":3000", router)
}