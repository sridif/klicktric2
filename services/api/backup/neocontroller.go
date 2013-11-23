package api

import (
  "github.com/stretchrcom/goweb"
  "github.com/stretchrcom/goweb/context"
  //"net/url"
  "fmt"
  //"time"
  "db"
  //"github.com/jmcvetta/neoism"
  //"strconv"
)


type NeoData struct {

  Properties []string

}
type NeoController struct {
  Datas []*NeoData
}

func (r *NeoController) Read(brand string, ctx context.Context) error{

  rowname := ctx.HttpRequest().URL.RawQuery 
  fmt.Println(rowname)
  data := new(NeoData)
  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")
  data.Properties = db.ReadValues("tg","neoindex", rowname)
  return goweb.API.RespondWithData(ctx , data)
  
}