package api

import (
  "github.com/stretchrcom/goweb"
  "github.com/stretchrcom/goweb/context"
  //"net/url"
  "fmt"
  //"time"
  "db"
  "strings"
  "sort"
  "strconv"
)


type NounphraseController struct {
  Datas []*IceData
}

func (r *NounphraseController) Read(brand string, ctx context.Context) error{

  rowname := ctx.HttpRequest().URL.RawQuery 
  rowname = strings.Replace(rowname, "%20" , " ", -1)
  ctx.HttpResponseWriter().Header().Set("Access-Control-Allow-Origin", "*") 
  act_rowname := db.ReadColumnsSimple("ice", "EntitymapVer01", rowname)

  nphrases := []string{" "}
  art_ids := []string{" "}
  fmt.Println(rowname)
  fmt.Println(act_rowname)

  np_dict := make(map[string][]string)
  np_count := make(map[string]float64)
  for _ , act := range act_rowname { 
    articles := db.ReadColumnsSimple("ice", "EntityIndex", act)
    
    for _ , article := range articles {
      curr_nphrases := db.ReadColumnsSimple("ice", "NounPhrase", article )
      //nphrases = append(nphrases , curr_nphrases[0] )
      for _, np := range curr_nphrases {
        np_dict[np] = append( np_dict[np] , article  )
      }
     
 
    }
  } 

  fmt.Println("the high freq")

  /*
  for k, v := range np_dict {

    np_count[k]= len(v) 
    if len(v) > 1 {
     // fmt.Println( k )
     nphrases = append(nphrases, k)
     curr_id_str := ""
     for _, new_id := range v { 
        
         curr_id_str = curr_id_str + "#" + new_id

     }

     art_ids = append( art_ids , curr_id_str )
    }
  } */

  
  for k, v := range np_dict {
    if len(v) > 1 {
      tf := float64( len(v) )
      df_str := db.ReadColumnsSimple("ice", "InvNounPhraseDF", k )[0]
      df , err := strconv.ParseFloat(df_str , 10)

      //if err != nil {
      //np_count[k] = 0
      
      if err == nil {
       if df < float64(20) {
        np_count[k] = tf / df
       }
      }
      
      //fmt.Println("tfidf : %s_%g_%g", k, tf, df  )
      // float64(len(document))
    }
  }
  

  fmt.Println("entering soriting")
  vs := NewValSorter(np_count)
  vs.Sort()
  
  limit := 10
  if limit > len(vs.Keys) {
    limit = len(vs.Keys)
  }

  for k := 0 ; k< limit ; k++ {
    nphrases = append(nphrases, vs.Keys[k])
    curr_id_str := ""
    for _, new_id := range np_dict[vs.Keys[k]] {
       curr_id_str = curr_id_str + "#" + new_id
    }  
    art_ids = append(art_ids , curr_id_str )
  }   

  fmt.Println(vs.Keys)    
  //fmt.Println(np_dict)
  fmt.Println("rest")
  //fmt.Println(nphrases)
  data := new(IceData)
  data.Output = nphrases
  data.TimeStamp= art_ids
  return goweb.API.RespondWithData(ctx , data)

}

type ValSorter struct {
        Keys []string
        Vals []float64
}
 
func NewValSorter(m map[string]float64) *ValSorter {
        vs := &ValSorter{
                Keys: make([]string, 0, len(m)),
                Vals: make([]float64, 0, len(m)),
        }
        for k, v := range m {
                vs.Keys = append(vs.Keys, k)
                vs.Vals = append(vs.Vals, v)
        }
        return vs
}
 
func (vs *ValSorter) Sort() {
        sort.Sort(vs)
}
 
func (vs *ValSorter) Len() int           { return len(vs.Vals) }
func (vs *ValSorter) Less(i, j int) bool { return vs.Vals[i] > vs.Vals[j] }
func (vs *ValSorter) Swap(i, j int) {
        vs.Vals[i], vs.Vals[j] = vs.Vals[j], vs.Vals[i]
        vs.Keys[i], vs.Keys[j] = vs.Keys[j], vs.Keys[i]
}