package domain

type OnDiscover func(primeNumber uint64)

type DiscoverPrimeNumbers struct {
	Qty        uint64
	OnDiscover OnDiscover
}

func (d *DiscoverPrimeNumbers) Execute() {
	var numbersToTest []uint64
	if d.Qty == 1 {
		numbersToTest = []uint64{1}
		fromNumbers(numbersToTest, d.OnDiscover)
	} else if d.Qty == 2 {
		numbersToTest = []uint64{1, 2}
		fromNumbers(numbersToTest, d.OnDiscover)
	} else {
		numbersToTest = []uint64{1, 2}
		var i uint64 = 3
		for i <= d.Qty {
			numbersToTest = append(numbersToTest, i)
			if len(numbersToTest) == 1000 || i+2 >= d.Qty {
				fromNumbers(numbersToTest, d.OnDiscover)
				numbersToTest = []uint64{}
			}
			i = i + 2
		}
	}
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
