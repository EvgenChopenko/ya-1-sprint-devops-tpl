package main

import	(
 "log"
 "github.com/EvgenChopenko/ya-1-sprint-devops-tpl/internal/apiserver"
)
func main() {

	if err := apiserver.Start(); err != nil {
		log.Fatal(err)
	} 
}
