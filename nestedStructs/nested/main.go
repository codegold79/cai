package main

import "fmt"

// Play with nested structs.
func main() {
	type disk struct {
		dynamic bool
		format  string
	}

	type accessory struct {
		name    string
		purpose string
	}

	type devices struct {
		processor struct {
			speed float32
			cores int
		}
		memory struct {
			capacity int
			memType  string
		}
		storage disk
		addon   accessory
	}

	// Use field names in first level, mixed use in nested level.
	vivoBook := devices{
		processor: struct {
			speed float32
			cores int
		}{4.6, 4},
		memory: struct {
			capacity int
			memType  string
		}{
			capacity: 8,
			memType:  "DDR4",
		},
		storage: disk{false, "512 GB PCle SSD"},
		addon: accessory{
			name:    "ScreenPad 2.0",
			purpose: "secondary 5.65 in interactive touchscreen",
		},
	}

	// Omit field names in first level, mixed use in nested level.
	vivoBook2 := devices{
		struct {
			speed float32
			cores int
		}{4.6, 4},
		struct {
			capacity int
			memType  string
		}{
			capacity: 8,
			memType:  "DDR4",
		},
		disk{false, "512 GB PCle SSD"},
		accessory{
			purpose: "secondary 5.65 in interactive touchscreen",
			name:    "ScreenPad 2.0",
		},
	}

	// Structs can be compared. Output: true
	fmt.Println(vivoBook == vivoBook2)
}
