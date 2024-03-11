package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/andrewwillette/sshcmd"
)

var (
	PrivateKeyPath = os.Getenv("SERVER_LOG_PRIVATE_KEY")
	IPAddress      = os.Getenv("SERVER_LOG_IP")
	User           = os.Getenv("SERVER_LOG_USER")
	LogLocation    = os.Getenv("SERVER_LOG_LOG_LOCATION")
)

func main() {
	ServerLogs()
}

func ServerLogs() {
	parseFlags()
	logs, err := sshcmd.RemoteRun(User, IPAddress, PrivateKeyPath, "cat "+LogLocation)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", logs)
}

func parseFlags() {
	var privatekey, ip, user, logLocation string
	const defaultEmpty = ""
	flag.StringVar(&privatekey, "privatekey", defaultEmpty, "absolute path to the private key")
	flag.StringVar(&ip, "ip", defaultEmpty, "IP address of the server")
	flag.StringVar(&user, "user", defaultEmpty, "user to connect to the server as")
	flag.StringVar(&logLocation, "loglocation", defaultEmpty, "absolute path to location of the log file")
	flag.Parse()
	if privatekey != "" {
		PrivateKeyPath = privatekey
	}
	if ip != "" {
		IPAddress = ip
	}
	if user != "" {
		User = user
	}
	if logLocation != "" {
		LogLocation = logLocation
	}
}
