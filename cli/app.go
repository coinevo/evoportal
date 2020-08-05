package cli

import (
	"log"
	"net/url"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var app = kingpin.New("evoportal", "EVO DApp Server")

var evoRPC = app.Flag("evo-rpc", "URL of evo RPC service").Envar("EVO_RPC").Default("").String()

func Run() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func getEvoRPCURL() *url.URL {
	if *evoRPC == "" {
		log.Fatalln("Please set EVO_RPC to evod's RPC URL")
	}

	url, err := url.Parse(*evoRPC)
	if err != nil {
		log.Fatalln("EVO_RPC URL:", *evoRPC)
	}

	if url.User == nil {
		log.Fatalln("EVO_RPC URL (must specify user & password):", *evoRPC)
	}

	return url
}
