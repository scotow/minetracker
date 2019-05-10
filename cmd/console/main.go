package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	. "github.com/scotow/skyblocktracker"
)

var (
	flagHostname = flag.String("h", "", "mcrcon server hostname")
	flagPort     = flag.Int("p", 0, "mcrcon server port")
	flagPassword = flag.String("P", "", "mcrcon server password (optional)")
	flagInterval = flag.Duration("i", time.Second*30, "checking interval")
	flagSelf     = flag.String("s", "", "don't notify if this player is online or concerned")
)

func main() {
	flag.Parse()

	if *flagHostname == "" || *flagPort <= 0 || *flagInterval == 0 {
		flag.Usage()
		os.Exit(1)
	}

	cred := Credentials{Hostname: *flagHostname, Port: *flagPort, Password: *flagPassword}
	tracker := NewTracker(cred, *flagInterval, onChange)

	report := make(chan error)
	err := tracker.Start(report)
	checkError(err)

	err = <-report
	checkError(err)
}

func onChange(online, connect, disconnect []string) {
	if *flagSelf != "" {
		if Contains(online, *flagSelf) {
			return
		}

		connect = Remove(connect, *flagSelf)
		disconnect = Remove(connect, *flagSelf)
	}

	if len(connect) > 0 {
		fmt.Printf("%s connected.\n", FormatPlayerList(connect))
	}
	if len(disconnect) > 0 {
		fmt.Printf("%s disconnected.\n", FormatPlayerList(disconnect))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
