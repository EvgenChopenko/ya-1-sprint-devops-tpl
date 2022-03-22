package main
 
import (
    "fmt"
    "time"
)
 
func main() {
    c := make(chan int) // Делает канал для связи

    timeout := time.After(100 * time.Second)
    for {
        go sleepyGopher(1, c)
        go sleepyGopher(10, c)
        select { // Оператор select
            case gopherID := <-c: // Ждет, когда проснется гофер
                fmt.Println("gopher ", gopherID, " has finished sleeping")
            case <-timeout: // Ждет окончания времени
                fmt.Println("my patience ran out")
                return // Сдается и возвращается
            }
    }
}
 
func sleepyGopher(id int, c chan int) { // Объявляет канал как аргумент
    time.Sleep(3 * time.Second)
    fmt.Println("... ", id, " snore ...")
    c <- id // Отправляет значение обратно к main
}


func sleepyGopher2(id int, c chan int) { // Объявляет канал как аргумент
    time.Sleep(10 * time.Second)
    fmt.Println("... ", id, " snore ...")
    c <- id // Отправляет значение обратно к main
}