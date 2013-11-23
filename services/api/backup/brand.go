package api

import (

  "db"
  "strings"
  "fmt"
  "time"
  "github.com/streadway/simpleuuid"
  "encoding/json"
  "math/rand"  
)
 
func GetRanking(name string )(Data []string,Trend []string){

  Data = db.ReadValues("universum", "GlobalRankings", name)  
  //OldData = db.ReadValues("universum", "GlobalRankings", name)
 
  for _ = range Data {
    i := rand.Int()
    
    if i%2 == 0 {
      Trend= append(Trend, "1")
    }  
    if i%2 != 0 {
      Trend= append(Trend, "0")
    } 
  }
  return

}
/*
func GetFollwers(name string, presenceMap map[string]string, since time.Time, until time.Time) (Data []string, Timestamp []string){  

  fbpages := presenceMap["FB"]
  Data, Tstamp := HelperAggregate("universum", "FacebookPage" ,"likes" , fbpages, since,until)
  //twpages := presenceMap["TW"]

  //inpages := presenceMap["IN"]
  return
}*

func GetLinkedinTypeA(name string, id []string, since time.Time, until time.Time)(Data []string, Timestamp []string){

*
  values := db.ReadValues("universum","LinkedinCompanies",id)
  columns := db.ReadColumns("universum","LinkedinCompanies",id) 
  data, timestamp = GetMetricWindow(name, values, columns, since, until) 

*

  Data, Tstamp := HelperAggregate("universum", "LinkedinCompanies" , name, id, since,until)
  for i := range Tstamp {
    Timestamp = append(Timestamp, Tstamp[i].String())    
  }

  return

}
func GetTwitterTypeA(name string, id []string, since time.Time, until time.Time)(Data []string, Timestamp []string){

*
  values := db.ReadValues("universum","TwitterUserShow",id)
  columns := db.ReadColumns("universum","TwitterUserShow",id) 
  data, timestamp = GetMetricWindow(name, values, columns, since, until)

*

  Data, Tstamp := HelperAggregate("universum", "TwitterUserShow" , name, id, since,until)

  for i := range Tstamp {
    Timestamp = append(Timestamp, Tstamp[i].String())    
  }
  
  return

}

func GetTwitterTypeB(name string, id []string, since time.Time, until time.Time)(Data []string, Timestamp []string){

 * values := db.ReadValues("universum","TwitterUserTimelineAggregated",id)
  columns := db.ReadColumns("universum","TwitterUserTimelineAggregated",id)* 

  Data, Tstamp := HelperAggregate("universum", "TwitterUserTimelineAggregated" , name, id, since,until)
  for i := range Tstamp {
    Timestamp = append(Timestamp, Tstamp[i].String())    
  }
  
  return

}

func GetFacebookTypeB(name string, id []string, since time.Time, until time.Time)(Data []string, Timestamp []string){
  //facebook =[]string{"0","0","0","0"}
 
*  values := db.ReadValues("universum","FacebookStreamAggregated",id)
  columns := db.ReadColumns("universum","FacebookStreamAggregated",id) 
  data, timestamp = GetMetricWindow(name, values, columns, since, until)  
*

  Data, Tstamp := HelperAggregate("universum", "FacebookStreamAggregated" , name, id, since,until)
  for i := range Tstamp {
    Timestamp = append(Timestamp, Tstamp[i].String())    
  }
  
  return

} 
func GetFacebookTypeA(name string, id []string, since time.Time, until time.Time)(Data []string, Timestamp []string){

  Data, Tstamp := HelperAggregate("universum", "FacebookPage" , name, id, since,until)
  for i := range Tstamp {
    Timestamp = append(Timestamp, Tstamp[i].String())    
  }

  return

} 
*/


type Content struct {
  Geo string
  NumFollowers string
  Comments_count string
  Posts_count string
  Likes string
  Likes_count string 
  Shares_count string
  Followers_count string
  Post_id string
  Post string
  Share string
  Retweet_count string

}
func GetFromJson(jsonstr string, name string) ( string ){
  //fmt.Println(jsonstr)
  
  var content []Content
  err := json.Unmarshal([]byte("["+ jsonstr + "]"), &content)  
  if err != nil {
     fmt.Println("error:", err)
     return "x"
  }
  
  switch name {
    case "geo":
      return content[0].Geo
    case "Post" : 
      return content[0].Post
    case "likes_count" :
      return content[0].Likes_count
    case "shares_count" :
      return content[0].Shares_count
    case "comments_count" :
      return content[0].Comments_count
    case "likes" :
      return content[0].Likes
    case "followers_count":
      return content[0].Followers_count
    case "numFollowers" :
      return content[0].NumFollowers
    case "retweet_count" :
      return content[0].Retweet_count
    default :
      return "x"
  }
  return "x"
  
 
}

func GetMetricWindowTIME(name string, values []string, columns []string, since time.Time , until time.Time)(data []string, tstamp []time.Time) {

//  tstamp :=[]time.Time{}
  count :=0
  for i := range values {

      testtime,_ := simpleuuid.NewString(columns[i])
      recorded := testtime.Time()
      fmt.Println(recorded)
      if recorded.After(until){

        return

      }
      if recorded.After(since) {
      fmt.Println(recorded)

      data =append(data, GetFromJson(values[i],name))
      //timestamp =append(timestamp,recorded)
      tstamp = append(tstamp,recorded)
      count = count +1
      }
   
  } 
  return

} 
func GetMetricWindow(name string, values []string, columns []string, since time.Time , until time.Time)(data []string, timestamp []string) {

  tstamp :=[]time.Time{}
  count :=0
  for i := range values {

      testtime,_ := simpleuuid.NewString(columns[i])
      recorded := testtime.Time()
      fmt.Println(recorded)
      if recorded.After(until){

        return

      }
      if recorded.After(since) {
      fmt.Println(recorded)

      data =append(data, GetFromJson(values[i],name))
      timestamp =append(timestamp,recorded.String())
      tstamp = append(tstamp,recorded)
      count = count +1
      }


   
  } 
  return

}

func GetAttributes(name string, id string, since time.Time, until time.Time)(data []string, timestamp []string){

  values := db.ReadValues("universum","AttributesByCompanyId",id)
  
  //columns := db.ReadColumns("universum","AttributesByCompanyId",id) 
  data, timestamp = GetMetricWindow(name, values, values, since, until)
  
  return

}



func HelperAggregate(keyspace string, columnfamily string, name string, id []string, since time.Time, until time.Time, resolution string) (Data []string, Tstamp []time.Time){

  //Tstamp := []time.Time{} 
  for i := range id {
    fmt.Println(id[i])
    fmt.Println(keyspace + " " +columnfamily + " " + id[i])
    values := db.ReadValues(keyspace, columnfamily, id[i])
    columns := db.ReadColumns(keyspace, columnfamily, id[i]) 
    data, tstamp := GetMetricWindowTIME(name, values, columns, since, until)
    if i == 0 {
      Data = data
      Tstamp = tstamp
    }
    if i>= 1 {
      Data, Tstamp = Aggregate(Data, Tstamp, data, tstamp)
    }    
  } 

  return   

} 


/*
func GetFacebookPostLikes(id string, since time.Time, until time.Time)(data []string, timestamp []string){
  fmt.Println(id)

  values := db.ReadValues("universum","FacebookStreamAggregated",id)
  columns := db.ReadColumns("universum","FacebookStreamAggregated",id) 

    
  data, timestamp = GetMetricWindow("likes_count", values, columns, since, until)  

  return
}
*/
func GetPages(brandname string )(pages []string){
  
  values:= db.Read("universum","manualEntry",brandname)
  fmt.Println("wowowo")
  if (len(values)>0){
    split := strings.Split(values[0],"\n")
    for i:= range split {
      split2 := strings.Split(split[i],"//")
      if(len(split2) > 3){
      pages = append(pages,split2[1],split2[0], split2[3], split2[2])
      }
    }

  }
 
  return

}

func GetFirstFacebookId(brandname string)(id string){
  id = "40796308305"

  values:= db.Read("universum","manualEntry",brandname)
  fmt.Println("wowowo")
  if (len(values)>0){
    split := strings.Split(values[0],"\n")
    for i:= range split {
      split2 := strings.Split(split[i],"//")
      if(len(split2)>3){
        if(split2[0]=="FB:"){
          id=split2[1]
        }
      }
    }
  }

  return

}
func GetFirstLinkedinId(brandname string)(id string){
  id = "1694"

  values:= db.Read("universum","manualEntry",brandname)
  fmt.Println("wowowo")
  if (len(values)>0){
    split := strings.Split(values[0],"\n")
    for i:= range split {
      split2 := strings.Split(split[i],"//")
      if(len(split2)>3){
        if(split2[0]=="IN:"){
          id=split2[1]
        }
      }
    }
  }

  return

}
func GetFirstTwitterId(brandname string)(id string){

  id = "18881750"
  values:= db.Read("universum","manualEntry",brandname)
  
  fmt.Println("wowowo")
  if (len(values)>0){
    split := strings.Split(values[0],"\n")
    for i:= range split {
      split2 := strings.Split(split[i],"//")
      if(len(split2)>3){
        if(split2[0]=="TW:"){
          id=split2[1]
        }
      }
    }
  }

  return

}


func GetFacebookPageShare(id string)(data []string){
 
  data =[]string{"0","0","0","0"}
  values := db.ReadValues("universum","FacebookStream",id)
 
  for i := range data{
    if (i < len(values)) {
      split := strings.Split(values[i],"\",\"")
      split1 := strings.Split(split[3],"\"")
      data[i]=split1[2]

   }
  } 

  return
}

func GetFacebookPageComment(id string)(data []string){
 
  data =[]string{"0","0","0","0"}
  values := db.ReadValues("universum","FacebookStream",id)
 
  for i := range data{
    if (i < len(values)) {
      split := strings.Split(values[i],"\",\"")
      split1 := strings.Split(split[5],"\"")
      data[i]=split1[2]

   }
  } 
 
  return
}

func GetTwitterRetweets(id string)(twitter []string){

  twitter =[]string{"0","0","0","0"}
  valuesT := db.ReadValues("universum","TwitterUserTimeline",id)

  count := -1
  // count is set to iterate until 4
  // for the latest 4 posts
  currentid := "0"
  for i := range valuesT{
      //fmt.Println("current id : " + currentid + "\n")
      split := strings.Split(valuesT[i],",")
      split0 := strings.Split(split[0],"\"")
      readid := split0[3]
      if (readid != currentid){
        count = count +1
        currentid = readid

        if( count < 4 ){
                  fmt.Println(currentid)
          split2 := strings.Split(split[1],"\"")
          twitter[i]= split2[3]
        }
        if (count >4){
        break
       }
      }

 }
  return 
}

func GetLinkedinFollowers(id string)(linkedin []string){

  linkedin =[]string{"0","0","0","0"}
  valuesL := db.ReadValues("universum","LinkedinCompanies",id)
  for i := range linkedin{
    if (i<len(valuesL)){
    split := strings.Split(valuesL[i],",")
    split2 :=	strings.Split(split[2],"\"")
    linkedin[i]= split2[3]
    //fmt.Printf(split2[3])
    //split2 = strings.Split(split[1],"\"")
    //social[0] = split2[3]
    }
  }
  return
  
}




