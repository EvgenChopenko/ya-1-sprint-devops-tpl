package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	 m "github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/metric"
	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/monitor"
	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/sender"
	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/storage"
)

func main() {
    metric := m.NewMetric() 
	interval := time.Duration(2) * time.Second
	store := storage.NewStorage()
	mon := monitor.New(interval,metric)
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	for  {
		c := make(chan *m.Metric) // Делает канал для связи
		b := make(chan bool)
		go mon.Read(c)	

		//time.Sleep(time.Second)
		go sender.ReadPush((time.Duration(10) * time.Second), store, b)
		select {
			case metric := <-c:
				store.Append(metric)
			case status := <-b:
				if status {
					fmt.Println("Metric Sended")
				}
			case <-sigc:
				fmt.Println("Close")
				return
			}
	}
}