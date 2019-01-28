package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func init() {
	register()
	loadConsumer()
	loadPublisher()
}

func main() {

	if os.Args[1] == "publish" {
		fmt.Println("Publishing task...")
		tn := os.Args[2]
		arg1, _ := strconv.ParseFloat(os.Args[3], 64)
		arg2, _ := strconv.ParseFloat(os.Args[4], 64)
		publish(tn, arg1, arg2) //publishing task with task name & the arguments from the command line
	}

	if os.Args[1] == "consume" {
		fmt.Println("Consuming task...")
		consume() //consuming the tasks from the queue
	}

}

func redis() string {
	u := url.URL{
		Scheme: "redis",
		Host:   "127.0.0.1:6379",
		Path:   "0",
	}
	return u.String()
}
