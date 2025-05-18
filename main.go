package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	// Record start time
	startTime := time.Now()

	// List of URLs to scrape
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	// Open the output file
	file, err := os.Create("web_scraped.jl")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	// Initialize the Colly collector
	c := colly.NewCollector()

	// Define the scraping logic
	c.OnHTML("body", func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.Text)
		data := map[string]string{
			"url":  e.Request.URL.String(),
			"text": text,
		}
		jsonData, _ := json.Marshal(data)
		file.WriteString(string(jsonData) + "\n")
		fmt.Println("Scraped:", e.Request.URL.String())
	})

	// Visit each URL in the list
	for _, url := range urls {
		err := c.Visit(url)
		if err != nil {
			log.Printf("Error visiting %s: %v", url, err)
		}
	}

	// Record end time and print duration
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("\nScraping completed in %.2f seconds\n", duration.Seconds())
}
