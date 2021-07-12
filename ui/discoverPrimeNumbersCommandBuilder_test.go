package ui

import (
	"bytes"
	"ignacio83/discover-prime-numbers-go/domain"
	"testing"
)

func TestBuildDiscoverPrimeNumbersCommand(t *testing.T) {
	t.Run("Should build command when input is valid", func(t *testing.T) {
		reader := bytes.Buffer{}
		writer := bytes.Buffer{}

		reader.Write([]byte("20\n"))

		builder := DiscoverPrimeNumbersCommandFromUserInputFactory{Writer: &writer, Reader: &reader}

		commandGot, _ := builder.Build()
		commandWant := &domain.DiscoverPrimeNumbersCommand{Qty: 20}

		messageGot := writer.String()
		messageWant := "Should discovery how many prime numbers? "

		if messageGot != messageWant {
			t.Fatalf("got %q want %q", messageGot, messageWant)
		}

		if commandGot.Qty != commandWant.Qty {
			t.Fatalf("got qty %d want %d", commandGot.Qty, commandWant.Qty)
		}
	})

	t.Run("Should throw error when input is blank", func(t *testing.T) {
		reader := bytes.Buffer{}
		writer := bytes.Buffer{}

		reader.Write([]byte("\n"))

		builder := DiscoverPrimeNumbersCommandFromUserInputFactory{Writer: &writer, Reader: &reader}

		_, errGot := builder.Build()

		messageGot := writer.String()
		messageWant := "Should discovery how many prime numbers? "

		if messageGot != messageWant {
			t.Fatalf("got %q want %q", messageGot, messageWant)
		}
		assertError(t, errGot, "qty is required")
	})

	t.Run("Should throw error when input is a letter", func(t *testing.T) {
		reader := bytes.Buffer{}
		writer := bytes.Buffer{}

		reader.Write([]byte("20A\n"))

		builder := DiscoverPrimeNumbersCommandFromUserInputFactory{Writer: &writer, Reader: &reader}

		_, errGot := builder.Build()

		messageGot := writer.String()
		messageWant := "Should discovery how many prime numbers? "

		if messageGot != messageWant {
			t.Fatalf("got %q want %q", messageGot, messageWant)
		}
		assertError(t, errGot, "Qty must be a number. Value: \"20A\": strconv.Atoi: parsing \"20A\": invalid syntax")
	})
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatalf("didn't get an error but wanted %v", want)
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
