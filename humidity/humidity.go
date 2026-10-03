package humidity

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func humditySensor(
	transferSensor chan<- map[float64]float64,
	wg *sync.WaitGroup,
	ctx context.Context,
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
			fmt.Println("Я датчик номер - ", numSensor, "начинаю изменять влажность воздуха!")
			time.Sleep(1 * time.Second)

			transferSensor <- map[float64]float64{
				coordinate : 40 + rand.Float64()*40, // 40–80%,
			}

			fmt.Println("Я датчик номер - ", numSensor, "передал показания влажности воздуха по координатам - ", coordinate)

		}
	}
}

func poolHumiDity(ctx context.Context,sensorCount int) <- chan map[float64]float64 {
	humdityChan := make(chan map[float64]float64)

	wg := &sync.WaitGroup{}

	for i:=1; i<=sensorCount; i++ {
		wg.Add(1)
		go humditySensor(humdityChan, wg, ctx, i)
	}

	go func() {
		wg.Wait()
		close(humdityChan)
	}()

	return  humdityChan
}