package domain

type DiscoverPrimeNumbersCommand struct {
	Qty int
}

const maxNumberToTest = 1000

func (d *DiscoverPrimeNumbersCommand) Execute() []int {
	var numbersToTest []int
	for i := 1; i < maxNumberToTest; i++ {
		numbersToTest = append(numbersToTest, i)
	}
	return fromNumbers(numbersToTest, d.Qty)
}

func fromNumbers(numbers []int, qty int) []int {
	var primes []int
	for _, number := range numbers {
		if isPrime(number) {
			primes = append(primes, number)
		}
		if len(primes) == qty {
			break
		}
	}
	return primes
}

func isPrime(number int) bool {
	var prime = true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			prime = false
			break
		}
	}
	return prime
}
