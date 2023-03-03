package logger

import (
	"log"
	"time"
)

// TimeTemplate for logger
const timeTemplate = time.RFC822 // "2006.01.02 15:04:05"

// Simple request log
func Request(url, from string) {
	log.Printf("[%s] 🔎 New Request: %s From: %s\n", time.Now().Format(timeTemplate), url, from)
}

// Simple logging
func Log(msg string) {
	log.Printf("[%s] ⚠ Log: %s\n", time.Now().Format(timeTemplate), msg)
}

// Simple error log
func Error(err error, msg string) {
	log.Printf("[%s] ❌ Error! %s: %s\n", time.Now().Format(timeTemplate), msg, err)
}
