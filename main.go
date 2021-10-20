package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kpeu3i/gods4"
	"golang.org/x/net/websocket"
)

const IP = "10.0.0.216"

func main() {
	cs := gods4.Find()
	if len(cs) == 0 {
		panic("No connected DS4 controllers found")
	}

	c := cs[0]

	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}

	log.Printf("* Controller #1 | %-10s | name: %s, connection: %s\n", "Connect", c, c.ConnectionType())

	origin := "http://" + IP + "/"
	url := "ws://" + IP + ":80/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to %s", url)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals

		err := ws.Close()
		if err != nil {
			panic(err)
		}
		log.Printf("* IP %s | %-10s | bye!\n", url, "Disconnect")

		err = c.Disconnect()
		if err != nil {
			panic(err)
		}
		log.Printf("* Controller #1 | %-10s | bye!\n", "Disconnect")
	}()

	c.On(gods4.EventCrossPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "Cross")
		send(ws, "toggle")
		return nil
	})
	c.On(gods4.EventCrossRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "Cross")
		send(ws, "toggle")
		return nil
	})

	c.On(gods4.EventDPadUpPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "DPadUp")
		send(ws, "r1On")
		return nil
	})
	c.On(gods4.EventDPadUpRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "DPadUp")
		send(ws, "r1Off")
		return nil
	})

	c.On(gods4.EventDPadRightPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "DPadUp")
		send(ws, "r2On")
		return nil
	})
	c.On(gods4.EventDPadRightRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "DPadUp")
		send(ws, "r2Off")
		return nil
	})

	c.On(gods4.EventDPadDownPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "DPadUp")
		send(ws, "r3On")
		return nil
	})
	c.On(gods4.EventDPadDownRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "DPadUp")
		send(ws, "r3Off")
		return nil
	})

	c.On(gods4.EventDPadLeftPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "DPadUp")
		send(ws, "r4On")
		return nil
	})
	c.On(gods4.EventDPadLeftRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "DPadUp")
		send(ws, "r4Off")
		return nil
	})

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

	log.Fatal(c.Listen())
}

func send(ws *websocket.Conn, msg string) error {
	if _, err := ws.Write([]byte(msg)); err != nil {
		log.Fatal(err)
	}
	var buf = make([]byte, 512)
	var n int
	var err error
	if n, err = ws.Read(buf); err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Received: %s.\n", buf[:n])
	return nil
}
