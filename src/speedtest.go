package main

import (
	"fmt"
	"log"
	"os/exec"
	"encoding/json"
)

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
	var peerServer = speed["server"].(map[string]interface{})["sponsor"] + speed["server"].(map[string]interface{})["name"] + speed["server"].(map[string]interface{})["country"]

	fmt.Println(speedtestDownloadSpeed)
	fmt.Println(speedtestUploadSpeed)
	fmt.Println(myPublicIP)
	fmt.Println(peerServer)

}

