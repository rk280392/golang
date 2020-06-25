package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"time"
)


type Speed struct {
	Download float64 `json:"download"`
	Upload float64   `json:"upload"`
	Server struct {
		Sponsor string `json:"sponsor"`
		Name string    `json:"name"` 
		Country string `json:""country"`  
	} `json:"server"`
	Client struct {
		Ip string `json:"ip"`
	} `json:"client"`
}
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "speedtestuser"
	dbPass := "wifi123!"
	dbName := "speedtest"
	db, errdb := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if errdb != nil {
		panic(errdb.Error())
	}
	return db
}

func main() {
	start := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(start)
	cmd := exec.Command("python", "speedtest.py", "--json")
	out, err  := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
}
	var speed Speed
	err1 := json.Unmarshal(out, &speed)

	if err1 != nil {
		log.Fatal(err1)
}
	var speedtestDownloadSpeed = speed.Download / 1000000
	var speedtestUploadSpeed = speed.Upload / 1000000
	var myPublicIP = speed.Client.Ip
	var peerServer =  speed.Server.Sponsor + " " + speed.Server.Name + " " + speed.Server.Country

	db := dbConn()
	stmt, _ := db.Prepare("INSERT INTO speedtest (TimeStamp, PublicIp, Peers, UploadSpeed, DownloadSpeed) Values(?,?,?,?,?)")
	stmt.Exec(start,myPublicIP,peerServer,speedtestUploadSpeed,speedtestDownloadSpeed)
    }


