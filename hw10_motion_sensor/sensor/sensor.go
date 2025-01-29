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
		for i := 0; i < 60; i++ {
			c <- generator()
			time.Sleep(time.Second)
		}
	}()
	return c
}
