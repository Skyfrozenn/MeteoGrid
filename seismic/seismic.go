package seismic

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func seismicSensor(
	transferSensor chan <- map[float64]float64, // канал с передачей активности и координат
	ctx context.Context, // контекст для отмены горутины
	wg *sync.WaitGroup, // вейт группа для ожидания и завершения
	numSensor int, // номер датчика

) {
	defer wg.Done()

	for {
		select {
		case <- ctx.Done():
			fmt.Println("Я датчик номер - ", numSensor, "заканчиваю свою работу!")
			return
		default:

			coordinate := 59 + rand.Float64() * 2

			fmt.Println("Я датчик номер - ", numSensor, "собираю показания влажности воздуха по координатам = ", coordinate)
			time.Sleep(1 * time.Second)

			active := 1.5 + rand.Float64() * 10 // активность

			transferSensor <- map[float64]float64{ // передача в канал мапы
				coordinate : active,
			}

			fmt.Println("Я датчик номер - ", numSensor, "передал показания в цетр!")


		}
	}
}


func poolSeismic(
	ctx context.Context, // контекст для отмены наших датчиков
	sensorCount int, // количество датчиков
) <- chan map[float64]float64 {

	seismicChan := make(chan map[float64]float64) // канал 
	
	wg := &sync.WaitGroup{} // группа для ожидания завершения горутин и закрытия канала

	for i:=1; i<=sensorCount;i++ {
		wg.Add(1)
		go seismicSensor(seismicChan, ctx, wg, i)
	}

	go func() {
		wg.Wait()
		close(seismicChan)
	}()

	return seismicChan

}