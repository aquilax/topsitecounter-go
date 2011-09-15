package counter

import (
  "fmt"
  "http"
  "strings"
  "encoding/base64"
  "appengine/datastore"
)

type MySite struct{
  url string
  created datastore.Time
  views int64
  last_access datastore.Time
}

func init() {
      http.HandleFunc("/", handler)
      http.HandleFunc("/t", impression)
}

func handler(w http.ResponseWriter, r *http.Request) {
      fmt.Fprint(w, "Hello, world!")
}

func Decode(decBuf, enc []byte, e64 *base64.Encoding) []byte {
  maxDecLen := e64.DecodedLen(len(enc))
  if decBuf == nil || len(decBuf) < maxDecLen {
          decBuf = make([]byte, maxDecLen)
  }
  n, _ := e64.Decode(decBuf, enc)
  return decBuf[0:n]
}

func impression(w http.ResponseWriter, r *http.Request) {
  referer := r.Header.Get("Referer")
  ip := r.RemoteAddr
  save(referer, ip)
  header := w.Header()
  header.Set("Content-Type", "image/gif")
  trans_gif_64 := "R0lGODlhAQABAIAAAAAAAAAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw=="

  dec := Decode(nil, []byte(trans_gif_64), base64.StdEncoding)

  w.Write(dec)
}

func save(referer, ip string){
  if referer == "" {
    return
  }
  if strings.HasPrefix(referer, "www.") {
    referer = referer[4:]
  }
}
