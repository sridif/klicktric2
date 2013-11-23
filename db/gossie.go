/*
  See Tests at db_test.go to use the functions here.
*/

package db

import (
 "github.com/sridif/gossie/src/gossie"
 "fmt"
 "strings"
 //"bytes"
 //"os"
 "github.com/streadway/simpleuuid"
)

type Tweet struct {
  UserID string `cf:"unsure" key:"UserID" mapping:"compact" cols:"Author"`
  Author string
  
}
type Data struct {
  BrandName string `cf:"users" key:"BrandName" cols:"ColumnName"`
  ColumnName string
  Value string
}
type Column struct {
     Name      []byte
     Value     []byte
     Ttl       int32
     Timestamp int64
}

type UUID []byte
func (value UUID) String() string {
     var r []string
     var s int
     for _, size := range [5]int{4, 2, 2, 2, 6} {
     	 var v int64
	     for i := 0; i < size; i++ {
	     	   v = v << 8
		       v = v | int64(value[s+i])
		       	 }
				r = append(r, fmt.Sprintf("%0*x", size*2, v))
				  s += size
				  }
				  return strings.Join(r, "-")
}

// read columns with string comparator
func ReadColumnsSimple(ks string, cf string, key string)(values []string){

  pool, _ := gossie.NewConnectionPool([]string{"localhost:9160"}, ks, gossie.PoolOptions{Size: 50, Timeout: 3000})


  sl := &gossie.Slice{
     Count : 2000,
  }
  
  row , _ := pool.Reader().Cf(cf).Slice(sl).Get([]byte(key))
  if row == nil{
    fmt.Println("nil row")
    return
  }

/*

   dl := &gossie.Slice{
     Start : t1.Bytes(),
     End : t2.Bytes(),
     Count : 4,
     Reversed : true,
   }
*/
 
  //fmt.Println(row.Columns)
  for i := range row.Columns {

  c := &Column{
    Name:      row.Columns[i].Name,
    Value:     row.Columns[i].Value,
    Timestamp: row.Columns[i].Timestamp,
    Ttl:       row.Columns[i].Ttl,
  }

  values= append(values, string(c.Name) )}

  return
  
}
func ReadValuesSimple(ks string, cf string, key string)(values []string){

  pool, _ := gossie.NewConnectionPool([]string{"localhost:9160"}, ks, gossie.PoolOptions{Size: 50, Timeout: 3000})

  row , _ := pool.Reader().Cf(cf).Get([]byte(key))
  if row == nil{
    fmt.Println("nil row")
    return
  }
 
  //fmt.Println(row.Columns)
  for i := range row.Columns {

  c := &Column{
    Name:      row.Columns[i].Name,
    Value:     row.Columns[i].Value,
    Timestamp: row.Columns[i].Timestamp,
    Ttl:       row.Columns[i].Ttl,
  }

  values= append(values, string(c.Value) )}

  return
  
}

// read columns with uuid comparator
func ReadColumns(ks string, cf string, key string)(values []string){

  pool, _ := gossie.NewConnectionPool([]string{"localhost:9160"}, ks, gossie.PoolOptions{Size: 50, Timeout: 3000})

  row , _ := pool.Reader().Cf(cf).Get([]byte(key))
  if row == nil{
    fmt.Println("nil row")
    return
  }
 
  //fmt.Println(row.Columns)
  for i := range row.Columns {

  c := &Column{
    Name:      row.Columns[i].Name,
    Value:     row.Columns[i].Value,
    Timestamp: row.Columns[i].Timestamp,
    Ttl:       row.Columns[i].Ttl,
  }
  var testtime simpleuuid.UUID = c.Name
  /*fmt.Println("column value : ", string(c.Value))
  //tm, _ := gossie.ParseUUID(c.Name)
  //fmt.Println()
  //fmt.Println("column name : ", testtime.Time())
  var b bytes.Buffer // A Buffer needs no initialization.
  b.Write(c.Name)
  fmt.Fprintf( &b, "world")
  b.WriteTo(os.Stdout)

//b  fmt.Println("s : ", string(row.Columns[i].Name))*/
  values= append(values, testtime.String())}

  return
  
}
type Row struct {
     Key     []byte
     Columns []*Column
}
/*
type Column struct {
     Name      []byte
     Value     []byte
     Ttl       int32
     Timestamp int64
}*/
func WriteValue(ks string, cf string, key string, col string, value string){

  pool, err := gossie.NewConnectionPool([]string{"localhost:9160"},  ks, gossie.PoolOptions{Size: 50, Timeout: 3000})
 
  if err != nil {
    fmt.Printf("haha U are fucked :P \n")
    return
    //Do something. This is amazing.
  }
  
  //mapping, err := gossie.NewMapping(&Data{})

  fmt.Println(col + " writing ... ")
  //data := &Data{key, col, value }
  //row, err := mapping.Map(data)
  column := &gossie.Column{Name : []byte(col), Value : []byte(value)} 
  columns := []*gossie.Column{column}
  row := &gossie.Row{Key : []byte(key), Columns : columns} 
  err = pool.Writer().Insert(cf, row).Run()

  if err != nil {

    fmt.Printf(" Insert failed \n %v",err)
    // do something
  }

}

func Write(ks string, cf string, key string, values []string){
  
  fmt.Println(ks)
  pool, err := gossie.NewConnectionPool([]string{"localhost:9160"},  ks, gossie.PoolOptions{Size: 50, Timeout: 3000})
 
  if err != nil {
    fmt.Printf("haha U are fucked :P \n")
    return
    //Do something. This is amazing.
  }
  
  mapping, err := gossie.NewMapping(&Tweet{})

  for i := range values {
    fmt.Println(values[i] + " writing ... ")
    tweet := &Tweet{key, values[i]}
    row, err := mapping.Map(tweet)
    err = pool.Writer().Insert(cf, row).Run()
    if err != nil {
      fmt.Printf(" Insert failed \n")
      // do something
    }
  }  
  

}

func WriteNewRow(key string, values []string) {
  
  pool , err := gossie.NewConnectionPool([]string{"localhost:9160"}, "TestGossie", gossie.PoolOptions{Size: 50, Timeout: 3000})

  if err != nil {
    fmt.Printf("haha U are fucked :P \n")
    return
    //Do something. This is amazing.
  }

  mapping, err := gossie.NewMapping(&Tweet{})

  for i := range values {
    fmt.Println(values[i] + " writing ... ")
    tweet := &Tweet{key,values[i]}
    row, err := mapping.Map(tweet)
    err = pool.Writer().Insert("universum", row).Run()
    if err != nil {
      fmt.Printf(" Insert failed \n")
      // do something
    }
  }  

  //fmt.Println(mapping.Cf());
  
}

func ReadValue(ks string, cf string, key string, col string)(values []string){

   pool, _ := gossie.NewConnectionPool([]string{"localhost:9160"}, ks, gossie.PoolOptions{Size: 50, Timeout: 3000})
  row , _ := pool.Reader().Cf(cf).Get([]byte(key))
  if row == nil{
    fmt.Println("nil row")
    return
  }
  fmt.Println(row.Columns)
  for i := range row.Columns {

    if strings.Contains(string(row.Columns[i].Name),col) {
    //if string(row.Columns[i].Name) == col {
    values= append(values, string(row.Columns[i].Value))
    }
     
  }
  
  return
  
}

func SliceReadValues2(ks string, cf string, key string)(values []string){
  
   dl := &gossie.Slice{
     Count : 120,
     Reversed : true,
   }
  fmt.Println("*")

  pool, _ := gossie.NewConnectionPool([]string{"localhost:9160"}, ks, gossie.PoolOptions{Size: 50, Timeout: 3000})

  row , _ := pool.Reader().Cf(cf).Slice(dl).Get([]byte(key))

  fmt.Println("*")
  if row == nil{
    fmt.Println("nil row")
    return 
  }
  //fmt.Println(row.Columns)
  for i := range row.Columns {
    values= append(values, string(row.Columns[i].Value))     
  }
  
  return
  
}

func SliceReadValues(ks string, cf string, key string, time1 string, time2 string)(values []string){
  
   t1, _ := simpleuuid.NewString(time1)
   t2, _ := simpleuuid.NewString(time2)
   fmt.Println(t1.Bytes())

   fmt.Println("*") 
   fmt.Println( t2.Compare(t2))
   dl := &gossie.Slice{
     Start : t1.Bytes(),
     End : t2.Bytes(),
     Count : 4,
     Reversed : true,
   }
  fmt.Println("*")

  pool, _ := gossie.NewConnectionPool([]string{"localhost:9160"}, ks, gossie.PoolOptions{Size: 50, Timeout: 3000})

  row , _ := pool.Reader().Cf(cf).Slice(dl).Get([]byte(key))

  fmt.Println("*")
  if row == nil{
    fmt.Println("nil row")
    return 
  }
  //fmt.Println(row.Columns)
  for i := range row.Columns {
    values= append(values, string(row.Columns[i].Value))     
  }
  
  return
  
}

func ReadValues(ks string, cf string, key string)(values []string){

   pool, _ := gossie.NewConnectionPool([]string{"localhost:9160"}, ks, gossie.PoolOptions{Size: 50, Timeout: 3000})
  row , _ := pool.Reader().Cf(cf).Get([]byte(key))
  if row == nil{
    fmt.Println("nil row")
    return 
  }
  //fmt.Println(row.Columns)
  for i := range row.Columns {

   
    values= append(values, string(row.Columns[i].Value))
    
     
  }
  
  return
  
}

func Read(ks string, cf string, key string)(values []string){

   pool, _ := gossie.NewConnectionPool([]string{"localhost:9160"}, ks, gossie.PoolOptions{Size: 50, Timeout: 3000})
  row , _ := pool.Reader().Cf(cf).Get([]byte(key))
  if row == nil{
    fmt.Println("nil row")
    return
  }
  for i := range row.Columns {

    length := len(row.Columns[i].Name)
    values= append(values, string(row.Columns[i].Name[2:length-1]))

  }
  return
  
}
func ReadRow(key string) (values []string){
 
  pool, _ := gossie.NewConnectionPool([]string{"localhost:9160"}, "TestGossie", gossie.PoolOptions{Size: 50, Timeout: 3000})
  row , _ := pool.Reader().Cf("universum").Get([]byte(key))
  if row == nil{
    fmt.Println("nil row")
    return
  }
  for i := range row.Columns {

    length := len(row.Columns[i].Name)
    values= append(values, string(row.Columns[i].Name[2:length-1]))

  }
  return

}






