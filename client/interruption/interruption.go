package interruption

import (
	"log"
	"os"
	"os/signal"
)

func GetInterruption() chan os.Signal {
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	return interrupt
}

func ListenInterruption(cInterrupt chan os.Signal) {
	select {
	case <-cInterrupt:
		log.Panic("Caught interrupt signal (interruption.ListenInterruption())")
		return
	}
}
