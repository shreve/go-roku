package roku

import (
	"net"
	"bytes"
	"log"
	"regexp"
	"time"
)

const ssdpAddr = "239.255.255.250:1900"
const searchRequest =
`M-SEARCH * HTTP/1.1
Host: 239.255.255.250:1900
Man: "ssdp:discover"
ST: roku:ecp
` // Extra newline required

func Discover() (Client, error) {

	client := Client{}

	// Bind a local UDP socket to listen for responses
	sock, err := net.ListenPacket("udp", ":0")
	if err != nil { return client, err }
	defer sock.Close()

	// Write the search request to the SSDP address
	destAddr, err := net.ResolveUDPAddr("udp", ssdpAddr)
	if err != nil { return client, err }
	sock.WriteTo([]byte(searchRequest), destAddr)

	for {

		// Read in responses (UDP datagrams typically 8kb)
		buf := make([]byte, 8192)
		sock.SetReadDeadline(time.Now().Add(time.Second * 3))
		_, _, err := sock.ReadFrom(buf)
		if err != nil { return client, err }

		// Validate response from Roku Device
		if ! bytes.Contains(buf, []byte("uuid:roku:ecp")) {
			log.Println("This response doesn't look like it came from a Roku device. Skipping.")
			continue
		}

		// Search for the location header and pull out the URL
		re := regexp.MustCompile(`(?i)location: (.*)\r\n`)
		address := re.FindSubmatch(buf)
		if len(address) < 2 {
			log.Println("Unable to parse device address from response. Skipping.")
			continue
		}

		return Connect(string(address[1]))
	}
}
