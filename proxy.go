package main 
import (
    "log"
    "net/url"
    "net/http"
    "net/http/httputil"
    ) 
func NewSingleHostReverseProxy(url *url.URL) *httputil.ReverseProxy {
  rp := httputil.NewSingleHostReverseProxy(url)
  oldDirector := rp.Director
  rp.Director = func(r *http.Request) {
    oldDirector(r)
    r.Host = url.Host
  }
  return rp
}

func main() {
  log.SetPrefix("代理 ")
  log.Println("开始")
    serverUrl, err := url.Parse("http://ender.appspot.com")
    if err != nil {
      log.Fatal("URL failed to parse")
    }
  reverseProxy := NewSingleHostReverseProxy(serverUrl)
  http.Handle("/", reverseProxy)
  log.Println("监听:", 9000)
  if err = http.ListenAndServe(":9000", nil); err != nil {
    log.Fatal(err)
  }
  log.Println("服务启动")
}
