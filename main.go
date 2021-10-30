package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/ccding/go-stun/stun"
)

func main() {
	var serverAddr = flag.String("s", "stun.qwq.pink:3478", "STUN server address")
	var debugMode = flag.Bool("debug", false, "debug mode")

	flag.Parse()

	client := stun.NewClient()
	client.SetServerAddr(*serverAddr)
	if *debugMode {
		client.SetVerbose(true)
		//client.SetVVerbose(true)
	}

	nat, host, err := client.Discover()
	if err != nil {
		fmt.Print(`{"code": -1, "msg": "` + err.Error() + `"}`)
		return
	}
	if host != nil {
		fmt.Print(`{"code": 1, "msg": "OK", "nat": "` + nat.String() + `", "ip": "` + host.IP() + `", "port": "` + strconv.Itoa(int(host.Port())) + `", "ip_family": "` + strconv.Itoa(int(host.Family())) + `"}`)
	}else{
		fmt.Print(`{"code": 0, "msg": "NO_HOST", "nat": "` + nat.String() + `"}`)
	}
}
