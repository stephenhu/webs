package main

import (
  "flag"
  "fmt"
  "net/http"

  "github.com/fatih/color"
)


var (
  
  addr string
  port string

)


func main() {

  flag.StringVar(&addr, "a", APP_DEFAULT_ADDR,
    "webs bind address, default: 0.0.0.0")

  flag.StringVar(&port, "p", APP_DEFAULT_PORT,
    "webs port, default: 8000")

  flag.Parse()

  a := fmt.Sprintf(":%s", port)

  color.Cyan(fmt.Sprintf("Starting %s", version()))
  color.Green(fmt.Sprintf("Listening on port %s\n", port))

  fmt.Println(http.ListenAndServe(a, http.FileServer(http.Dir(
    APP_PWD))))

  
} // main

