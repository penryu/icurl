package main


import (
  "fmt"
)

// #include <curl/curl.h>
import "C"

func fetch(url string, filename string) uint64 {
  fmt.Printf("%s => %s\n", url, filename)
  return 0
}

func main() {
  urls := map[string]string{
    "https://penryu.app/": "index.html",
    "https://penryu.app/favicon.ico": "favicon.ico",
  }

  for filename, url := range urls {
    fetch(url, filename)
  }

}

