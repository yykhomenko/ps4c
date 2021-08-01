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

	// Cross
	// EventCrossPress
	// EventCrossRelease
	c.On(gods4.EventCrossPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "Cross")
		toggle(ws)
		return nil
	})

	c.On(gods4.EventCrossRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "Cross")
		toggle(ws)
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

func toggle(ws *websocket.Conn) error {
	if _, err := ws.Write([]byte("toggle")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	var err error
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Received: %s.\n", msg[:n])
	return nil
}
