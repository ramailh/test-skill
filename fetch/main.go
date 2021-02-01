package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ramailh/backend/fetch/props"
	"github.com/ramailh/backend/fetch/rest/router"
	"github.com/subosito/gotenv"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	env := ".env"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	if err := gotenv.Load(env); err != nil {
		log.Println(err)
	}

	props.Setup()
}

func main() {

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)

		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

		errs <- fmt.Errorf("%v", <-c)
	}()

	go func() {
		rtr := router.NewRouter()
		if err := rtr.Run(":" + props.Port); err != nil {
			errs <- err
		}
	}()

	log.Fatal(<-errs)
}
