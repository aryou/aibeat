package main

import (
	"os"
	
	"github.com/elastic/beats/libbeat/beat"

	"github.com/aryou/aibeat/beater"
)

func main() {
	// fmt.Printf("Hello my first BEAT!")
	err := beat.Run("aibeat", "", beater.New())
	if err != nil {
		os.Exit(1)
	}
}
