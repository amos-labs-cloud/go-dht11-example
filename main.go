package main

import (
	"github.com/d2r2/go-dht"
	"github.com/d2r2/go-logger"
	"github.com/stianeikeland/go-rpio/v4"
	"log"
)

func main() {
	log.Println("Opening rpio")
	err := rpio.Open() // Open up our access to the GPIO pins (requires sudo)
	if err != nil {
		log.Panicf("could not open rpio: %s", err)
	}
	defer rpio.Close()

	err = logger.ChangePackageLogLevel("dht", logger.FatalLevel) // This turns off very noisy default logging
	if err != nil {                                              // in the dht-11 library we are using
		log.Fatalf("could not turn off dht logger: %s", err)
	}

	pinNumber := 18 // This is the pin that we are going to be interacting with, it is what we connected the sensor to
	tries := 10     // The reading of the sensor is unreliable, and the library has retries built in

	temperature, humidity, _, err := dht.ReadDHTxxWithRetry(dht.DHT11, int(pinNumber), false, tries)
	if err != nil {
		log.Fatalf("unable to read sensor after: %d tries err: %s", tries, err)
	}

	log.Printf("temperature: %0.2f°C humidity: %0.2f%%\n", temperature, humidity) // Print the information we gathered
}
