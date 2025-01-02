package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ccding/go-stun/stun"
	"os"
	"strings"
)

type NatResult struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Nat      string `json:"nat"`
	IP       string `json:"ip"`
	Port     uint16 `json:"port"`
	IPFamily uint16 `json:"ip_family"`
}

var returnJson *bool

func main() {
	var serverAddr = flag.String("s", "stun.qwq.pink:3478", "STUN server address")
	var ip = flag.String("i", "", "bind ip")
	var debugMode = flag.Bool("debug", false, "debug mode")
	returnJson = flag.Bool("json", false, "return json")

	flag.Parse()

	client := stun.NewClient()
	client.SetServerAddr(*serverAddr)
	client.SetLocalIP(*ip)
	if *debugMode {
		client.SetVerbose(true)
		client.SetVVerbose(true)
	}

	nat, host, err := client.Discover()
	if err != nil {
		printResult(NatResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	if host != nil {
		printResult(NatResult{
			Code:     1,
			Msg:      "ok",
			Nat:      nat.String(),
			IP:       host.IP(),
			Port:     host.Port(),
			IPFamily: host.Family(),
		})
	} else {
		printResult(NatResult{
			Code: 0,
			Msg:  "NO_HOST",
			Nat:  nat.String(),
		})
	}
}

func printResult(result NatResult) {
	if *returnJson {
		var rt, err = json.Marshal(result)
		if err != nil {
			fmt.Printf("%d|%s\n", -1, err.Error())
		} else {
			fmt.Println(string(rt))
		}
	} else {
		if result.Code == 1 {
			fmt.Printf("%d|%s|%s|%s|%d|%d\n", result.Code, result.Msg, strings.Replace(result.Nat, " ", "-", -1), result.IP, result.Port, result.IPFamily)
		} else if result.Code == 0 {
			fmt.Printf("%d|%s|%s\n", result.Code, result.Msg, strings.Replace(result.Nat, " ", "-", -1))
		} else {
			fmt.Printf("%d|%s\n", result.Code, result.Msg)
		}
	}
	os.Exit(0)
}
