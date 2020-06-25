package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func getPeerServer(server interface{}) string {
	var peerServer strings.Builder

	peerServer.WriteString(fmt.Sprintf("%v, ", server.(map[string]interface{})["sponsor"]))
	peerServer.WriteString(fmt.Sprintf("%v, ", server.(map[string]interface{})["name"]))
	peerServer.WriteString(fmt.Sprintf("%v", server.(map[string]interface{})["country"]))

	return peerServer.String()
}

func main() {
	cmd := exec.Command("python", "speedtest.py", "--json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	var speed map[string]interface{}
	json.Unmarshal([]byte(out), &speed)
	var speedtestDownloadSpeed = speed["download"]
	var speedtestUploadSpeed = speed["upload"]
	var myPublicIP = speed["client"].(map[string]interface{})["ip"]
	var peerServer = getPeerServer(speed["server"])

	fmt.Println(speedtestDownloadSpeed)
	fmt.Println(speedtestUploadSpeed)
	fmt.Println(myPublicIP)
	fmt.Println(peerServer)


}
