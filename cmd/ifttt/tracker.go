package ifttt

import (
	"fmt"
	"github.com/scotow/skyblocktracker"
	"log"
)

func main() {
	online, err := skyblocktracker.OnlinePlayers("127.0.0.1", 25575, "a86a9c22f37baa6727979708b172d7c1")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(online)
}
