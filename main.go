package main

import (
    "fmt"
	"time"
	"strconv"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"encoding/base64"
	"os"
	config "github.com/joho/godotenv"
)

type Data struct {
	Value string
}

func main() {
	err := config.Load(".env")
  		if err != nil {
			log.Fatal("Error loading .env file")
			os.Exit(2)
		}

	for {
		get()
		}
}

func get() {
	resp, err := http.Get(fmt.Sprintf("https://%s@%s/v1/kv/%s/%s/env?token=%s",os.Getenv("BASIC_AUTH"), os.	Getenv("URL_CONSUL"), os.Getenv("APP"), os.Getenv("ENV"), os.Getenv("TOKEN")))
	if err != nil {
		fmt.Println(err)
		return
	}
	
	f, err := os.Create(fmt.Sprintf("%s", os.Getenv("KEY")))
	if err != nil {
		fmt.Println(err)
		return
	}
	
	t := time.Now()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []Data
	json.Unmarshal([]byte(body), &data)
	uDec, err := base64.StdEncoding.DecodeString(string(data[0].Value))
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return
	}

	l, err := f.WriteString(string(uDec))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	
	fmt.Println("Updated from consul - -", t.Format("[02/Jan/2006:15:04:05 +0700]"))
	fmt.Println(l, "bytes written value successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	i := os.Getenv("INTERVAL")
	interval, err := strconv.Atoi(i)
	time.Sleep( time.Duration(interval) * time.Second)
}