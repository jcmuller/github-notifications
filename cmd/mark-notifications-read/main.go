package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jcmuller/github-notifications/internal/utils"
)

func parseDuration(duration string) (t time.Time, err error) {
	// Ensure it's a negative duration
	if strings.Index(duration, "-") == -1 {
		duration = fmt.Sprintf("-%s", duration)
	}
	d, err := time.ParseDuration(duration)
	if err != nil {
		return
	}

	return time.Now().Add(d), nil
}

func getLimit() (t time.Time, err error) {
	var duration string
	var timestamp string

	flag.StringVar(&duration, "duration", "24h", "Time expressed in duration. 15m, 1h, ...")
	flag.StringVar(&timestamp, "timestamp", "", "Timestamp as of when to clear notification")
	flag.Parse()

	if len(timestamp) == 0 {
		t, err = parseDuration(duration)

		if err != nil {
			log.Fatalf("Could not parse duration: %v\n", err)
		}
	} else {
		t, err = time.Parse("2006-01-02T15:04", timestamp)
		if err != nil {
			log.Fatalf("Could not parse time: %v\n", err)
		}
	}

	return
}

func main() {
	t, err := getLimit()
	if err != nil {
		log.Fatalf("Could not get limit: %v\n", err)
	}

	fmt.Printf("Marking all notifications read as of %s\n", t.Format(time.RFC1123Z))

	client := http.Client{Timeout: time.Second * 2}

	body := []byte(fmt.Sprintf(`{"last_read_at":"%s"}`, t.Format(time.RFC3339)))

	b := bytes.NewBuffer(body)

	req, err := utils.Request(http.MethodPut, b)
	if err != nil {
		log.Fatalf("Error creating request: %v\n", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error executing request: %v\n", err)
	}

	fmt.Printf("Response: %s\n", res.Status)

	fmt.Println("Done.")
}
