package storage

import (
	"fmt"

	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/metric"
)

type Storage struct {
	Rows []*metric.Metric
}

func NewStorage() (*Storage) {
	s := &Storage{}
	return s
}

func (s *Storage) Append(row *metric.Metric){
	s.Rows = append(s.Rows, row)
	fmt.Println("add metric")
}


func (s *Storage) Remove (i int) {
	
    s.Rows[i] = s.Rows[len(s.Rows)-1]
    s.Rows =  s.Rows[:len(s.Rows)-1]
}