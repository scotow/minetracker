package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/scotow/skyblocktracker"
	. "github.com/scotow/skyblocktracker/notifier"
	. "github.com/scotow/skyblocktracker/tracker"
)

var (
	flagHostname = flag.String("h", "", "mcrcon server hostname")
	flagPort     = flag.Int("p", 0, "mcrcon server port")
	flagPassword = flag.String("P", "", "mcrcon server password (optional)")
	flagTitle    = flag.String("t", "Minecraft", "notigo notification title")

	flagConnInterval = flag.Duration("i", time.Second*30, "checking interval for connections")
	flagConnSelf     = flag.String("s", "", "don't notify if this player is online or concerned")
	flagConnKey      = flag.String("k", "", "notigo key(s) for connections")

	flagEntityInterval = flag.Duration("I", time.Second*30, "checking interval for entity")
	flagWait           = flag.Duration("w", time.Hour, "waiting interval")
	flagEntityId       = flag.String("e", "", "entity id")
	flagEntityName     = flag.String("E", "", "entity name")
	flagEntityKey      = flag.String("K", "", "notigo key(s) for entity")
)

func main() {
	flag.Parse()

	if *flagHostname == "" || *flagPort <= 0 {
		flag.Usage()
		os.Exit(1)
	}

	report := make(chan error)
	server := NewServer(*flagHostname, *flagPort, *flagPassword, report)

	hasTracker := false
	if *flagConnInterval > 0 && *flagConnKey != "" {
		connNotifier := NewNotigoNotifier(*flagConnKey, *flagTitle)
		_ = server.Add(NewConnectionTracker(*flagConnSelf, *flagConnInterval), connNotifier)
		hasTracker = true
	}
	if *flagEntityInterval > 0 && *flagWait > 0 && *flagEntityId != "" && *flagEntityName != "" && *flagEntityKey != "" {
		keys := strings.Split(*flagEntityKey, ",")
		notifiers := make([]Notifier, len(keys))
		for i, v := range keys {
			notifiers[i] = NewNotigoNotifier(v, *flagTitle)
		}

		entityNotifier := NewMultiNotifier(notifiers...)
		if hasTracker {
			time.Sleep(time.Second * 5)
		} else {
			hasTracker = true
		}
		_ = server.Add(NewEntityTracker(*flagEntityId, *flagEntityName, *flagEntityInterval, *flagWait), entityNotifier)
	}

	if !hasTracker {
		fmt.Println("No tracker initiated.")
		os.Exit(1)
	}

	err := <-report
	fmt.Println(err.Error())
	os.Exit(1)
}
