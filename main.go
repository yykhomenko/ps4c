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

	// Cross
	// EventCrossPress
	// EventCrossRelease
	c.On(gods4.EventCrossPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "Cross")
		return nil
	})

	c.On(gods4.EventCrossRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "Cross")
		return nil
	})

	// Circle
	// EventCirclePress
	// EventCircleRelease

	// Square
	// EventSquarePress
	// EventSquareRelease

	// Triangle
	// EventTrianglePress
	// EventTriangleRelease

	// L1
	// EventL1Press
	// EventL1Release
	c.On(gods4.EventL1Press, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "L1")
		return nil
	})

	c.On(gods4.EventL1Release, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "L1")
		return nil
	})

	// L2
	// EventL2Press
	// EventL2Release
	c.On(gods4.EventL2Press, func(data interface{}) error {
		b := data.(byte)
		log.Printf("* Controller #1 | %-10s | state: press, level: %v\n", "L2", b)
		return nil
	})

	c.On(gods4.EventL2Release, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "L2")
		return nil
	})

	// L3
	// EventL3Press
	// EventL3Release
	c.On(gods4.EventL3Press, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "L3")
		return nil
	})

	c.On(gods4.EventL3Release, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "L3")
		return nil
	})

	// R1
	// EventR1Press
	// EventR1Release
	c.On(gods4.EventR1Press, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "R1")
		return nil
	})

	c.On(gods4.EventR1Release, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "R1")
		return nil
	})

	// R2
	// EventR2Press
	// EventR2Release
	c.On(gods4.EventR2Press, func(data interface{}) error {
		b := data.(byte)
		log.Printf("* Controller #1 | %-10s | state: press, level: %v\n", "R2", b)
		return nil
	})

	c.On(gods4.EventR2Release, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "R2")
		return nil
	})

	// R3
	// EventR3Press
	// EventR3Release
	c.On(gods4.EventR3Press, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "R3")
		return nil
	})

	c.On(gods4.EventR3Release, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "R3")
		return nil
	})

	// D-pad up
	// EventDPadUpPress
	// EventDPadUpRelease
	c.On(gods4.EventDPadUpPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "DPadUp")
		return nil
	})

	c.On(gods4.EventDPadUpRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "DPadUp")
		return nil
	})

	// D-pad down
	// EventDPadDownPress
	// EventDPadDownRelease

	// D-pad left
	// EventDPadLeftPress
	// EventDPadLeftRelease

	// D-pad right
	// EventDPadRightPress
	// EventDPadRightRelease

	// Share
	// EventSharePress
	// EventShareRelease

	// Options
	// EventOptionsPress
	// EventOptionsRelease

	// Touchpad
	// EventTouchpadSwipe
	// EventTouchpadPress
	// EventTouchpadRelease

	// PS
	// EventPSPress
	// EventPSRelease

	// Left stick
	// EventLeftStickMove
	c.On(gods4.EventLeftStickMove, func(data interface{}) error {
		stick := data.(gods4.Stick)
		log.Printf("* Controller #1 | %-10s | x: %v, y: %v\n", "LeftStick", stick.X, stick.Y)
		return nil
	})

	// // Right stick
	// EventRightStickMove
	c.On(gods4.EventRightStickMove, func(data interface{}) error {
		stick := data.(gods4.Stick)
		log.Printf("* Controller #1 | %-10s | x: %v, y: %v\n", "RightStick", stick.X, stick.Y)
		return nil
	})

	// Accelerometer
	// EventAccelerometerUpdate
	// c.On(gods4.EventAccelerometerUpdate, func(data interface{}) error {
	// 	a := data.(gods4.Accelerometer)
	// 	log.Printf("* Controller #1 | %-10s | x: %v, y: %v, z: %v\n", "Accelerometer", a.X, a.Y, a.Z)
	// 	return nil
	// })

	// Gyroscope
	// EventGyroscopeUpdate
	// c.On(gods4.EventGyroscopeUpdate, func(data interface{}) error {
	// 	g := data.(gods4.Gyroscope)
	// 	log.Printf("* Controller #1 | %-10s | roll: %v, yaw: %v, pitch: %v\n", "Gyroscope", g.Roll, g.Yaw, g.Pitch)
	// 	return nil
	// })

	// Battery
	// EventBatteryUpdate
	c.On(gods4.EventBatteryUpdate, func(data interface{}) error {
		b := data.(gods4.Battery)
		log.Printf("* Controller #1 | %-10s | capacity: %v%%, charging: %v, cable: %v\n", "Battery",
			b.Capacity,
			b.IsCharging,
			b.IsCableConnected,
		)
		return nil
	})

	err = c.Listen()
	if err != nil {
		panic(err)
	}
}

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
