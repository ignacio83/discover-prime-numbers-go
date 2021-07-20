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
	fmt.Fprintf(s.Writer, howManyMessage)
	scanner := bufio.NewScanner(s.Reader)
	scanner.Scan()
	qtyStr := scanner.Text()

	if len(qtyStr) == 0 {
		return nil, errors.New("qty is required")
	}

	qty, err := strconv.ParseUint(qtyStr, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Qty must be a number. Value: %q", qtyStr))
	}
	return &domain.DiscoverPrimeNumbers{Qty: qty, OnDiscover: onDiscover}, nil
}
