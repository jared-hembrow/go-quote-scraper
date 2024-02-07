# Quote Scraper

## Overview

This is a simple Go program that scrapes quotes from a website (https://quotes.toscrape.com/). It utilizes the `net/http` and `golang.org/x/net/html` packages for fetching and parsing HTML content.

## Features

- Fetches HTML content from a specified URL.
- Parses the HTML content to extract quotes and their authors.
- Navigates to the next page recursively to scrape more quotes.
- Prints the scraped quotes.

## Usage

To use this program, simply run the main.go file:

```bash
go run main.go
```

This will output the scraped quotes to the console.

### Project Structure

- `main.go`: Contains the main program logic.
- `README.md`: Provides information about the project.

### Lessons Learned

Through this project, several key lessons were learned:

- **Web Scraping:** Gain a practical understanding of web scraping techniques using Go.
- **HTML Parsing:** Learn how to parse HTML content using the golang.org/x/net/html package.
- **HTTP Requests:** Understand how to make HTTP requests and handle responses using the net/http package.
- **Recursion:** Apply recursion to navigate through paginated content for scraping.

### Next Steps

Future enhancements to this project could include:

- Error handling improvements for robustness.
- Enhanced user interface for better interaction.
- Support for scraping additional content types beyond quotes.
