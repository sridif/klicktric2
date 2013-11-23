package api

import ( 
  "fmt"
  "db"
)

// needs to return the id
func GetPresences(brandname string)( map[string][]string) {
  values:= db.ReadColumnsSimple("universum","PrecenseIdsByCompanyId",brandname)
  var presenceMap = map[string] []string {
    "FB" : []string{},
    "TW" : []string{},
    "IN" : []string{},
  }
  for i := range values {
    fmt.Println(values[i][0:2])
    site := values[i][0:2]
    siteid := values[i][2:len(values[i])]
    presenceMap[site] = append(presenceMap[site] , siteid)  
    
  }
  fmt.Println(presenceMap)
  // mapping the string variable 
  return presenceMap
}

func GetPresencesfromAttribute(attribute string) (map[string][]string){

  values:= db.ReadColumnsSimple("universum", "PrecenseIdsByAttributeValue" , attribute)

  var presenceMap = map[string] []string {
    "FB" : []string{},
    "TW" : []string{},
    "IN" : []string{},
  }
  for i := range values {
    fmt.Println(values[i][0:2])
    site := values[i][0:2]
    siteid := values[i][2:len(values[i])]
    presenceMap[site] = append(presenceMap[site] , siteid)  
    
  }

  fmt.Println(presenceMap)
  // mapping the string variable 
  return presenceMap
  
}