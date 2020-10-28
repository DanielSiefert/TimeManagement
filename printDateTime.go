package main

import (
	"os"
	"time"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	
	textFile, err := os.Create("datetime.txt")
	check(err)
	currentTime := time.Now()
	dateTimeString := ("Current Time: " + currentTime.String() + "/n")
	textFile.WriteString(dateTimeString)
    check(err)
}
