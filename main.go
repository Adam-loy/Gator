package main

import (
  "fmt"
  "log"
  "github.con/Adam-loy/Gator/internal/config"
)

func main() {
  cfg, err := config.Read()
  if err != nil {
    log.Fatalf("error reading config: %v",err)
  }
  fmt.Printf("Read config: %+v\n",cfg)

  err = cfg.SetUser("lane")

  cfg, err = config.Read()
  if err != nil {
    log.Fatalf("error reading config: %v",err)
  }
  fmt.Printf("Read config again: %+v\n",cfg)
}