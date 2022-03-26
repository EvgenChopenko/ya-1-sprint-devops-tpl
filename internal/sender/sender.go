package sender

import (
	"fmt"
	"time"
	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/storage"
)

func ReadPush(duration time.Duration, store *storage.Storage){
	<- time.After(duration)
	count := len(store.Rows)
	if count > 0 {
		for i := 0; i < count; i++ {
			fmt.Println(i)
			if store.Rows[i] != nil {
				ok := sendMetric(store.Rows[i].GetMetricList())
				if ok {
					copy(store.Rows[i:], store.Rows[i+1:])
				
				} else{
					fmt.Println("No Send")
				
				}

			}
			
		}

	}
	


}

func sendMetric(list []string) bool  {
	fmt.Println(list)
	return true
}