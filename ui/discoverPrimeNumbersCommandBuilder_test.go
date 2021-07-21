package ui

import (
	"bytes"
	"ignacio83/discover-prime-numbers-go/domain"
	"testing"
)

func TestBuildDiscoverPrimeNumbers(t *testing.T) {
	t.Run("Should build when input is valid", func(t *testing.T) {
		reader := bytes.Buffer{}
		writer := bytes.Buffer{}

		reader.Write([]byte("20\n"))

		builder := DiscoverPrimeNumbersFromUserInputFactory{Writer: &writer, Reader: &reader}

		var callbackGot uint64
		callback := func(primeNumber uint64) {
			callbackGot = primeNumber
		}

		commandGot, _ := builder.Build(callback)
		commandWant := domain.NewDiscoverPrimeNumbers(20, callback)

		messageGot := writer.String()
		messageWant := "Should discovery prime numbers until: "

		if messageGot != messageWant {
			t.Fatalf("got %q want %q", messageGot, messageWant)
		}

		if commandGot.End != commandWant.End {
			t.Fatalf("got end %d want %d", commandGot.End, commandWant.End)
		}

		var callbackWant uint64 = 10
		callback(callbackWant)

		if callbackGot != 10 {
			t.Fatalf("callback return %d want %d", callbackGot, callbackWant)
		}

	})

	t.Run("Should throw error when end input is blank", func(t *testing.T) {
		reader := bytes.Buffer{}
		writer := bytes.Buffer{}

		reader.Write([]byte("\n"))

		builder := DiscoverPrimeNumbersFromUserInputFactory{Writer: &writer, Reader: &reader}

		_, errGot := builder.Build(func(primeNumber uint64) {
			// Nothing
		})

		messageGot := writer.String()
		messageWant := "Should discovery prime numbers until: "

		if messageGot != messageWant {
			t.Fatalf("got %q want %q", messageGot, messageWant)
		}
		assertError(t, errGot, "end is required")
	})

	t.Run("Should throw error when input is a letter", func(t *testing.T) {
		reader := bytes.Buffer{}
		writer := bytes.Buffer{}

		reader.Write([]byte("20A\n"))

		builder := DiscoverPrimeNumbersFromUserInputFactory{Writer: &writer, Reader: &reader}

		_, errGot := builder.Build(func(primeNumber uint64) {
			// Nothing
		})

		messageGot := writer.String()
		messageWant := "Should discovery prime numbers until: "

		if messageGot != messageWant {
			t.Fatalf("got %q want %q", messageGot, messageWant)
		}
		assertError(t, errGot, "End must be a number. Value: \"20A\": strconv.ParseUint: parsing \"20A\": invalid syntax")
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
