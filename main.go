package main

import (
	"fmt"
	"github.com/d2r2/go-dht"
	"github.com/d2r2/go-logger"
	"github.com/stianeikeland/go-rpio/v4"
	"log"
)

func main() {
	log.Println("Opening rpio")
	err := rpio.Open()
	if err != nil {
		log.Panicf("could not open rpio: %s", err)
	}
	defer rpio.Close()

	err = logger.ChangePackageLogLevel("dht", logger.FatalLevel)
	if err != nil {
		log.Fatalf("could not turn off dht logger: %s", err)
	}

	pinNumber := 18
	tries := 10

	temperature, humidity, _, err := dht.ReadDHTxxWithRetry(dht.DHT11, int(pinNumber), false, tries)
	if err != nil {
		log.Fatalf("unable to read sensor after: %d tries err: %s", tries, err)
	}

	fmt.Printf("temperature: %0.2f°C humidity: %0.2f%%\n", temperature, humidity)
}
