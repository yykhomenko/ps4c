package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kpeu3i/gods4"
)

func main() {
	cs := gods4.Find()
	if len(cs) == 0 {
		panic("No connected DS4 controllers found")
	}

	c := cs[0]
	err := c.Connect()
	if err != nil {
		panic(err)
	}

	log.Printf("* Controller #1 | %-10s | name: %s, connection: %s\n", "Connect", c, c.ConnectionType())

	onProgramTerminate(c)

	c.On(gods4.EventBatteryUpdate, func(data interface{}) error {
		battery := data.(gods4.Battery)
		log.Printf("* Controller #1 | %-10s | capacity: %v%%, charging: %v, cable: %v\n",
			"Battery",
			battery.Capacity,
			battery.IsCharging,
			battery.IsCableConnected,
		)

		return nil
	})

	c.On(gods4.EventCrossPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "Cross")

		return nil
	})

	c.On(gods4.EventCrossRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "Cross")

		return nil
	})

	c.On(gods4.EventLeftStickMove, func(data interface{}) error {
		stick := data.(gods4.Stick)
		log.Printf("* Controller #1 | %-10s | x: %v, y: %v\n", "RightStick", stick.X, stick.Y)

		return nil
	})

	c.On(gods4.EventLeftStickMove, func(data interface{}) error {
		stick := data.(gods4.Stick)
		log.Printf("* Controller #1 | %-10s | x: %v, y: %v\n", "RightStick", stick.X, stick.Y)

		return nil
	})

	err = c.Listen()
	if err != nil {
		panic(err)
	}
}

// Disconnect controller when a program is terminated
func onProgramTerminate(c *gods4.Controller) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		err := c.Disconnect()
		if err != nil {
			panic(err)
		}
		log.Printf("* Controller #1 | %-10s | bye!\n", "Disconnect")
	}()
}
