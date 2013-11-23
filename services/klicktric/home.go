package klicktric

import (

  "html/template"
  "net/http"  
  "net/url"
  "encoding/xml"
  "log"
  "io/ioutil"
  "encoding/json"
  "bytes"
  "strings"
)

func Api(rw http.ResponseWriter, req *http.Request){

   t, _:= template.ParseFiles(Template + "api.html")
   t.Execute(rw, nil) 
  
}  


type Body struct{
  Client_id string
  Client_secret string
  Code string
  Grant_type string
  Redirect_uri string
}

func Game(rw http.ResponseWriter, req *http.Request){
  
   if len(req.URL.RawQuery) > 10 {
   code := req.URL.RawQuery[len("code="):]
   log.Println(code)
   energy := &Body{"greengridz",
                    "bOdxOSyXFLnQIWe8wIhbdnW6",
                    code,
                    "authorization_code",
                    "http://www.klicktric.com/game"} 

   buf, _ := xml.Marshal(energy)
   log.Printf(string(buf))
   body := bytes.NewBuffer(buf)
   resp , err := http.Post("https://app.energimolnet.se/oauth2/grant", "text/html", body )
   
   
   //6593cd5e93837d00a06bfb73e155885f665c572e
   log.Println(err)
   log.Println(resp) 

   resp , err = http.PostForm("https://app.energimolnet.se/oauth2/grant", url.Values{ "client_id" : {"greengridz"} ,
       "client_secret" : {"bOdxOSyXFLnQIWe8wIhbdnW6"},
       "code" : {code},
       "grant_type": {"authorization_code"},
       "redirect_uri": {"http://www.klicktric.com/game"}  })

   log.Println(err)
   log.Println(resp) 
 
   body2, _ := ioutil.ReadAll(resp.Body)
   log.Println(string(body2) )
   var data DataSeg2
   err = json.Unmarshal(body2, &data)

   //ck2 := &http.Cookie{Name :"access_token" , Value : data.Access_token } 
  // http.SetCookie(rw, ck2)

   // get meter id 
   
   resp2 , _ := http.Get("http://app.energimolnet.se/api/1.1/users/me/meters?access_token=" + data.Access_token )
   body3, _ := ioutil.ReadAll(resp2.Body)
   log.Println(string(body3) )
   split := strings.Split(string(body3) , "\"")
   log.Println(split[3]) 


   // get meter data
   resp2 , _ = http.Get("http://app.energimolnet.se/api/1.1/users/me/series/"+ split[3]+ "?query=[20131001]&access_token=" + data.Access_token )
   body3, _ = ioutil.ReadAll(resp2.Body)
   log.Println(string(body3))

  // get meter data
   resp2 , _ = http.Get("http://app.energimolnet.se/api/1.1/users/me/series/"+ split[3]+ "?query=[20131002]&access_token=" + data.Access_token )
   body4, _ := ioutil.ReadAll(resp2.Body)
   log.Println(string(body4))


   
 //  var data2 Meter
 //  json.Unmarshal(body3, &data2)
 //  log.Println(data2)

    data2 :=&DataSeg2{"empty" , string(body3) + string(body4)}
   t, _:= template.ParseFiles(Template + "game.html" )
   t.Execute(rw, data2) 
   }
   if len(req.URL.RawQuery) < 10 {
  //   at, _ :=req.Cookie("access_token")
     data := &DataSeg2{"empty" ," [[0 ,0]]"}

   t, _:= template.ParseFiles(Template + "game.html" )
   t.Execute(rw, data)
   }

  
}
type Meter struct {
  _Id string
}
type DataSeg2 struct {
 Access_token string
 Meterdata string

 
}
type DataSeg struct {

  Data UseridSeg

}

type UseridSeg struct {

  User_id int

}
type User_info struct {
   Id string
   Gender string
   Name string
}
 
func GetUserInfo (userid string) User_info {
  appid := "?access_token=1376200959286911|uVPsuyz-BE-or0YhVsEa4XSLR4Q"
   resp , _ := http.Get("https://graph.facebook.com/" + userid + appid )
   body, _ := ioutil.ReadAll(resp.Body)
   log.Println(string(body) )
   var data User_info
   json.Unmarshal(body, &data)
   return data
}
func Profile(rw http.ResponseWriter, req *http.Request){
   
   req.ParseForm()

   if len(req.Form["access_token"]) > 0 { 
   appid := "&access_token=1376200959286911|uVPsuyz-BE-or0YhVsEa4XSLR4Q"
   resp , _ := http.Get("https://graph.facebook.com/debug_token?input_token=" + req.Form["access_token"][0] +appid )

   log.Println(resp)

   body, _ := ioutil.ReadAll(resp.Body)
   log.Println(string(body) )
   var data DataSeg
   err := json.Unmarshal(body, &data)
   log.Println(data.Data.User_id)
   log.Println(err)
   userid := string(data.Data.User_id)
   GetUserInfo( userid )   
   ck2 := &http.Cookie{Name :"userid" , Value : userid} 
   http.SetCookie(rw, ck2)
   
  

   }

   t, _:= template.ParseFiles(Template + "profile.html")
   t.Execute(rw, nil) 
 
}

func fblogin(rw http.ResponseWriter, req *http.Request){

    t, _ := template.ParseFiles(Template + "temp_profile.html")
    t.Execute(rw,nil)
    
}

func BaseHandler(rw http.ResponseWriter, req *http.Request) {

    t, _ := template.ParseFiles(Template + "home.html")
    t.Execute(rw,nil)

}		
