package api

import (
  "github.com/stretchrcom/goweb"
  "github.com/stretchrcom/goweb/context"
  //"net/url"
  "fmt"
  //"time"
  "db"
)


type EntitylistController struct {
  Datas []*IceData
}

func (r *EntitylistController) Read(brand string, ctx context.Context) error{

  rowname := ctx.HttpRequest().URL.RawQuery 
  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")  
  

  newsclusters := db.ReadColumnsSimple("ice", "entitylist", rowname )  

  fmt.Println("rest")
  data := new(IceData)
  data.Output = newsclusters
  data.TimeStamp=[]string{"testing", "timestamp" }
  return goweb.API.RespondWithData(ctx , data)

}