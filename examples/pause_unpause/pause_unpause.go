package main

import (
	"fmt"
	"os"
	"time"

	"github.com/naposproject/go-utorrent"
)

func main() {
	c, err := utorrent.NewClient(&utorrent.Client{
		API:      "http://192.168.1.163:8085/gui",
		Username: "admin",
		Password: os.Getenv("TORRENT_PASSWORD"),
	})
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	fmt.Printf("Pausing torrent..\n")
	err = c.PauseTorrent("E4BE9E4DB876E3E3179778B03E906297BE5C8DBE")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	time.Sleep(10 * time.Second)

	fmt.Printf("Unpausing torrent..\n")
	err = c.UnPauseTorrent("E4BE9E4DB876E3E3179778B03E906297BE5C8DBE")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	os.Exit(0)
}
