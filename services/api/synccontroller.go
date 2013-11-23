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


type SyncData struct{

  Status string

}
 
type SyncController struct {
}
 
func (r *SyncController) Read(brand string, ctx context.Context) error{

  fmt.Println(brand)
  query := ctx.HttpRequest().URL.RawQuery 
  fmt.Println(query)
  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")
  data := new(SyncData)
  data.Status = query + " synced with " + brand
  
  return goweb.API.RespondWithData(ctx , data)
 
}
