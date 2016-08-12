package main

import "log"
import "mapreduce"
import "os"
import "time"

const (
	SELF = "54.201.255.6:1337"
	INPUT_FILE = "input.txt"
	CHAT_FILE = "/home/ubuntu/chat/chat.txt"
	SLEEP_TIME = 5
)

func doMapreduce() {
	mr := mapreduce.MakeMapReduce(1, 1, INPUT_FILE, SELF)
	// Wait for the MR to finish
	<-mr.DoneChannel
}

func main() {

	for {
		doMapreduce()
		// Reset the chat file for the next file
		file, err := os.Create(CHAT_FILE)
		if err != nil {
			log.Fatal("Clear chat file: ", err)
		}
		file.Close()
		time.Sleep(SLEEP_TIME * time.Second)	
	}

}
