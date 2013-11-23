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


type RewardData struct{

  Badges []string
  Points string

}
type RewardController struct {
}
 
func (r *RewardController) Read(brand string, ctx context.Context) error{


  
  fmt.Println(brand)
  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")
  data := new(RewardData)
  data.Badges = []string{"badge1", "badge2"}
  data.Points = "32"
 
  return goweb.API.RespondWithData(ctx , data)
 
}
