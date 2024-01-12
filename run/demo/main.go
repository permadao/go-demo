package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/permadao/go-demo/demo"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	rdsURL := "redis://@localhost:6379/4"
	sqlDSN := "root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	d := demo.New(sqlDSN, rdsURL)
	d.Run(":8080")

	<-signals
	d.Close()
}
