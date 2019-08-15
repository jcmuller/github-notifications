package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jcmuller/github-notifications/internal/utils"
)

func getDate() (t time.Time, err error) {
	var d time.Duration

	if len(os.Args) == 1 {
		d, err = time.ParseDuration("-24h")
		if err != nil {
			return
		}

		t = time.Now().Add(d)
	} else {
		// Mon/Day
		// 01/02 03:04:05PM ‘06 -0700
		t, err = time.Parse("2006/01/02 15:04", os.Args[1])

		if err != nil {
			err = nil
			d, err = time.ParseDuration(os.Args[1])
			t = time.Now().Add(d)
		}
	}

	return
}

func main() {
	t, err := getDate()
	if err != nil {
		log.Fatalf("Could not get a date: %v\n", err)
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
