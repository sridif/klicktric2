package klicktric

import (
  "net/http"
//  "html/template"
  "log"
)

var Template = "/home/ubuntu/dev/src/services/klicktric/templates/"

func StartServer(){

  http.HandleFunc("/",BaseHandler)
  http.HandleFunc("/api", Api)
  http.HandleFunc("/game", Game)
  http.HandleFunc("/profile", Profile)
  http.HandleFunc("/callback",fblogin)
  log.Fatal(http.ListenAndServe(":9921", nil))

}

