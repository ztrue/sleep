package main

import (
  "flag"
  "fmt"
  "log"
  "net/http"
  "regexp"
  "strconv"
  "time"
)

var host = flag.String("host", "localhost", "http service host")
var port = flag.String("port", "8000", "http service port")

var re, _ = regexp.Compile("\\d+")

func main() {
  flag.Parse()

  addr := fmt.Sprintf("%s:%s", *host, *port)

  http.HandleFunc("/", handle)
  log.Fatal(http.ListenAndServe(addr, nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
  timeoutString := re.FindString(r.URL.Path)
  timeout, err := strconv.Atoi(timeoutString)
  if err != nil {
    log.Println(err)
    timeout = 0
  }
  log.Println(timeout)
  time.Sleep(time.Duration(timeout) * time.Second)
  w.Write([]byte(timeoutString))
}
