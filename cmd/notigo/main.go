package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/scotow/notigo"
	. "github.com/scotow/skyblocktracker"
)

var (
	flagHostname  = flag.String("h", "", "mcrcon server hostname")
	flagPort      = flag.Int("p", 0, "mcrcon server port")
	flagPassword  = flag.String("P", "", "mcrcon server password (optional)")
	flagInterval  = flag.Duration("i", time.Second*30, "checking interval")
	flagSelf      = flag.String("s", "", "don't notify if this player is online")
	flagNotigoKey = flag.String("k", "", "notigo key")
)

var (
	key notigo.Key
)

func main() {
	flag.Parse()

	if *flagHostname == "" || *flagPort <= 0 || *flagInterval == 0 || *flagNotigoKey == "" {
		flag.Usage()
		os.Exit(1)
	}

	key = notigo.Key(*flagNotigoKey)

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

	var lines []string
	if len(connect) > 0 {
		lines = append(lines, fmt.Sprintf("%s connected.", FormatPlayerList(connect)))
	}
	if len(disconnect) > 0 {
		lines = append(lines, fmt.Sprintf("%s disconnected.", FormatPlayerList(disconnect)))
	}

	if len(lines) > 0 {
		err := key.Send(notigo.NewNotification("Skyblock", strings.Join(lines, "\n")))
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
