package main

import (
	log "github.com/lfkeitel/verbose/v5"
)

func init() {
	// Configure logger to print stderr using logfmt lines
	// log.ClearTransports() // Clear default transports
	// log.AddTransport(log.NewTextTransportWith(log.NewLogFmtFormatter()))
}

func main() {
	// do stuff...
	log.WithFields(log.Fields{
		"foo":    "bar",
		"result": 3,
	}).Info("This happened")
}
