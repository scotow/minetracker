package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	. "github.com/scotow/skyblocktracker"
	. "github.com/scotow/skyblocktracker/notifier"
	. "github.com/scotow/skyblocktracker/tracker"
)

var (
	flagHostname  = flag.String("h", "", "mcrcon server hostname")
	flagPort      = flag.Int("p", 0, "mcrcon server port")
	flagPassword  = flag.String("P", "", "mcrcon server password (optional)")
	flagInterval  = flag.Duration("i", time.Second*30, "checking interval")
	flagSelf      = flag.String("s", "", "don't notify if this player is online or concerned")
	flagNotigoKey = flag.String("k", "", "notigo key")
)

func main() {
	flag.Parse()

	if *flagHostname == "" || *flagPort <= 0 || *flagInterval == 0 || *flagNotigoKey == "" {
		flag.Usage()
		os.Exit(1)
	}

	report := make(chan error)

	server := NewServer(*flagHostname, *flagPort, *flagPassword, report)
	notifier := NewNotigoNotifier(*flagNotigoKey, "Skyblock")
	_ = server.Add(NewConnectionTracker(*flagSelf), notifier, *flagInterval)
	_ = server.Add(NewEntityTracker("wandering_trader", "Wandering Trader"), notifier, *flagInterval)

	err := <-report
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
