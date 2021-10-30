package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/ccding/go-stun/stun"
)

func main() {
	var serverAddr = flag.String("s", "stun.qwq.pink:3478", "STUN server address")
	var debugMode = flag.Bool("debug", false, "debug mode")
	var returnJson = flag.Bool("json", false, "return json")

	flag.Parse()

	client := stun.NewClient()
	client.SetServerAddr(*serverAddr)
	if *debugMode {
		client.SetVerbose(true)
		//client.SetVVerbose(true)
	}

	nat, host, err := client.Discover()
	if err != nil {
		if *returnJson {
			fmt.Print(`{"code": -1, "msg": "` + err.Error() + `"}`)
		} else {
			fmt.Print(strings.Join([]string{
				"-1",
				err.Error(),
			}, "|"))
		}
		return
	}
	if host != nil {
		if *returnJson {
			fmt.Print(`{"code": 1, "msg": "OK", "nat": "` + nat.String() + `", "ip": "` + host.IP() + `", "port": "` + strconv.Itoa(int(host.Port())) + `", "ip_family": "` + strconv.Itoa(int(host.Family())) + `"}`)
		} else {
			fmt.Print(strings.Join([]string{
				"1",
				"OK",
				strings.Replace(nat.String(), " ", "-", -1),
				host.IP(),
				strconv.Itoa(int(host.Port())),
				strconv.Itoa(int(host.Family())),
			}, "|"))
		}
	} else {
		if *returnJson {
			fmt.Print(`{"code": 0, "msg": "NO_HOST", "nat": "` + nat.String() + `"}`)
		} else {
			fmt.Print(strings.Join([]string{
				"0",
				"NO_HOST",
				strings.Replace(nat.String(), " ", "-", -1),
			}, "|"))
		}
	}
}
