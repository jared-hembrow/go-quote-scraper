// Package main provides a program to scrape quotes from a website.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Scrap struct represents a scraping instance.
type Scrap struct {
	url        string
	stringBody string
	quoteList  []Quote
}

// Quote struct represents a quote with its author.
type Quote struct {
	quote  string
	author string
}

// getPage fetches the HTML page from the provided URL.
func (p *Scrap) getPage() {
	fmt.Println("Getting page: ", p.url)
	resp, err := http.Get(p.url)
	if err != nil {
		log.Fatal("Error:", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error:", err)
	}

	p.stringBody = string(body)
}

// parsePage parses the HTML page and extracts quotes.
func (p *Scrap) parsePage() {
	if p.stringBody == "" {
		log.Println("Error: Empty page")
		return
	}

	doc, err := html.Parse(strings.NewReader(p.stringBody))
	if err != nil {
		log.Fatal("Error:", err)
	}
	p.forEachNode(doc, startElement, endElement)
}

// extractQuote extracts quotes from the HTML node.
func (p *Scrap) extractQuote(n *html.Node) {
	child := n.FirstChild
	var newQuote Quote
	for i := 0; i < 10 && child != nil; i++ {
		if child.Type == html.ElementNode && child.FirstChild != nil {
			for _, a := range child.Attr {
				if a.Key == "class" && a.Val == "text" {
					newQuote.quote = strings.TrimSpace(child.FirstChild.Data)
				}
			}
			if strings.TrimSpace(child.FirstChild.Data) == "by" {
				if child.FirstChild.NextSibling != nil && child.FirstChild.NextSibling.FirstChild != nil {
					newQuote.author = child.FirstChild.NextSibling.FirstChild.Data
				}
			}
		}
		child = child.NextSibling
	}
	if newQuote.quote != "" && newQuote.author != "" {
		p.quoteList = append(p.quoteList, newQuote)
	}
}

// gotoNextPage navigates to the next page for scraping.
func (p *Scrap) gotoNextPage(n *html.Node) {
	child := n.FirstChild
	for i := 0; i < 10 && child != nil; i++ {
		for _, a := range child.Attr {
			if a.Key == "href" {
				p.url = "https://quotes.toscrape.com" + a.Val
				p.getPage()
				p.parsePage()
				break
			}
		}
		child = child.NextSibling
	}
}

// forEachNode iterates over HTML nodes recursively.
func (p *Scrap) forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	for _, a := range n.Attr {
		if a.Key == "class" && a.Val == "quote" {
			p.extractQuote(n)
		}
		if a.Key == "class" && a.Val == "next" {
			p.gotoNextPage(n)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		p.forEachNode(c, pre, post)
	}
}

// startElement prints the start HTML element tag.
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("<%s>\n", n.Data)
	}
}

// endElement prints the end HTML element tag.
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("</%s>\n", n.Data)
	}
}

func main() {
	// Initialize a Scrap instance with the URL.
	scrappy := Scrap{
		url: "https://quotes.toscrape.com/",
	}
	// Fetch the page and parse it.
	scrappy.getPage()
	scrappy.parsePage()
	// Print the scraped quotes.
	for _, q := range scrappy.quoteList {
		fmt.Printf("\n\t%s\n\tBy:\t%s\n\n", q.quote, q.author)
	}
}
