package api

import (
  "fmt"
  "time"
  "strconv"
)

func AggregateOverDay(metric []string,timestamp []time.Time){


  data := []string{}
  newtime:= []time.Time{}
  current:= timestamp[0]

  for i :=range timestamp{

    t:= timestamp[i].Truncate(24*time.Hour)    
    if !current.Equal(t){
      
      current = t
      data = append(data,metric[i])
      newtime=append(newtime , t)       
    }
    
  }
  
  Aggregate (data,newtime,data,newtime)
  fmt.Println(data)
}

/*
func AggregateSTRING( metric []string, tstamp1 []string, metric2 []string, tstamp2  []string) ([]string, []time.Time){
 
  timestamp1:= []time.Time
  timestamp2 := []time.Time
  const shortForm = "2006-Jan-02"
  //since, _ := time.Parse(shortForm, since_str)

  for i := range timestamp1 {

    presentTime, _ = time.Parse(shortForm, tsamp1[i])
    timestamp1 = aggregate( timestamp1, presentTime )    
  }

  for i := range timestamp2 {
    presentTime, _ = time.Parse(shortForm, tsamp2[i])
    timestamp2 = aggregate(timestamp2, tstamp2[i])
    
  }
 

}*/

func Aggregate( metric1 []string, timestamp1 []time.Time, metric2 []string, timestamp2 []time.Time ) ([]string , []time.Time){

  aggregatedmetric := []string{}
  aggregatedtimestamp := []time.Time{} 
  index1 := 0
  length1 := len(metric1)
  index2 := 0
  length2 := len(metric2)

  index1valid := true
  index2valid := true
  latestmetric1 := "0"
  latestmetric2 := "0"

  for ( index1valid || index2valid ) {
    if index1valid && (timestamp1[index1].Equal(timestamp2[index2]) || timestamp1[index1].Before(timestamp2[index2])  || !index2valid ) {
        latestmetric1 = metric1[index1]
        aggregatedmetric = append(aggregatedmetric, Add(latestmetric1, latestmetric2) )
        aggregatedtimestamp = append(aggregatedtimestamp, timestamp1[index1])
 
        fmt.Println("index1 - ps1 :", index1)

        index1 ++
        if index1 >= length1 {
          index1valid = false
          index1 --
        }
      continue
    }
    if index2valid && (timestamp2[index2].Before(timestamp1[index1]) || !index1valid ) {
        latestmetric2 = metric2[index2]
        aggregatedmetric = append(aggregatedmetric,Add(latestmetric1, latestmetric2) )
        aggregatedtimestamp = append(aggregatedtimestamp, timestamp2[index2])
        fmt.Println("index2 - ps2 :", index2)

        index2 ++
        if index2 >= length2 {
          index2valid = false
          index2 --
        }
      continue
    }

  }
  fmt.Println(aggregatedmetric)  

 return aggregatedmetric, aggregatedtimestamp 
}

func Add(a string, b string) (c string) {

  c = a + " + " + b
  inta, _ := strconv.ParseInt(a,10,64)
  intb, _ := strconv.ParseInt(b,10,64)
  c = strconv.FormatInt(inta + intb , 10)
  return
}
  /*
  for {
    if index1valid {
        
      if timestamp1[index1].Before(timestamp2[index2]) {
        aggregatedmetric = append(aggregatedmetric, metric1[index1])
        aggregatedtimestamp = append(aggregatedtimestamp, timestamp1[index1])
        index1 ++
        fmt.Println("index1 - ps1 :", index1)
        if index1 >= length1 {
          fmt.Println("just append the remaining metric2 stuff ")
          index1valid = false 
          //break
        }
        continue
      }
      if !index2valid {
          aggregatedmetric = append(aggregatedmetric, metric1[index1])
          aggregatedtimestamp = append(aggregatedtimestamp, timestamp1[index1])
          index1 ++
          fmt.Println("index1 - ps1 :", index1)

          if index1 >= length1 {
            fmt.Println("just append the remaining metric2 stuff ")
            index1valid = false 
            //break
          }
          continue
      } 
    }

    if index2valid {
  
      aggregatedmetric = append(aggregatedmetric, metric2[index2])
      aggregatedtimestamp = append(aggregatedtimestamp, timestamp2[index2])
      index2++
      fmt.Println("index2 -ps2:" , index2)
      if index2 >= length2 {
        if index1valid {
          for {
             aggregatedmetric = append(aggregatedmetric, metric2[index1])
             aggregatedtimestamp = append(aggregatedtimestamp, timestamp2[index1])
             index1++
             fmt.Println("index1 -ps3", index1)
             if index1 >= length1 {
	       break
	     }
              
          }
       } 
       index2valid = false
      }
      continue
    }
    break
  } 
  */
