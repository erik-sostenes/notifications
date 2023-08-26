package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	c "github.com/erik-sostenes/notifications-api/cmd/bootstrap/container"
	"github.com/erik-sostenes/notifications-api/pkg/server"
)

func main() {
	// Sets logger flags
	log.SetFlags(log.Flags() | log.Llongfile)

	routes, err := c.Injector()
	if err != nil {
		log.Fatal(err)
	}

	svr := server.New(*routes)

	go func() {
		if err := svr.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	svr.Server.Shutdown(ctx)

	log.Println("server has been canceled")
}
