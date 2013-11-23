package swatbots

import (
  "net/http"
  "html/template"
  "log"
)

var Template = "/home/ubuntu/dev/src/services/swatbots/templates/"

func StartServer(){


  http.HandleFunc("/",BaseHandler)
  log.Fatal(http.ListenAndServe(":9941", nil))

  http.HandleFunc("/test",testHandler)
  log.Fatal(http.ListenAndServe(":9941", nil))

}


func BaseHandler(rw http.ResponseWriter, req *http.Request) {

    t, _ := template.ParseFiles(Template + "home.html" , Template + "home.js")
    t.Execute(rw,nil)

}

func testHandler(rw http.ResponseWriter, req *http.Request) {

    t, _ := template.ParseFiles(Template + "home.html" , Template + "home.js")
    t.Execute(rw,nil)

}