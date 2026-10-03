package presure

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var pressureData = make(map[float64]float64)
var highPressureData = make(map[float64]float64)

func PressureHub(wg *sync.WaitGroup) {
	defer wg.Done()

	pressureCtx, pressureCancel := context.WithCancel(context.Background())

	pressureTransfer := poolPresure(pressureCtx, 2)

	go func() {
		time.Sleep(3 * time.Second)
		pressureCancel()
	}()

	for data := range pressureTransfer {
		for k, v := range data {
			if v >= 1020 {
				fmt.Println("")
				fmt.Println("Центр обнаружил высокое атмосферное давление по координатам =", k)
				highPressureData[k] = v
				fmt.Println("")
			} else {
				pressureData[k] = v
			}
		}
	}

	fmt.Println("")

	if len(pressureData) > 0 {
		fmt.Println("")
		fmt.Println("обычное атмосферное давление")
		fmt.Println("")
		for k, v := range pressureData {
			fmt.Println(k, "-", v)
		}
	} else {
		fmt.Println("")
		fmt.Println("Обычное атмосферное давление не обнаружено!")
		fmt.Println("")
	}
	 
	fmt.Println("")

	 
	if len(highPressureData) > 0 {
		fmt.Println("")
		fmt.Println("Высокое атмосферное давление")
		fmt.Println("")
		for k, v := range highPressureData {
			fmt.Println(k, "-", v)
		}
	} else {
		fmt.Println("")
		fmt.Println("Высокое атмосферное давление не обнаружено!")
		fmt.Println("")
	}
	 
	fmt.Println("")
}