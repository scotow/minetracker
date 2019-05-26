package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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
	flagTitle    = flag.String("t", "Minecraft", "notigo notification title")

	flagConnInterval  = flag.Duration("i", time.Second*30, "checking interval for connections")
	flagConnSelf      = flag.String("s", "", "don't notify if this player is online or concerned")
	flagConnNotigoKey = flag.String("n", "", "notigo key(s) for connections")

	flagEntityInterval       = flag.Duration("I", time.Second*30, "checking interval for entity")
	flagEntityId             = flag.String("e", "", "entity id")
	flagEntityName           = flag.String("E", "", "entity name")
	flagEntityNotigoKey      = flag.String("N", "", "notigo key(s) for entity")
	flagEntityDiscordKey     = flag.String("d", "", "discord key for entity")
	flagEntityDiscordChannel = flag.String("c", "", "discord channel id for entity")
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

	ct := NewConnectionTracker(*flagConnSelf, *flagConnInterval)
	et := NewEntityTracker(*flagEntityId, *flagEntityName, *flagEntityInterval)

	hasTracker := false
	if cn := connectionsNotifier(); cn != nil {
		hasTracker = true
		_ = server.Add(ct, cn)
	}
	if en := entityNotifiers(); en != nil {
		hasTracker = true
		_ = server.Add(et, en)
	}

	if !hasTracker {
		fmt.Println("No tracker initiated.")
		os.Exit(1)
	}

	checkError(<-report)
}

func connectionsNotifier() Notifier {
	if *flagConnInterval <= 0 || *flagConnNotigoKey == "" {
		return nil
	}

	return parseNotigoKeys(*flagConnNotigoKey)
}

func entityNotifiers() Notifier {
	enn := entityNotigoNotifiers()
	edn := entityDiscordNotifier()

	if enn != nil && edn != nil {
		return NewMultiNotifier(enn, edn)
	}

	if enn != nil {
		return enn
	}

	return edn
}

func entityNotigoNotifiers() Notifier {
	if *flagEntityInterval <= 0 || *flagEntityId == "" || *flagEntityName == "" || *flagEntityNotigoKey == "" {
		return nil
	}

	return parseNotigoKeys(*flagEntityNotigoKey)
}

func entityDiscordNotifier() Notifier {
	if *flagEntityInterval <= 0 || *flagEntityId == "" || *flagEntityName == "" || *flagEntityDiscordKey == "" || *flagEntityDiscordChannel == "" {
		return nil
	}

	dn, err := NewDiscordNotifier(*flagEntityDiscordKey, *flagEntityDiscordChannel)
	checkError(err)

	return dn
}

func parseNotigoKeys(data string) *MultiNotifier {
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
