package runtimeMetric

import (
	"fmt"
	"strconv"
)
type counter struct {
	name string
	value uint64
	typeMetrix string 
}

func NewCounter(name string) *counter{
	c := &counter{
		name: name,
		value: uint64(0),
		typeMetrix: "counter",
	}
	return c
}

func (c *counter) Increment(){
	c.value = c.value +1
}

func (c *counter) Update(value uint64){
	c.value = value
}
func(c* counter) GetName() string{
	return c.name
}
func(c* counter) GetType() string{
	return c.typeMetrix
}

func(c* counter) GetFullMetric() string{
	return fmt.Sprintf("%s/%s/%s",c.typeMetrix,c.name,strconv.FormatUint(c.value,10) )
}