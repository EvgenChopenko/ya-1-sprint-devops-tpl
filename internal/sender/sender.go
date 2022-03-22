package sender

import (
	"fmt"
	"time"
	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/storage"
)

func ReadPush(duration time.Duration, store *storage.Storage, b chan bool){
	time.Sleep(duration)
	if len(store.Rows) > 0 {
		for i := 0; i < len(store.Rows); i++ {
			if store.Rows[i] != nil {
				ok := sendMetric(store.Rows[i].GetMetriList())
				if ok {
					store.Rows[i] = nil
					b <- true
				} else{
					fmt.Println("No Send")
					b <- false
				}

			}
			
		}

	}
	


}

func sendMetric(list []string) bool  {
	fmt.Println(list)
	return true
}