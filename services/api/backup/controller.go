/* 
 1: query to handle api keys and link it to accounts.
 2: 
 
*/
package api

import (

  "github.com/stretchrcom/goweb"
  "github.com/stretchrcom/goweb/context"
  //"net/http"
  "net/url"
  "fmt"
  "time"
  "db"
)

type Data struct {
     Brand   string
     Metrics []string
     TimeStamp []string
}

type RankingData struct {
     Type   string
     Brands []string
     Change []string
}


type UniversumController struct {
  Datas []*Data
}

/*
//only the first one is returned
func GetFirstMetricName(query string) (metric string){

  values, _ := url.ParseQuery(query)
  //fmt.Println( values)
  metrics := values["metric"]
  //fmt.Println( metrics[0])
  if len(metrics) > 0 {
     metric = metrics[0]
  }
  
  return 

} */
func GetQueryValueList( name string, query string) (metrics []string){

  values, _ := url.ParseQuery(query)
  //fmt.Println( values )
  metrics = values[name]
  //fmt.Println( metrics[0])
 
 
  return 

}
func GetQueryValue( name string, query string) (metric string){

  values, _ := url.ParseQuery(query)
  //fmt.Println( values)
  metrics := values[name]
  //fmt.Println( metrics[0])
  if len(metrics) > 0 {
     metric = metrics[0]
  }
 
  return 

}

func DoesItExist(value string, valuearray []string ) (bool) {
 
 for i := range valuearray {
    if (valuearray[i] == value) {
      return true
    }

 }

 return false
  
}

func AndMap(presenceMap1 map[string][]string, presenceMap2 map[string][]string)(map[string][]string){
 
// refactor this !
  fbpages := presenceMap1["FB"]
  newfbpages := []string{}
  for i := range fbpages {
    if DoesItExist(fbpages[i], presenceMap2["FB"]){
      newfbpages = append(newfbpages, fbpages[i])
    }
           
   
  }

  twpages := presenceMap1["TW"]
  newtwpages := []string{}
  for i := range twpages {
    if DoesItExist(twpages[i], presenceMap2["TW"]){
      newtwpages = append(newtwpages, twpages[i])
    }
           
   
  }

  inpages := presenceMap1["IN"]
  newinpages := []string{}
  for i := range inpages {
    if DoesItExist(inpages[i], presenceMap2["IN"]){
      newinpages = append(newinpages, inpages[i])
    }
           
   
  }

  presenceMap := map[string][]string {
    "FB" : newfbpages,
    "IN" : newinpages,
    "TW" : newtwpages,
  }
  return presenceMap
}


// the mapper that maps metric_name to metrics.
func GetMetrics(metric_name string, query string, since time.Time, until time.Time, resolution string)(fmetrics []string, ftimestamp []string){
  metrics := []string{}
  timestamp := []string{}

  brands := GetQueryValueList( "brand" , query )
 
  if brands == nil{
    brands = []string{"empty"}
  }

 ftstamp := []time.Time{}
  
 for ii := range brands {
   brand := brands[ii]  
  /*

  */

  fmt.Println(brands)
  //metric_name := GetQueryValue("name",query)
  
  // if there is an attribute go do an and.
  attribute := GetQueryValue("attribute", query)
  presenceMap := map[string][]string{} 
  if attribute != "" { 
   if brand == "all" {
     presenceMap = GetPresencesfromAttribute(attribute)
    }
    if brand != "all" {
     fmt.Printf("inside all")
     presenceMap1 := GetPresences(brand)  
     presenceMap2 := GetPresencesfromAttribute(attribute)
     presenceMap = AndMap( presenceMap1, presenceMap2 )
    }
  }

  // defenition of type A & B metrics
  // type A : attibutes linked to the presence pages
  // (eg : followers count)
  // type B : attributes linked to the content on presence pages 
  // (eg : total retweets on the pages content)
 
  tstamp := []time.Time{}

  switch metric_name {

    case "ranking" :
      metrics, timestamp = GetRanking(GetQueryValue("basedon", query))  

    case "followers" :
      // metrics, tstamp = GetFacebookTypeA("likes", presenceMap["FB"], since, until)
      metrics, tstamp = HelperAggregate("universum", "FacebookPage" , "likes",presenceMap["FB"], since,until , resolution)
      
      metrics2 , tstamp2 := HelperAggregate("universum", "TwitterUserShow" , "followers_count",presenceMap["TW"], since,until, resolution)

      if len(metrics2) > 0 {
         metrics , tstamp = Aggregate(metrics, tstamp, metrics2 , tstamp2)

      }
 
      metrics3, tstamp3 := HelperAggregate( "universum", "LinkedinCompanies" , "numFollowers",presenceMap["IN"], since,until, resolution)
     if len(metrics3) > 0 {
     	 metrics , tstamp = Aggregate(metrics, tstamp, metrics3 , tstamp3)

     }

    case "totalengagement":
      // metrics, tstamp = GetFacebookTypeA("likes", presenceMap["FB"], since, until)
     // metrics, tstamp = HelperAggregate("universum", "FacebookPage" , "likes", presenceMap["FB"], since, until , resolution)
        metrics, tstamp =HelperAggregate("universum", "FacebookStreamAggregated" , "comments_count",presenceMap["FB"], since,until , resolution)

     
      metrics2 , tstamp2 := HelperAggregate("universum", "FacebookStreamAggregated" , "comments_count",presenceMap["FB"], since,until, resolution)


      if len(metrics2) > 0 {
         metrics , tstamp = Aggregate(metrics, tstamp, metrics2 , tstamp2)

      }
 
      metrics3, tstamp3 := HelperAggregate("universum", "TwitterUserTimelineAggregated", "retweet_count", presenceMap["TW"], since,until, resolution)

     if len(metrics3) > 0 {
     	 metrics , tstamp = Aggregate(metrics, tstamp, metrics3 , tstamp3)

     }


   
        
    case "fblikes" : 
      // metrics,tstamp = GetFacebookTypeA("likes", presenceMap["FB"], since,until)
      metrics, tstamp = HelperAggregate("universum", "FacebookPage" , "likes",presenceMap["FB"], since,until, resolution)


    case "twfollowers" :
      // metrics, tstamp = GetTwitterTypeA("followers_count",presenceMap["TW"], since, until )
      metrics, tstamp = HelperAggregate("universum", "TwitterUserShow" , "followers_count",presenceMap["TW"], since,until, resolution)


    case "twretweets" :
      // metrics,tstamp = GetTwitterTypeB("retweet_count", presenceMap["TW"], since, until, resolution)
      metrics, tstamp = HelperAggregate("universum", "TwitterUserTimelineAggregated", "retweet_count", presenceMap["TW"], since,until ,resolution)


    case "infollowers" :
      // metrics, tstamp = GetLinkedinTypeA("numFollowers", presenceMap["IN"], since , until)
      metrics, tstamp = HelperAggregate("universum", "LinkedinCompanies" , "numFollowers",presenceMap["IN"], since,until, resolution)


    case "fbpostlikes" :
      // metrics,tstamp = GetFacebookTypeB("likes_count",presenceMap["IN"],since,until)
      metrics, tstamp = HelperAggregate("universum", "FacebookStreamAggregated" , "comments_count",presenceMap["FB"], since,until, resolution)


    case "fbpostshare" :
      //metrics,tstamp = GetFacebookTypeB("shares_count",presenceMap["INOB"],since,until)   
      metrics, tstamp = HelperAggregate("universum", "FacebookStreamAggregated" , "comments_count",presenceMap["FB"], since,until, resolution)

    case "fbpostcomment" :
      //metrics,tstamp = GetFacebookTypeB("comments_count",presenceMap["IN"],since,until)
       metrics, tstamp = HelperAggregate("universum", "FacebookStreamAggregated" , "comments_count",presenceMap["FB"], since,until, resolution)


    case "geo" :
     metrics,timestamp = GetAttributes("geo",brand,since,until)
          

    default :
       metrics = []string{}
       timestamp = []string{}
    }

   if (ii == 0 ) {
     fmetrics = metrics
     ftstamp = tstamp
     
   }
   if (ii>=1) {
     fmetrics , ftstamp = Aggregate(fmetrics, ftstamp, metrics, tstamp) 

  }
}
  for i := range ftstamp {
    timestamp = append( timestamp, ftstamp[i].String() )    
  }
  
  ftimestamp = timestamp

  return 

}

//add error catching here.
func (r *UniversumController) Read(brand string, ctx context.Context) error {

  fmt.Println("rest")
  query := ctx.HttpRequest().URL.RawQuery
  //metric_name := GetQueryValue("name",query) 
  since_str := GetQueryValue("since",query) 
  const shortForm = "2006-Jan-02" 
  since, _ := time.Parse(shortForm, since_str) 

  until_str := GetQueryValue("until", query) 
  until, _ := time.Parse(shortForm, until_str) 
  resolution := "day"   

  key := GetQueryValue("apikey",query) 
  fmt.Println("the api key is " + key) 
  user_id := db.Read("universum","apikeys",key)

  if(len(user_id) <1) {
      errObject := "unverified api key"
      return goweb.API.RespondWithError(ctx, 500, errObject)
  }


  if brand == "ranking" {
    ranking := new(RankingData)
    ranking.Type = brand
    ranking.Brands , ranking.Change = GetMetrics(brand, query, since, until, resolution )
    return goweb.API.RespondWithData(ctx, ranking)
  }

  data := new(Data)
  data.Brand  = brand
  data.Metrics, data.TimeStamp = GetMetrics(brand,query, since, until, resolution )
 
  fmt.Println("booyay")
  return goweb.API.RespondWithData(ctx , data)

}