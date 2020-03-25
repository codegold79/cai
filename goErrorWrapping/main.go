package main

import (
	"errors"
	"fmt"
)

func main() {
	err := runApp()
	fmt.Printf("\nerr value: %v \n type: %[1]T\n", err)

	fmt.Printf("\nerr.Error: %v\n type: %[1]T\n", err.Error())

	// I thought err was supposed to unwrap until it found an error that
	// matched possErr, but the result is false. What am I missing?
	possErr := errors.New("laptop not in possession")
	fmt.Println("result of errors is:", errors.Is(err, possErr))
}

// locateStore, goToStore, and buyLaptop return both wrapped and not wrapped errors
func locateStore() locateError {
	return locateError{
		msg:"no store nearby",
		fault: errors.New("no store nearby"),
	}
}

func goToStore() error {
	err := locateStore()
	return err
}

func buyLaptop() error {
	err := goToStore()
	return err
}

// setupLaptop, writeProgram, runApp all return wrapped errors.
func setupLaptop() error {
	err := errors.New("laptop not in possession")
	return fmt.Errorf("laptop setup failed: %w", err)
}

func writeProgram() error {
	err := setupLaptop()
	return fmt.Errorf("start coding failed: %w", err)
}

func runApp() error {
	err := writeProgram()
	return fmt.Errorf("write program failed: %w", err)
}
