package api

import (
  "github.com/stretchrcom/goweb"
  "github.com/stretchrcom/goweb/context"
  //"net/url"
  "fmt"
  //"time"
  "db"
  "strings"
)

type ClusterData struct {
  
  Node0 []string
  Node1 []string
  Node2 []string
  Node3 []string
  Nodek []string
  Nodev []float64
}

type ClusterController struct {
  
  Datas []*ClusterData
}

func (r *ClusterController) Read(brand string, ctx context.Context) error{
  data := new(ClusterData)
  

  rowname := ctx.HttpRequest().URL.RawQuery 
  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*") 
  data.Node0 = []string{rowname}

  data.Node1 =[]string{}
  data.Node2 =[]string{}
  data.Node3 =[]string{}
  data.Nodek = []string{}
  data.Nodev = []float64{}

  fmt.Println(rowname)
  newsclusters := db.ReadValues("ice", "ParsedVer02", rowname )  

  //for _, cluster := range newsclusters {
    //fmt.Println(cluster)
  //}

  entities := strings.Split(newsclusters[1], "#")  
  
 
  id_counter := make(map[string]float64)
  for _ , entity := range entities {
    ids := db.ReadColumnsSimple( "ice", "EntityIndex" , entity)

    curr_len := len(ids)
    if (curr_len < 50) { 
     for _,id := range ids{
          id_counter[id]= id_counter[id] + float64(1)/float64(curr_len)
       }
    }
    
         
   }
   Length := id_counter[rowname]
   //fmt.Println("checkign the id_counter")
   for k, v := range id_counter {
      fmt.Println(k)
      fmt.Println(v)  
      switch  {

      case v < Length/3 :
           data.Node3 = append(data.Node3, k)
      case v < Length*2/3:  
           data.Node2 = append(data.Node2, k)
      case v < Length:
           data.Node1 = append(data.Node1, k)
      }
      data.Nodek = append(data.Nodek, k)
      data.Nodev = append(data.Nodev, v)

  } 
  //fmt.Println(id_counter)
 

  return goweb.API.RespondWithData(ctx , data)

}

func getTitleIdfromId() []string {
   
   return []string{"booyay"}
}