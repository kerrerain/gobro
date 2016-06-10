package main

import (
	"bufio"
	"flag"
	//	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Date        time.Time
	Description string
	Amount      float32
}

func parseTime(input string) time.Time {
	time, err := time.Parse("02/01/2006", input)
	if err != nil {
		log.Fatal(err)
	}
	return time
}

func parseAmount(amount string) float32 {
	amount = strings.Replace(amount, ",", ".", 1)
	amountFloat, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		log.Fatal(err)
	}
	return float32(amountFloat)
}

func processLine(line string) Entry {
	fields := strings.Split(line, ";")
	return Entry{parseTime(fields[0]), fields[1], parseAmount(fields[2])}
}

// As seen in http://blog.mongodb.org/post/80579086742/running-mongodb-queries-concurrently-with-go
func createDBSession() *mgo.Session {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Timeout:  60 * time.Second,
		Database: "expenses-analyzer"}
	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	return mongoSession
}

func importEntriesIntoDB(mongoSession *mgo.Session, entries []Entry) {
	expenses := mongoSession.DB("").C("expenses")
	for _, entry := range entries {
		expenses.Insert(entry)
	}
}

func processFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	countLine := 0
	var entries []Entry

	for scanner.Scan() {
		if countLine > 7 {
			entries = append(entries, processLine(scanner.Text()))
		}
		countLine++
	}

	mongoSession := createDBSession()

	importEntriesIntoDB(mongoSession, entries)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseFileName() string {
	fileFlag := flag.String("f", "", "The path to the file")
	flag.Parse()
	return *fileFlag
}

func main() {
	file, err := os.Open(parseFileName())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	processFile(file)
}
