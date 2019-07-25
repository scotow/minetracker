package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/scotow/minetracker"
	. "github.com/scotow/minetracker/notifier"
	. "github.com/scotow/minetracker/notifier/notigo"
	. "github.com/scotow/minetracker/runner"
	. "github.com/scotow/minetracker/tracker"
)

var (
	flagHostname = flag.String("h", "127.0.0.1", "mcrcon server hostname")
	flagPort     = flag.Int("p", 25575, "mcrcon server port")
	flagPassword = flag.String("P", "", "mcrcon server password (optional)")
	flagTitle    = flag.String("t", "Minecraft", "notigo notification title")

	flagConnInterval = flag.Duration("i", time.Second*30, "checking interval for connections")
	flagConnExclude  = flag.String("x", "", "don't notify if this player is concerned")
	flagConnSilence  = flag.String("s", "", "don't notify if this player is online")
	flagConnKey      = flag.String("k", "", "notigo key(s) for connections")

	flagEntityInterval = flag.Duration("I", time.Second*30, "checking interval for entity")
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
	cred := Credentials{Hostname: *flagHostname, Port: *flagPort, Password: *flagPassword}
	runner, err := NewDirectRunner(cred)
	if err != nil {
		checkError(err)
	}
	server := NewServer(runner, report)

	hasTracker := false
	if *flagConnInterval > 0 && *flagConnKey != "" {
		hasTracker = true
		_ = server.Add(NewConnectionTracker(*flagConnExclude, *flagConnSilence, *flagConnInterval), parseKeys(*flagConnKey))
	}
	if *flagEntityInterval > 0 && *flagEntityId != "" && *flagEntityName != "" && *flagEntityKey != "" {
		hasTracker = true
		_ = server.Add(NewEntityTracker(*flagEntityId, *flagEntityName, *flagEntityInterval), parseKeys(*flagEntityKey))
	}

	if !hasTracker {
		fmt.Println("No tracker initiated.")
		os.Exit(1)
	}

	checkError(<-report)
}

func parseKeys(data string) *MultiNotifier {
	keys := strings.Split(data, ",")
	notifiers := make([]Notifier, len(keys))
	for i, v := range keys {
		notifiers[i] = NewNotigoNotifier(v, *flagTitle)
	}

	return NewMultiNotifier(notifiers...)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
