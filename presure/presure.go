package presure


import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func presureSensor(
	transferSensor chan <- map[float64]float64, 
	ctx context.Context, 
	wg *sync.WaitGroup, 
	numSensor int, 

) {
	defer wg.Done()

	for {
		select {
		case <- ctx.Done():
			fmt.Println("Я датчик номер - ", numSensor, "заканчиваю свою работу!")
			return
		default:

			coordinate := 59 + rand.Float64() * 2

			fmt.Println("Я датчик номер - ", numSensor, "собираю показания атмосферного давления по координатам = ", coordinate)
			time.Sleep(1 * time.Second)

			pressure := 1000 + rand.Float64()*40

			transferSensor <- map[float64]float64{ // передача в канал мапы
				coordinate : pressure,
			}

			fmt.Println("Я датчик номер - ", numSensor, "передал показания в цетр!")


		}
	}
}


func poolPresure(
	ctx context.Context, // контекст для отмены наших датчиков
	sensorCount int, // количество датчиков
) <- chan map[float64]float64 {

	presureChan := make(chan map[float64]float64) // канал 
	
	wg := &sync.WaitGroup{} // группа для ожидания завершения горутин и закрытия канала

	for i:=1; i<=sensorCount;i++ {
		wg.Add(1)
		go presureSensor(presureChan, ctx, wg, i)
	}

	go func() {
		wg.Wait()
		close(presureChan)
	}()

	return presureChan

}