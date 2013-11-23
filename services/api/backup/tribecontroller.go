package api

import (
  "github.com/stretchrcom/goweb"
  "github.com/stretchrcom/goweb/context"
  //"net/url"
  "fmt"
  //"time"
  "db"
  "github.com/jmcvetta/neoism"
)


type TribeData struct {

  People []string
  Company []string
  Links []string

}
type TribeController struct {
  Datas []*TribeData
}

func (r *TribeController) Read(brand string, ctx context.Context) error{

  rowname := ctx.HttpRequest().URL.RawQuery 
  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")  
  

  //newsclusters := db.ReadColumnsSimple("ice", "entitylist", rowname )  
  
  neodb,_ := neoism.Connect("http://localhost:7474/db/data")
  
  People, _ := neodb.LegacyNodeIndex("Company")
  Res , _ :=People.Find("name" , "Sriram Elango")
  fmt.Println( Res)  
  //fmt.Println(Res)
  fmt.Println(rowname)
  data := new(TribeData)
  data.People = db.ReadColumnsSimple("tg", "neoindex" ,"people" )
  data.Company = db.ReadColumnsSimple("tg", "neoindex" , "company" )
  data.Links = db.ReadColumnsSimple("tg","neoindex" , "rel" )

  /*
  data.People = []string{"tom" , "mark"}
  data.Company =[]string{"universum", "pronto" }
  data.Links = []string{ "tom#universum" , "mark#universum" , "mark#pronto", "tom#pronto" }
  */

  return goweb.API.RespondWithData(ctx , data)

}