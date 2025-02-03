package sensor

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func RandomGenerator() int {
	a := big.NewInt(100000)

	// Генерируем случайное число
	n, err := rand.Int(rand.Reader, a)
	if err != nil {
		fmt.Println("Ошибка генерации случайного числа:", err)
		return 0
	}
	return int(n.Int64())
}

func ReadChan(generator func() int) <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)
		stop := time.After(1 * time.Minute)

		for {
			select {
			case c <- generator():
				time.Sleep(500 * time.Millisecond)
			case <-stop:
				return
			}
		}
	}()

	return c
}
