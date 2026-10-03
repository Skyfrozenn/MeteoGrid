package seismic

import (
	"context"
	"fmt"
	"sync"
	"time"
)


var seismicData = make(map[float64]float64) // свод обычной активности

var hightseismicData = make(map[float64]float64) // высокая активность



func SeismicHub(wg *sync.WaitGroup) {

	defer wg.Done()

	ctxSeismo, seismoCancel := context.WithCancel(context.Background()) // контекст для отмены сейсмо датчиков

	seismicTransfer := poolSeismic(ctxSeismo, 6) // фунция возвращающая канал

	go func() {
		time.Sleep(3 * time.Second)
		seismoCancel()
	}()
	
		
	for data := range seismicTransfer {
		for k,v := range data {
			if v > 7 {
				fmt.Println("")
				fmt.Println("Дата центр сейсмо - активности  обнаружил высокую сейсмо активность по координатам = ", k)
				hightseismicData[k] = v
				fmt.Println("")
			} else {
				seismicData[k] = v
			}
	
		}
	}

	fmt.Println("")

	if len(seismicData) > 0 {
		fmt.Println("")
		fmt.Println("Обычное сейсмо активность")
		fmt.Println("")
		for k, v := range seismicData {
			fmt.Println(k, "-", v)
		}
	} else {
		fmt.Println("")
		fmt.Println("Обычное сейсмо активность не обнаружено!")
		fmt.Println("")
	}
	 
	fmt.Println("")

	 
	if len(hightseismicData) > 0 {
		fmt.Println("")
		fmt.Println("Высокое сейсмо активность!")
		fmt.Println("")
		for k, v := range hightseismicData {
			fmt.Println(k, "-", v)
		}
	} else {
		fmt.Println("")
		fmt.Println("Высокое сейсмо активность не обнаружена!")
		fmt.Println("")
	}
	 
	fmt.Println("")
	 


}