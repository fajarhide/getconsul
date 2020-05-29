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
	for {
		err := config.Load(".env")
  		if err != nil {
			log.Fatal("Error loading .env file")
			os.Exit(2)
  		}
		
		f, err := os.Create(fmt.Sprintf("%s", os.Getenv("KEY")))
    	if err != nil {
    	    fmt.Println(err)
    	    return
		}
		
		t := time.Now()
	
		resp, err := http.Get(fmt.Sprintf("https://%s@%s/v1/kv/%s/%s/env?token=%s",os.Getenv("BASIC_AUTH"), os.	Getenv("URL_CONSUL"), os.Getenv("APP"), os.Getenv("ENV"), os.Getenv("TOKEN")))
		if err != nil {
			log.Fatalln(err)
		}
	
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
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
    	fmt.Println(l, "bytes written successfully")
    	err = f.Close()
    	if err != nil {
    	    fmt.Println(err)
    	    return
		}
		
		i := os.Getenv("INTERVAL")
		interval, err := strconv.Atoi(i)
		fmt.Println("Updated at", t.Format("15:04:05 2006-01-02 \n"))
		time.Sleep( time.Duration(interval)  * time.Second)
	}
}	