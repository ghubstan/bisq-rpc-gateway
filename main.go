package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"golang.org/x/bisq-grpc-gateway/client"
	"golang.org/x/bisq-grpc-gateway/proxy"
	"golang.org/x/bisq-grpc-gateway/server"
	"os"
	"strings"
)

type Configuration struct {
	RpcServerAddress string
	RpcProxyAddress  string
}

func getConfiguration() Configuration {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		glog.Fatal(err)
	}
	return configuration
}

func getParams() string {
	var params string = ""
	if len(os.Args) > 1 {
		var tokens bytes.Buffer
		for _, a := range os.Args[1:] {
			if strings.IndexAny(a, " ") > -1 || strings.IndexAny(a, "\t") > -1 {
				tokens.WriteString("\"" + a + "\"" + " ") // replace quotes os.Args removed
			} else {
				tokens.WriteString(a + " ")
			}
		}
		params = strings.TrimRight(tokens.String(), " ")
	}
	return params
}

func main() {
	flag.Parse()
	defer glog.Flush()

	var params string = getParams()
	if len(params) == 0 {
		glog.Error("no program arg [server | proxy | client]")
		return
	}

	config := getConfiguration()

	var program string = strings.Fields(params)[0]
	switch program {
	case "server":
		if err := server.RunServer(config.RpcServerAddress); err != nil {
			glog.Fatal(err)
		}
	case "proxy":
		if err := proxy.RunProxy(config.RpcProxyAddress, config.RpcServerAddress); err != nil {
			glog.Fatal(err)
		}
	case "client":
		clientCommand := strings.TrimLeft(strings.Replace(params, "client", "", 1), " ")
		if err := client.RunClient(config.RpcServerAddress, clientCommand); err != nil {
			glog.Fatal(err)
		}
	default:
		glog.Error("unknown program  " + program)
	}

}
