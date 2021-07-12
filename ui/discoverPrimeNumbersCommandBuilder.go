package ui

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"ignacio83/discover-prime-numbers-go/domain"
	"io"
	"strconv"
)

const howManyMessage = "Should discovery how many prime numbers? "

type DiscoverPrimeNumbersCommandFromUserInputFactory struct {
	Writer io.Writer
	Reader io.Reader
}

func (s *DiscoverPrimeNumbersCommandFromUserInputFactory) Build() (*domain.DiscoverPrimeNumbersCommand, error) {
	fmt.Fprintf(s.Writer, howManyMessage)
	scanner := bufio.NewScanner(s.Reader)
	scanner.Scan()
	qtyStr := scanner.Text()

	if len(qtyStr) == 0 {
		return nil, errors.New("qty is required")
	}

	qty, err := strconv.Atoi(qtyStr)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Qty must be a number. Value: %q", qtyStr))
	}
	return &domain.DiscoverPrimeNumbersCommand{Qty: qty}, nil
}
