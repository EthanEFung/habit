package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func beginningOfMonth(date time.Time)  (time.Time) {
	return date.AddDate(0, 0, -date.Day() + 1)
}

func endOfMonth(date time.Time) (time.Time) {
	return date.AddDate(0, 1, -date.Day())
}

var directory string
func init() {
	defaultFileDir := filepath.Join(os.Getenv("HOME"), "Documents")
	flag.StringVar(&directory, "dir", defaultFileDir, "default directory to save csv files")
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("cannot execute habitual without a keyword")
	}

	var task string

	for i, arg := range args {
		task += arg
		if i < len(args)-1 {
			task += "-"
		}
	}

	now := time.Now().Local()
	month := now.Month().String()
	year := strconv.Itoa(now.Year())

	// make a file for the month
	fp := filepath.Join(directory, month+"-"+year+".csv")

	f, err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)

	writer := csv.NewWriter(f)

	err = writer.Write([]string{task, now.String()})
	check(err)
	writer.Flush()
	f.Close()

	// create month view of progress

  // first create a hash set of all the time task has been done
	f, err = os.OpenFile(fp, os.O_RDONLY, 0600)
	check(err)
	defer f.Close()
	
	reader := csv.NewReader(f)

	did := make(map[time.Time]bool)
	
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		check(err)

		curr, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", record[1])

		check(err)
		if record[0] == task {
			did[curr] = true
		}
	}

	first, last := beginningOfMonth(now), endOfMonth(now)
	curr := first

  var builder strings.Builder

	n := 1
	builder.WriteString("Monthly view: ________\n")
	for curr.Equal(last) || curr.Before(last) {
		if did[curr] {
			builder.WriteString("|✅")
		} else {
			day := curr.Day()
			builder.WriteString("|")
			if day < 10 {
				builder.WriteString("0")
			}
			str := strconv.Itoa(day)
			builder.WriteString(str)
		}
		if n % 7 == 0 {
			builder.WriteString("|\n _____________________\n")
		}
		curr = curr.Add(time.Hour * 24)
		n++
	}
	builder.WriteString("|")

	fmt.Println(builder.String())
}
