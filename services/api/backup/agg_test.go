package api

import (
  "testing"
  "fmt"
  "time"
)

func testTimeAggregation(*testing.T) {
  fmt.Println("booyay")

  time1 := time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC) 
  time2 := time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC)
  time3 := time.Date(2009, time.November, 12, 0, 0, 0, 0, time.UTC)
  time4 := time.Date(2009, time.November, 13, 0, 0, 0, 0, time.UTC)
  time5 := time.Date(2009, time.November, 14, 0, 0, 0, 0, time.UTC)
  
  timestamp1 := []time.Time{time2,time3,time5}
  timestamp2 := []time.Time{time1,time4}
  
  data1 := []string{"23", "45" , "60"}
  data2 := []string{"32", "49"}

  Aggregate(data1, timestamp1, data2, timestamp2)

}

func TestJsonStr(*testing.T){

  jsonstr := `{"Post":"0", "Share":"12" }`
  
  fmt.Println(GetFromJson(jsonstr,"Post"))


}