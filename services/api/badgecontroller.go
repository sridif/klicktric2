package api

import (
    "github.com/stretchrcom/goweb"
   "github.com/stretchrcom/goweb/context"
   "fmt"

)

type Gamedata struct{
  
  Names []string
  Images []string
  Desc  []string
  
}

type BadgeController struct{
 Data []*Gamedata

}

func (r *BadgeController) Read(query string, ctx context.Context) error{
  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*") 
  data := new(Gamedata)
  data.Names = []string{"energymolnet" , "fill profile"}
  data.Images = []string{"https://docs.google.com/drawings/d/1p-EVlpUSCjLiIgwdjOVZ24K6U5xH40AliayDG9vP2PA/pub?w=305&h=381", "https://docs.google.com/drawings/d/1B51Z6i4OLBaFsr2zuni6CyDp_hM9H09Aw9eWuS5j96M/pub?w=449&h=454"}
  data.Desc = []string{"Sync your enrgymolnet account" , "add description of the type of equipments you have"}

  fmt.Println(query)

  return goweb.API.RespondWithData(ctx, data)
}