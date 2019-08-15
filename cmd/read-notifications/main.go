package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jcmuller/github-notifications/internal/utils"
)

// Subject holds the subject
type Subject struct {
	Title string `json:"title"`
	Type  string `json:"type"`
	URL   string `json:"url"`
}

func (s *Subject) String() string {
	url, err := url.Parse(s.URL)
	if err != nil {
		panic(err)
	}

	url.Host = "github.com"
	url.Path = strings.Replace(url.Path, "/repos", "", 1)
	url.Path = strings.Replace(url.Path, "pulls/", "pull/", 1)

	s.URL = url.String()

	return fmt.Sprintf("[%s] %s -- %s", s.Type, s.Title, s.URL)
}

func (n *Notification) String() string {
	return fmt.Sprintf("%s: [%s]%s %s",
		fmtDuration(time.Since(n.When)),
		n.Reason,
		n.Subject.String(),
		n.Description,
	)
}

// Notification holds the notification
type Notification struct {
	When        time.Time `json:"updated_at"`
	Unread      bool      `json:"unread"`
	Reason      string    `json:"reason"`
	Subject     Subject   `json:"subject"`
	Description string    `json:"description"`
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute

	if h == 0 {
		return fmt.Sprintf("%2d minutes ago", m)
	}

	return fmt.Sprintf("%2d hours ago", h)
}

func main() {
	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := utils.Request(http.MethodGet, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v\n", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error executing request: %v\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v\n", err)
	}

	var notifications []Notification
	err = json.Unmarshal(body, &notifications)
	if err != nil {
		log.Fatalf("Error unmarshaling response body: %v\n", err)
	}

	for _, notification := range notifications {
		fmt.Println(notification.String())
	}
}
