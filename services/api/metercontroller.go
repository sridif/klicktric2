package api

import (
  "github.com/stretchrcom/goweb"
  "github.com/stretchrcom/goweb/context"
  //"net/url"
  "fmt"
  //"time"
  //"db" 
  //"math/rand"
  //"github.com/streadway/simpleuuid"
  //"time"
)


type MeterData struct{

  Resolution string
  Data []string

}
type MeterController struct {

}
 
func (r *MeterController) Read(brand string, ctx context.Context) error{


  query := ctx.HttpRequest().URL.RawQuery 
  fmt.Println(query)
  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")
  data := new(MeterData)
  data.Resolution = "Hourly"
  data.Data = []string{"12.3", "13.4", "16.7" , "12.3", "231.3"}

  return goweb.API.RespondWithData(ctx , data)
 
}
