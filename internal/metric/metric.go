package metric

import (
	"math/rand"
	"reflect"
	"runtime"

	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/metric/runtimeMetric"
)

type  Metric struct {
	list       []metricType
}

func  NewMetric() *Metric{
	m := &Metric{
		list: []metricType{ 
			runtimeMetric.NewGauge("Alloc"),
			runtimeMetric.NewGauge("BuckHashSys"),
			runtimeMetric.NewGauge("Frees"),
			runtimeMetric.NewGauge("GCCPUFraction"),
			runtimeMetric.NewGauge("GCSys"),
			runtimeMetric.NewGauge("HeapAlloc"),
			runtimeMetric.NewGauge("HeapIdle"),
			runtimeMetric.NewGauge("HeapInuse"),
			runtimeMetric.NewGauge("HeapObjects"),
			runtimeMetric.NewGauge("HeapReleased"),
			runtimeMetric.NewGauge("HeapSys"),
			runtimeMetric.NewGauge("LastGC"),
			runtimeMetric.NewGauge("Lookups"),
			runtimeMetric.NewGauge("MCacheInuse"),
			runtimeMetric.NewGauge("MCacheSys"),
			runtimeMetric.NewGauge("MSpanInuse"),
			runtimeMetric.NewGauge("MSpanSys"),
			runtimeMetric.NewGauge("Mallocs"),
			runtimeMetric.NewGauge("NextGC"),
			runtimeMetric.NewGauge("NumForcedGC"),
			runtimeMetric.NewGauge("NumGC"),
			runtimeMetric.NewGauge("OtherSys"),
			runtimeMetric.NewGauge("PauseTotalNs"),
			runtimeMetric.NewGauge("StackInuse"),
			runtimeMetric.NewGauge("StackSys"),
			runtimeMetric.NewGauge("Sys"),
			runtimeMetric.NewGauge("TotalAlloc"),
			runtimeMetric.NewCounter("PollCount"),
			runtimeMetric.NewGauge("RandomValue"),
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

func( m *Metric) GetMetriList() []string  {
	var  array []string

	for i := 0; i < len(m.list); i++ {
		array = append(array, m.list[i].GetFullMetric())
	}

	return array
}