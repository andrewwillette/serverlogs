package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

var (
	PrivateKeyPath = os.Getenv("SERVER_LOG_PRIVATE_KEY")
	IPAddress      = os.Getenv("SERVER_LOG_IP")
	User           = os.Getenv("SERVER_LOG_USER")
	LogLocation    = os.Getenv("SERVER_LOG_LOG_LOCATION")
)

func main() {
	parseFlags()
	logs, err := RemoteRun(User, IPAddress, PrivateKeyPath, "cat "+LogLocation)
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

func RemoteRun(user, ip, privateKeyFilePath, command string) (string, error) {
	privateKey, err := getFileAsString(privateKeyFilePath)
	if err != nil {
		return "", err
	}
	// run command on remote server
	key, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return "", err
	}
	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}
	client, err := ssh.Dial("tcp", net.JoinHostPort(ip, "22"), config)
	if err != nil {
		return "", err
	}
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(command)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func getFileAsString(filepath string) (string, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
