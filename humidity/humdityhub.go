package humidity

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var humidityData = make(map[float64]float64)

var highHumidityData = make(map[float64]float64)

func HumidityHub(wg *sync.WaitGroup) {

	defer wg.Done()

	humidityCtx, humidityCancel := context.WithCancel(context.Background()) // контекст для отмены горутин

	humidityTransfer := poolHumiDity(humidityCtx, 3)

	go func() {
		time.Sleep(3 * time.Second)
		humidityCancel()
	}()

	for data := range humidityTransfer {
		for k,v := range data {
			if v >= 70 {
				fmt.Println("Центр заметил высокую влажность воздуха!")
				highHumidityData[k] = v
			} else {
				humidityData[k] = v
			}

		}
	}

	fmt.Println("")
	fmt.Println("Обычная влажность воздуха = ")
	fmt.Println("")

	for k,v := range humidityData {
		fmt.Println(k, "- ", v)
	}

	fmt.Println("")
	fmt.Println("Высокая влажность воздуха = ")
	fmt.Println("")

	for k,v := range highHumidityData {
		fmt.Println(k, "- ", v)
	}

	fmt.Println("")

	
}