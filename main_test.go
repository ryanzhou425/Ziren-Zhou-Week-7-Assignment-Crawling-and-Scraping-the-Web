package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/gocolly/colly"
)

// Test URL fetching
func TestURLFetching(t *testing.T) {
	// Initialize the Colly collector
	c := colly.NewCollector()

	// Store the scraped text
	var scrapedText string
	c.OnHTML("body", func(e *colly.HTMLElement) {
		scrapedText = strings.TrimSpace(e.Text)
	})

	// Test URL
	testURL := "https://example.com"
	err := c.Visit(testURL)
	if err != nil {
		t.Errorf("Failed to visit URL %s: %v", testURL, err)
	}

	// Check if any content was scraped
	if scrapedText == "" {
		t.Errorf("Empty content")
	}
}

// Test JSON data generation
func TestJSONOutput(t *testing.T) {
	// Use struct to ensure field order
	type PageData struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	}

	data := PageData{
		URL:  "https://example.com",
		Text: "Sample text",
	}

	expected := `{"url":"https://example.com","text":"Sample text"}`

	// Convert struct to JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Errorf("Failed to marshal JSON: %v", err)
	}

	// Verify the output
	if string(jsonData) != expected {
		t.Errorf("Expected %s, but got %s", expected, string(jsonData))
	}
}

// Test file creation and writing
func TestFileCreationAndWriting(t *testing.T) {
	filePath := "test.jl"

	// Create the test file
	file, err := os.Create(filePath)
	if err != nil {
		t.Errorf("Failed to create file: %v", err)
	}
	defer file.Close()

	// Write test data
	testData := map[string]string{
		"url":  "https://example.com",
		"text": "Sample text",
	}
	jsonData, err := json.Marshal(testData)
	if err != nil {
		t.Errorf("Failed to marshal JSON: %v", err)
	}

	_, err = file.WriteString(string(jsonData) + "\n")
	if err != nil {
		t.Errorf("Failed to write to file: %v", err)
	}

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("File %s doesn't exist", filePath)
	}
}
