package metric

import (
	"math/rand"
	"reflect"
	"runtime"

	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/metric/runtimemetric"
)

type  Metric struct {
	list       []metricType
}

func  NewMetric() *Metric{
	m := &Metric{
		list: []metricType{ 
			runtimemetric.NewGauge("Alloc"),
			runtimemetric.NewGauge("BuckHashSys"),
			runtimemetric.NewGauge("Frees"),
			runtimemetric.NewGauge("GCCPUFraction"),
			runtimemetric.NewGauge("GCSys"),
			runtimemetric.NewGauge("HeapAlloc"),
			runtimemetric.NewGauge("HeapIdle"),
			runtimemetric.NewGauge("HeapInuse"),
			runtimemetric.NewGauge("HeapObjects"),
			runtimemetric.NewGauge("HeapReleased"),
			runtimemetric.NewGauge("HeapSys"),
			runtimemetric.NewGauge("LastGC"),
			runtimemetric.NewGauge("Lookups"),
			runtimemetric.NewGauge("MCacheInuse"),
			runtimemetric.NewGauge("MCacheSys"),
			runtimemetric.NewGauge("MSpanInuse"),
			runtimemetric.NewGauge("MSpanSys"),
			runtimemetric.NewGauge("Mallocs"),
			runtimemetric.NewGauge("NextGC"),
			runtimemetric.NewGauge("NumForcedGC"),
			runtimemetric.NewGauge("NumGC"),
			runtimemetric.NewGauge("OtherSys"),
			runtimemetric.NewGauge("PauseTotalNs"),
			runtimemetric.NewGauge("StackInuse"),
			runtimemetric.NewGauge("StackSys"),
			runtimemetric.NewGauge("Sys"),
			runtimemetric.NewGauge("TotalAlloc"),
			runtimemetric.NewCounter("PollCount"),
			runtimemetric.NewGauge("RandomValue"),
		},

	}
		


	return m
}


func NewCustomMetric( mtrT metricType) *Metric{
	
	m := &Metric{
		list: []metricType{ 
			mtrT,
		},
	}
	return m 

}


func (m *Metric)Update(rtm *runtime.MemStats) {
	for i := 0; i < len(m.list); i++ {
		if m.list[i].GetName() == "RandomValue"{
			m.list[i].Update(uint64(rand.Uint32())<<32 + uint64(rand.Uint32()))
		}
		if m.list[i].GetType() == "counter" && m.list[i].GetName() != "RandomValue" {
			m.list[i].Increment()
		}

		if m.list[i].GetType() == "gauge"{
			r := reflect.ValueOf(rtm)
			f := reflect.Indirect(r).FieldByName(m.list[i].GetName())
			
			if f.IsValid(){
				if f.Type().Name() == "float64"{
					m.list[i].Update(uint64(f.Float()))
				}
				if f.Type().Name() == "uint64"{
					m.list[i].Update(f.Uint())
				}
				if f.Type().Name() == "uint32"{
					m.list[i].Update(f.Uint())
				}

	
			}
			
		
			
		}
	}

	
}

func( m *Metric) GetMetricList() []string  {
	var  array []string

	for i := 0; i < len(m.list); i++ {
		array = append(array, m.list[i].GetFullMetric())
	}

	return array
}

func (m *Metric) GetMetrics() []metricType{
	var  array []metricType

	for i := 0; i < len(m.list); i++ {
		array = append(array, m.list[i])
	}

	return array
}
