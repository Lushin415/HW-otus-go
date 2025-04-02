package process

func Process(readChan <-chan int, writeChan chan<- float64) {
	defer close(writeChan)
	sum := 0
	count := 0
	arifmMid := 0.0

	for n := range readChan {
		sum += n
		count++

		if count == 10 { // Подсчитываем каждые 10 значений
			arifmMid += float64(sum) / 10
			writeChan <- arifmMid
			sum = 0
			count = 0
			arifmMid = 0.0
		}
	}
}
