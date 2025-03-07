package main

import (
	"fmt"

	"github.com/vladimirvivien/automi/collectors"
	"github.com/vladimirvivien/automi/emitters"
	"github.com/vladimirvivien/automi/stream"
)

func main() {
	type log struct{ Event, Src, Device, Result string }
	data := emitters.Slice([]log{
		log{Event: "request", Src: "/i/a", Device: "00:11:51:AA", Result: "accepted"},
		log{Event: "response", Src: "/i/a/", Device: "00:11:51:AA", Result: "served"},
		log{Event: "request", Src: "/i/b", Device: "00:11:22:33", Result: "accepted"},
		log{Event: "response", Src: "/i/b", Device: "00:11:22:33", Result: "served"},
		log{Event: "request", Src: "/i/c", Device: "00:11:51:AA", Result: "accepted"},
		log{Event: "response", Src: "/i/c", Device: "00:11:51:AA", Result: "served"},
		log{Event: "request", Src: "/i/d", Device: "00:BB:22:DD", Result: "accepted"},
		log{Event: "response", Src: "/i/d", Device: "00:BB:22:DD", Result: "served"},
	})

	stream := stream.New(data)

	stream.Filter(func(e log) bool {
		return (e.Event == "response")
	})

	// sort returns []log
	stream.Batch().SortByName("Src")

	var logs []log
	stream.Into(collectors.Func(func(data interface{}) error {
		logs = data.([]log)
		return nil
	}))

	// open the stream
	if err := <-stream.Open(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("in logs")
	for _, log := range logs {
		fmt.Printf("%v\n", log)
	}
}
