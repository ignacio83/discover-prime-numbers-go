package domain

import (
	"github.com/go-playground/validator/v10"
)

type OnDiscover func(primeNumber uint64)

type DiscoverPrimeNumbers struct {
	Start      uint64
	End        uint64 `validate:"required"`
	onDiscover OnDiscover
}

func NewDiscoverPrimeNumbers(end uint64, onDiscover OnDiscover) *DiscoverPrimeNumbers {
	return &DiscoverPrimeNumbers{Start: 0, End: end, onDiscover: onDiscover}
}

func (d *DiscoverPrimeNumbers) Execute() error {
	validate := validator.New()
	err := validate.Struct(d)
	if err != nil {
		return err
	}
	var numbersToTest []uint64
	if d.End == 1 {
		numbersToTest = []uint64{1}
		fromNumbers(numbersToTest, d.onDiscover)
	} else if d.End == 2 {
		numbersToTest = []uint64{1, 2}
		fromNumbers(numbersToTest, d.onDiscover)
	} else {
		numbersToTest = []uint64{1, 2}
		var i uint64 = 3
		for i <= d.End {
			numbersToTest = append(numbersToTest, i)
			if len(numbersToTest) == 1000 || i+2 >= d.End {
				fromNumbers(numbersToTest, d.onDiscover)
				numbersToTest = []uint64{}
			}
			i = i + 2
		}
	}
	return nil
}

func fromNumbers(numbersToTest []uint64, callback OnDiscover) {
	var qtyDiscover uint64 = 0
	for _, number := range numbersToTest {
		if isPrime(number) {
			if callback != nil {
				callback(number)
			}
			qtyDiscover++
		}
	}
}

func isPrime(number uint64) bool {
	var prime = true
	var i uint64 = 2
	for i < number {
		if number%i == 0 {
			prime = false
			break
		}
		i++
	}
	return prime
}
