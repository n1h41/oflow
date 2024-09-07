package main

import "n1h41/oflow/internal/server"

func main() {

  fiberServer := server.NewFiberServer()
  fiberServer.Run()
}
