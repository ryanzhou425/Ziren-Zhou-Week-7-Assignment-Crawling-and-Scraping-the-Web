# Project Overview
This project utilizes Go Colly to scrape pages from various Wikipedia sites, storing the extracted text data in JSON Lines format. In addition, it includes a comparative study with Python Scrapy. By implementing the same scraping functionality with both Go Colly and Scrapy, and recording the processing time, this project aims to analyze their differences in efficiency, stability, and scalability.

---

# Program Features
The core features of this project include:
- Webpage text extraction
- File storage
- Processing time logging

---

# How to Run
Use the following command to run the program: 
- go run main.go

If everything works correctly, the terminal will display: 
- Scraping completed in xxx seconds

Output generated: 
- web_scraped.jl

---

# How to Test
Run the following command to execute the unit tests: 
- go test

If the test runs successfully, it generates the file: 
- test.jl

---

# Performance Comparison
The performance comparison was conducted using the Python code from [WebFocusedCrawlWorkV001](WebFocusedCralWorkV001) on the same set of Wikipedia pages. Below are the results:
- Scrapy: 2.42 seconds 
- Go Colly: 0.64 seconds

The results revealed significant differences between Python Scrapy and Go Colly. Go is nearly 4 times faster than Python. This improvement is primarily due to Go's advantages in concurrency and memory management, indicating that Go Colly is a superior choice for handling large-scale web scraping tasks.

---

# Use of AI Assistants
- Searched for methods to define scraping logic in Go
- Searched for approaches to crawl URLs individually in Go
- Searched for how to initialize the Colly collector in Go

