package apiserver

import (
	"net/http"

	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/metric"
	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/storage"
)

func Start () error {
    store := storage.NewStorage()
	metric := metric.NewMetric()
	srv := newServer(store,metric)
	
	s := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: srv,
	}
	return s.ListenAndServe()
}