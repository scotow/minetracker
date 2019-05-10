package main

import (
	"flag"
	"fmt"
	"github.com/scotow/skyblocktracker"
	"os"
	"strings"
	"time"
)

var (
	flagHostname = flag.String("h", "", "mcrcon server hostname")
	flagPort     = flag.Int("p", 0, "mcrcon server port")
	flagPassword = flag.String("P", "", "mcrcon server password (optional)")
	flagInterval = flag.Duration("i", time.Second*30, "checking interval")
	flagSelf     = flag.String("s", "", "don't notify if this player is online")
)

func main() {
	flag.Parse()

	if *flagHostname == "" || *flagPort <= 0 || *flagInterval == 0 {
		flag.Usage()
		os.Exit(1)
	}

	cred := skyblocktracker.Credentials{Hostname: *flagHostname, Port: *flagPort, Password: *flagPassword}
	tracker := skyblocktracker.NewTracker(cred, *flagInterval, func(online, connect, disconnect []string) {
		if *flagSelf != "" && skyblocktracker.Contains(online, *flagSelf) {
			return
		}

		if len(connect) > 0 {
			fmt.Printf("%s connected.\n", formatPlayerList(connect))
		}
		if len(disconnect) > 0 {
			fmt.Printf("%s disconnected.\n", formatPlayerList(disconnect))
		}
	})

	report := make(chan error)
	err := tracker.Start(report)
	checkError(err)

	err = <-report
	checkError(err)
}

func formatPlayerList(players []string) string {
	switch len(players) {
	case 0:
		return ""
	case 1:
		return players[0]
	default:
		return fmt.Sprintf("%s and %s", strings.Join(players[0:len(players)-1], ", "), players[len(players)-1])
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
