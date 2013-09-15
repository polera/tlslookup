package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/polera/tlskit"
	"log"
	"os"
	"strconv"
	s "strings"
)

var (
	domainName = flag.String("domainName", "uncryptic.com", "FQDN of the domain you want to lookup.")
	serverPort = flag.Int("port", 443, "HTTPS port used by the domain you're looking up.")
	lookupFile = flag.String("file", "", "A file containing multiple hosts to lookup.  This overrides other parameters.")
)

func main() {
	flag.Parse()
	m := tlskit.TLSRequest{}

	if *lookupFile == "" {
		m.Requests = append(m.Requests, tlskit.Request{*domainName, int32(*serverPort)})
	} else {
		file, _ := os.Open(*lookupFile)
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			splitLine := s.Split(scanner.Text(), ",")
			port, _ := strconv.Atoi(splitLine[1])
			m.Requests = append(m.Requests, tlskit.Request{splitLine[0], int32(port)})
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}

	results, err := tlskit.Lookup(m)
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}

}
