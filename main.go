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
	// EventCirclePress   Event = "circle.press"
	// EventCircleRelease Event = "circle.release"
	//
	// // Square
	// EventSquarePress   Event = "square.press"
	// EventSquareRelease Event = "square.release"
	//
	// // Triangle
	// EventTrianglePress   Event = "triangle.press"
	// EventTriangleRelease Event = "triangle.release"
	//
	// // L1
	// EventL1Press   Event = "l1.press"
	// EventL1Release Event = "l1.release"
	//
	// // L2
	// EventL2Press   Event = "l2.press"
	// EventL2Release Event = "l2.release"
	//
	// // L3
	// EventL3Press   Event = "l3.press"
	// EventL3Release Event = "l3.release"
	//
	// // R1
	// EventR1Press   Event = "r1.press"
	// EventR1Release Event = "r1.release"
	//
	// // R2
	// EventR2Press   Event = "r2.press"
	// EventR2Release Event = "r2.release"
	//
	// // R3
	// EventR3Press   Event = "r3.press"
	// EventR3Release Event = "r3.release"
	//
	// // D-pad up
	// EventDPadUpPress   Event = "dpad_up.press"
	// EventDPadUpRelease Event = "dpad_up.release"
	//
	// // D-pad down
	// EventDPadDownPress   Event = "dpad_down.press"
	// EventDPadDownRelease Event = "dpad_down.release"
	//
	// // D-pad left
	// EventDPadLeftPress   Event = "dpad_left.press"
	// EventDPadLeftRelease Event = "dpad_left.release"
	//
	// // D-pad right
	// EventDPadRightPress   Event = "dpad_right.press"
	// EventDPadRightRelease Event = "dpad_right.release"
	//
	// // Share
	// EventSharePress   Event = "share.press"
	// EventShareRelease Event = "share.release"
	//
	// // Options
	// EventOptionsPress   Event = "options.press"
	// EventOptionsRelease Event = "options.release"
	//
	// // Touchpad
	// EventTouchpadSwipe   Event = "touchpad.swipe"
	// EventTouchpadPress   Event = "touchpad.press"
	// EventTouchpadRelease Event = "touchpad.release"
	//
	// // PS
	// EventPSPress   Event = "ps.press"
	// EventPSRelease Event = "ps.release"
	//
	// // Left stick
	// EventLeftStickMove Event = "left_stick.move"
	//
	// // Right stick
	// EventRightStickMove Event = "right_stick.move"
	//
	// // Accelerometer
	// EventAccelerometerUpdate Event = "accelerometer.update"
	//
	// // Gyroscope
	// EventGyroscopeUpdate Event = "gyroscope.update"
	//
	// // Battery
	// EventBatteryUpdate Event = "battery.update"

	c.On(gods4.EventBatteryUpdate, func(data interface{}) error {
		battery := data.(gods4.Battery)
		log.Printf("* Controller #1 | %-10s | capacity: %v%%, charging: %v, cable: %v\n", "Battery",
			battery.Capacity,
			battery.IsCharging,
			battery.IsCableConnected,
		)
		return nil
	})

	// c.On(gods4.EventGyroscopeUpdate, func(data interface{}) error {
	// 	g := data.(gods4.Gyroscope)
	// 	log.Printf("* Controller #1 | %-10s | roll: %v, yaw: %v, pitch: %v\n", "Gyroscope", g.Roll, g.Yaw, g.Pitch)
	// 	return nil
	// })

	// c.On(gods4.EventAccelerometerUpdate, func(data interface{}) error {
	// 	a := data.(gods4.Accelerometer)
	// 	log.Printf("* Controller #1 | %-10s | x: %v, y: %v, z: %v\n", "Accelerometer", a.X, a.Y, a.Z)
	// 	return nil
	// })

	// L1 =======================================================================
	L1(c)
	// ==========================================================================

	// L2 =======================================================================
	c.On(gods4.EventL2Press, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "L2")
		return nil
	})

	c.On(gods4.EventL2Release, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "L2")
		return nil
	})
	// ==========================================================================

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

func L1(c *gods4.Controller) {
	c.On(gods4.EventL1Press, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "L1")
		return nil
	})

	c.On(gods4.EventL1Release, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "L1")
		return nil
	})
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
