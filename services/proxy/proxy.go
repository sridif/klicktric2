/* 
   
Access point code

Todo: 
1: figuring out hte port number issues.
2: fluentze mention universum media pilot apps .. got go here.


*/

package proxy
import (
  "fmt"
  //"log"
  "flag"
  "net/http"
  "net/http/httputil"
)


var addr = flag.String("addr", "test.swatbots.com:9921", "address of the service ")

func StartService() {

  fmt.Println("test")
  proxy := httputil.ReverseProxy{Director : func(req *http.Request) {
    req.URL.Scheme = "http"
    
    switch req.Host {
     /*
      case "www.pinnar.com":
     
      case "pinnar.com":
        req.URL.Host = "localhost:" */
      case "games.swatbots.com":
        req.URL.Host = "localhost:9999"

      case "api.klicktric.com" : 
        req.URL.Host = "localhost:9925"

      case "www.klicktric.com" :
        req.URL.Host = "localhost:9921"

      case "www.swatbots.com":
        req.URL.Host = "localhost:9931"

      case "api.swatbots.com" :
        fmt.Println("inside api")
        req.URL.Host = "localhost:9925"
      
      case "test.swatbots.com":
         fmt.Println("Inside Test")
         req.URL.Host = "localhost:9941"
      
      default :
        fmt.Println("inside defualt")
        req.URL.Host = "localhost:9925"

    }
  
  }}

  err := http.ListenAndServe(*addr, &proxy)
  fmt.Println(err)
}

 






    
