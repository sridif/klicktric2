package api

import (
  "github.com/stretchrcom/goweb"
  "github.com/stretchrcom/goweb/context"
  "net/url"
  "fmt"
  //"time"
  "db" 
  //"math/rand"
  "github.com/streadway/simpleuuid"
  "time"
)


type SenorData struct{

  Value []string

}
type SensorController struct {
  
  Datas []*SenorData
}

func (r *SensorController) Read(brand string, ctx context.Context) error{


  fmt.Println(brand)
  query := ctx.HttpRequest().URL.RawQuery 
  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")  
  V, _ := url.ParseQuery(query)
  id:= V.Get("id")
  data:= V.Get("data") 
  
  //db.Write("qcoild", id , )  
  now := time.Now()
  uuid, _ := simpleuuid.NewTime(now)
  fmt.Println(uuid.String()) 
  fmt.Println(id)
  fmt.Println(data)
  data2 := new(SenorData)
  
  if brand == "read" {
  
    data2.Value = db.SliceReadValues2("swatbots", "qicoil", id)

  }
  if brand == "write"{
    db.WriteValue("swatbots", "qicoil", id , string(uuid.Bytes()), data)
  }

  return goweb.API.RespondWithData(ctx , data2)
 
}
