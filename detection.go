package main

import (
        "context"
        "fmt"
        "log"
)

func main() {
        // Create a new context
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()

        // Create a new Chrome instance
        c, err := chromedp.New(ctx, chromedp.WithLogf(log.Printf))
        if err != nil {
                log.Fatal(err)
        }
        defer c.Close()

        // Get the current active tab
        var currentURL string
        err = c.Run(ctx, chromedp.Navigate(`chrome://extensions`))
        if err != nil {
                log.Fatal(err)
        }
        err = c.Run(ctx, chromedp.Evaluate(`document.querySelector('extensions-frame').contentDocument.activeTab.url`, &currentURL))
        if err != nil {
                log.Fatal(err)
        }

        // Perform phishing checks (replace with your actual logic)
        if isPhishing(currentURL) {
                fmt.Println("Potential phishing site detected!")
        } else {
                fmt.Println("No phishing indicators found.")
        }
}

func isPhishing(url string) bool {
        // Implement your phishing detection logic here
        // (e.g., check for typosquatting, SSL issues, etc.)
        // Example: Check for suspicious URL patterns
        suspiciousPatterns := []string{
                "example.com.co", // Typosquatting
                "xn--example-com", // Punycode
        }
        for _, pattern := range suspiciousPatterns {
                if url == pattern {
                        return true
                }
        }
        return false
}