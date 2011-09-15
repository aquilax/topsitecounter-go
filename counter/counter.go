package counter

import (
  "log"
  "fmt"
  "http"
//  "time"
  "strconv"
  "strings"
  "appengine"
  "appengine/datastore"
  "appengine/memcache"
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

func impression(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  referer := r.Referer
  url, err := http.ParseURL(referer);
  if (err == nil) {
    ip := r.RemoteAddr
    save(&c, url.Host, ip)
  }
  header := w.Header()
  header.Set("Content-Type", "image/gif")
  // 1x1 px transparent gif
  dec := []byte {71, 73, 70, 56, 57, 97, 1, 0, 1, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 33, 249, 4, 1, 0, 0, 0, 0, 44, 0, 0, 0, 0, 1, 0, 1, 0, 0, 2, 2, 68, 1, 0, 59}
  w.Write(dec)
}

func save(c *appengine.Context, referer, ip string) {
  log.Println(referer)
  if referer == "" {
    return
  }
  if strings.HasPrefix(referer, "www.") {
    referer = referer[4:]
  }
  siteId := searchsite(c, referer)
  log.Println(siteId)
}

func searchsite(c *appengine.Context, referer string) (int64) {
  hash := "ref:" + referer
  item, err := memcache.Get(*c, hash);
  if (err != nil) {
    q := datastore.NewQuery("MySite").Filter("url", referer)
    log.Print(q);
/*    t := q.Run(*c)
/*    var x MySite
    key, err := t.Next(&x)
    if err == datastore.Done {
      //Not found
      mysite := MySite {
        url: referer,
        created: datastore.SecondsToTime(time.Seconds()),
        views : 1,
        last_access: datastore.SecondsToTime(time.Seconds()),
      }
      key, _ := datastore.Put(*c, datastore.NewIncompleteKey("MySite"), &mysite)
      return key.IntID()
    }
    return key.IntID()
*/
    return 0;
  }
  r, _ := strconv.Atoi64(string(item.Value))
  return r
}
