package api

import (
  "github.com/stretchrcom/goweb"
 
  "net/http"
  "log"  
  "time"
  "net"
  "fmt"
)

const (
  Address string = ":9925"
)

func StartService() {

  s := &http.Server{
    Addr:           Address,
    Handler:        goweb.DefaultHttpHandler(),
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }

  listener, _ := net.Listen("tcp", Address)

  synccontroller := new(SyncController)
  metercontroller := new(MeterController)  
  rewardcontroller := new(RewardController)
  profilecontroller := new(ProfileController)
  sensorcontroller := new(SensorController)
  badgecontroller := new(BadgeController) 

  goweb.MapController(synccontroller)
  goweb.MapController(metercontroller)
  goweb.MapController(sensorcontroller)
  goweb.MapController(badgecontroller)
  goweb.MapController(profilecontroller)
  goweb.MapController(rewardcontroller)

  fmt.Println("wats up") 
  log.Fatalf("Error in Serve: %s", s.Serve(listener))

}

func Documentation(){
/*


*/
}
