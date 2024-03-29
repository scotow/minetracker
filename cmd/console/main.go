package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	. "github.com/scotow/minetracker"
	. "github.com/scotow/minetracker/notifier"
	. "github.com/scotow/minetracker/runner"
	. "github.com/scotow/minetracker/tracker"
)

var (
	flagHostname = flag.String("h", "127.0.0.1", "mcrcon server hostname")
	flagPort     = flag.Int("p", 25575, "mcrcon server port")
	flagPassword = flag.String("P", "", "mcrcon server password (optional)")
	flagInterval = flag.Duration("i", time.Second*30, "checking interval")
	flagExclude  = flag.String("x", "", "don't notify if this player is concerned")
	flagSilence  = flag.String("s", "", "don't notify if this player is online")
)

func main() {
	flag.Parse()

	if *flagHostname == "" || *flagPort <= 0 || *flagInterval == 0 {
		flag.Usage()
		os.Exit(1)
	}

	report := make(chan error)
	cred := Credentials{Hostname: *flagHostname, Port: *flagPort, Password: *flagPassword}
	runner, err := NewDirectRunner(cred)
	if err != nil {
		checkError(err)
	}
	server := NewServer(runner, report)

	notifier := NewConsoleNotifier()
	_ = server.Add(NewConnectionTracker(*flagExclude, *flagSilence, *flagInterval), notifier)

	checkError(<-report)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
