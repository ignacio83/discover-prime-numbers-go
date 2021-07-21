package ui

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"ignacio83/discover-prime-numbers-go/domain"
	"io"
	"strconv"
)

const howManyMessage = "Should discovery prime numbers until: "

type DiscoverPrimeNumbersFromUserInputFactory struct {
	Writer io.Writer
	Reader io.Reader
}

func (s *DiscoverPrimeNumbersFromUserInputFactory) Build(onDiscover domain.OnDiscover) (*domain.DiscoverPrimeNumbers, error) {
	_, err := fmt.Fprintf(s.Writer, howManyMessage)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(s.Reader)
	scanner.Scan()
	qtyStr := scanner.Text()

	if len(qtyStr) == 0 {
		return nil, errors.New("end is required")
	}

	qty, err := strconv.ParseUint(qtyStr, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("End must be a number. Value: %q", qtyStr))
	}
	return domain.NewDiscoverPrimeNumbers(qty, onDiscover), nil
}
