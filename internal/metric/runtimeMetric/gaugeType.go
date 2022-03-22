package runtimeMetric

import (
	"fmt"
	"strconv"
)

type gauge struct {
	name string
	value  uint64
	typeMetrix string 
	
}

func NewGauge(name string) *gauge{
	g := &gauge{
		name: name,
		value: uint64(0),
		typeMetrix: "gauge",
	}
	return g
}

func (g *gauge) Update(value uint64){
	g.value = value
}

func (g *gauge) Increment(){
	g.value = g.value +1
}

func(g* gauge) GetName() string{
	return g.name
}

func(g* gauge) GetType() string{
	return g.typeMetrix
}

func(g* gauge) GetFullMetric() string{
	return fmt.Sprintf("%s/%s/%s",g.typeMetrix,g.name,strconv.FormatUint(g.value,10) )
}