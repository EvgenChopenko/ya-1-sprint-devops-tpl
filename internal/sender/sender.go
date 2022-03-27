package sender

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/storage"
)

func ReadPush(duration time.Duration, store *storage.Storage){
	<- time.After(duration)
	count := len(store.Rows)
	if count > 4 {
		for i := count-1; i >= 0 ; i-- {
			if store.Rows[i] != nil {
				ok := sendMetric(store.Rows[i].GetMetricList())
				if ok {
					fmt.Println(store.Rows[i].GetMetricList())
					store.Remove(i)
					
					fmt.Println("Send")
				} else{
					fmt.Println("No Send")
				
				}

			}
			
		}

	}
	


}



func sendMetric(list []string) bool  {
data := url.Values{}
data.Set("metric", "runserver")
client := &http.Client{}
for i := 0; i < len(list); i++ {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://127.0.0.1:8080/update/%s",list[i]),  strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error Sending")
		return false
	}
	request.Header.Set("Content-Type", "text/plain") 
	request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode()))) 
	response, err := client.Do(request)
    if err != nil {
		return false
    }
    // печатаем код ответа
    if response.StatusCode != http.StatusCreated {
		return false
	}
}

	return true
}