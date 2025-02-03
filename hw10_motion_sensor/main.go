package main

import (
	"fmt"

	"github.com/Lushin415/HW-otus-go/hw10_motion_sensor/process"
	"github.com/Lushin415/HW-otus-go/hw10_motion_sensor/sensor"
)

func main() {
	numChan := sensor.ReadChan(sensor.RandomGenerator)
	dataChan := make(chan float64)

	go process.Process(numChan, dataChan)

	for processedData := range dataChan {
		fmt.Printf("Среднее арифметическое 10 чисел: %.2f\n", processedData)
	}
}
