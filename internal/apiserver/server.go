package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/metric"
	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/metric/runtimemetric"
	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/storage"
)

type server struct {
	router *http.ServeMux
	store *storage.Storage
	metric *metric.Metric
}

func newServer(store *storage.Storage, metric *metric.Metric) *server{
	s := &server{
		router: http.NewServeMux(),
		store: store,
		metric: metric,
	}
	s.configureRoute()
	return s
}

func(s * server) configureRoute(){
	listMetric := metric.NewMetric().GetMetrics()
	for i := 0; i < len(listMetric); i++ {
		s.router.HandleFunc(fmt.Sprintf("/update/%s",listMetric[i].GetNameMetric()), s.handleWtiteMetric())
		fmt.Println(fmt.Sprintf("/update/%s",listMetric[i].GetNameMetric()))
	}
	
	
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}



func(s *server) handleWtiteMetric() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Only Post requests are allowed!", http.StatusMethodNotAllowed)
        return
    }
	if r.Header.Get("Content-Type") != "text/plain"{
		http.Error(w, "Only Content-Type text/plan are allowed!", http.StatusRequestHeaderFieldsTooLarge)
        return
	}
	urlParametrs := strings.Split(r.URL.Path,"/")

	if len(urlParametrs[4]) < 1 {
		http.Error(w, "Error value ", http.StatusBadGateway)
		return
	}
		value, err := strconv.ParseUint(urlParametrs[4], 10, 64)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error value ", http.StatusBadGateway)
			return
		}
	
	if urlParametrs[2] == "gauge"{
		tmt := runtimemetric.NewGauge(urlParametrs[3])
		tmt.Update(uint64(value))
		s.store.Append(metric.NewCustomMetric(tmt))
		w.WriteHeader(http.StatusCreated)
	}else if urlParametrs[2] == "counter" {
		tmt := runtimemetric.NewCounter(urlParametrs[3])
		tmt.Update(value)
		s.store.Append(metric.NewCustomMetric(tmt))
		w.WriteHeader(http.StatusCreated)
	}
	

}

}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
