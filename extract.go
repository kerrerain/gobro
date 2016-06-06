package main

import (
    "bufio"
    "os"
    "fmt"
    "log"
    "strings"
    "time"
    "strconv"
    "flag"
)

type Entry struct {
	Date time.Time
	Description string
	Amount	float32
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
	fields := strings.Split(line, ";");
	return Entry{parseTime(fields[0]), fields[1], parseAmount(fields[2])}
}

func processEntries(entries []Entry) {
	var sum float32
	for _, entry := range entries {
		sum = sum + entry.Amount
	}
	fmt.Println("Sum: ", sum);
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

    processEntries(entries)

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