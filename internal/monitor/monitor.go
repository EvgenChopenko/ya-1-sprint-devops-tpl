package monitor

import (
	"runtime"
	"time"
	//"fmt"
	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/metric"
)

type Monitor struct {
	tMetric *metric.Metric

	interval time.Duration
}

func New(duration time.Duration, metric *metric.Metric) *Monitor {
	
	m := &Monitor{
		
		interval: duration,
		tMetric: metric,
		

	}
	
			return m
}

func (m *Monitor) Read(c chan *metric.Metric){
	var rtm runtime.MemStats
	
		//time.AfterFunc(m.interval, func() {
		
	    time.Sleep(m.interval)

		// Read full mem stats
		runtime.ReadMemStats(&rtm)

		m.tMetric.Update(&rtm)
		
		c <- m.tMetric
	}

		
	

