package api

import (
  "github.com/stretchrcom/goweb"
  "github.com/stretchrcom/goweb/context"
  //"net/url"
  "fmt"
  "time"
  "db" 
  "math/rand"

)


type NewsbyidController struct {
  
  Datas []*IceData
}

func (r *NewsbyidController) Read(brand string, ctx context.Context) error{

  rd := rand.New(rand.NewSource(time.Now().UnixNano()))
  rowname := ctx.HttpRequest().URL.RawQuery 
  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")  
  
 
  if rowname == "random" {
      rid:= db.ReadColumnsSimple("ice", "ParsedVer01", "2013-09-18")
      rowname = rid[ rd.Intn( len(rid) ) ]
      //fmt.Println(rowname)
   }

  newsclusters := db.ReadValues("ice", "ParsedVer02", rowname )  

  //for _, cluster := range newsclusters {
    //fmt.Println(cluster)
  //}
  fmt.Println("rest")
  data := new(IceData)
  data.Output = newsclusters
  data.TimeStamp=[]string{"testing", "timestamp" }
  return goweb.API.RespondWithData(ctx , data)
}
