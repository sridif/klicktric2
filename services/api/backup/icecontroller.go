package api

import (
  "github.com/stretchrcom/goweb"
  "github.com/stretchrcom/goweb/context"
  //"net/url"
  "fmt"
  //"time"
  "db"

)

type IceData struct {
  
  Output []string
  TimeStamp []string

}

type IceController struct {
  
  Datas []*IceData
}

func (r *IceController) Read(brand string, ctx context.Context) error{

  rowname := ctx.HttpRequest().URL.RawQuery 
  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")  
  
  newsclusters := db.ReadValues("ice", "clusters", rowname )  

  for _, cluster := range newsclusters {
    fmt.Println(cluster)
  }
  fmt.Println("rest")
  data := new(IceData)
  data.Output = newsclusters
  data.TimeStamp=[]string{"testing", "timestamp" }
  return goweb.API.RespondWithData(ctx , data)
}
