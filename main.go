package main

import (
	"MeteoGrid/humidity"
	"MeteoGrid/presure"
	"MeteoGrid/seismic"
	"fmt"
	"sync"
)

 


func main() {
	wg := &sync.WaitGroup{}

	wg.Add(3)
	go seismic.SeismicHub(wg)
	go humidity.HumidityHub(wg)
	go presure.PressureHub(wg)

	wg.Wait()

	fmt.Println("Сейсмо центр закончил свое исследование")
	fmt.Println("Центр измерений влажности воздуха закончил свои иследования")
	fmt.Println("Центр измерения атмосферного давления закончил свои иследования")
}