package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatal("failed", err)
		os.Exit(1)
	}

	fmt.Println("Текущее время:", time.String())

}
