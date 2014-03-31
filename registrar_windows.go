package main

import (
  "encoding/json"
  "log"
  "os"
)

func WriteRegistry(state map[string]*FileState, path string) {
  tmp := path + ".new"
  file, err := os.Create(tmp)
  if err != nil {
    log.Printf("Failed to open .logstash-forwarder.new for writing: %s\n", err)
    return
  }

  encoder := json.NewEncoder(file)
  encoder.Encode(state)
  file.Close()

  old := path + ".old"
  os.Rename(path, old)


  if err != nil {
    log.Printf("Warning, could not rename %s to %s: %s\n", path, old, err)
  }

  os.Rename(tmp, path)
  
  if err != nil {
    log.Printf("Warning, could not rename %s to %s: %s\n", path, old, err)
  }
}
