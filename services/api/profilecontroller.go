package api

import (
  "github.com/stretchrcom/goweb"
  "github.com/stretchrcom/goweb/context"
  //"net/url"
  //"fmt"
  //"time"
  //"db" 
  //"math/rand"
  //"github.com/streadway/simpleuuid"
  //"time"
)


type ProfileData struct{

  Name string
  Fbid string
  Link string
  EnergySync string

}
type ProfileController struct {

}
 
func (r *ProfileController) Read(brand string, ctx context.Context) error{

  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")
  data := new(ProfileData)
  data.Name = "Sriram Elango"
  data.Fbid = "638449578"
  data.Link = "https://www.facebook.com/sridif"
  data.EnergySync ="True"
  return goweb.API.RespondWithData(ctx , data)
 
}
