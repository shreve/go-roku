package roku

import (
	"net"
	"net/http"
	"bytes"
	"regexp"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/xml"
)

type Client struct {
	address string
	client *http.Client
}

type AppList struct {
	Apps []App `xml:"app"`
}
type ActiveApp struct {
	XMLName xml.Name `xml:"active-app"`
	App App `xml:"app"`
}

type App struct {
	Id string  `xml:"id,attr"`
	Name string
}

type RokuInfo struct {
	UDN string `xml:"udn"`
	SerialNumber string `xml:"serial-number"`
	DeviceId string `xml:"device-id"`
	AdvertisingId string `xml:""`
	VendorName string `xml:"vendor-name"`
	ModelName string `xml:"model-name"`
	ModelNumber string `xml:"model-number"`
	ModelRegion string `xml:"model-region"`
	IsTV bool `xml:"is-tv"`
	IsStick bool `xml:"is-stick"`
	ScreenSize uint `xml:"screen-size"`
	PanelId uint `xml:"panel-id"`
	TunerType string `xml:"tuner-type"`
	SupportsEthernet string `xml:"supports-ethernet"`
	WiFiMAC string `xml:"wifi-mac"`
	WiFiDriver string `xml:"wifi-driver"`
	EthernetMAC string `xml:"ethernet-mac"`
	NetworkType string `xml:"network-type"`
	FriendlyDeviceName string `xml:"friendly-device-name"`
	FriendlyModelName string `xml:"friendly-model-name"`
	DefaultDeviceName string `xml:"default-device-name"`
	UserDeviceName string `xml:"user-device-name"`
	BuildNumber string `xml:"build-number"`
	SoftwareVersion string `xml:"software-version"`
	SoftwareBuild string `xml:"software-build"`
	SecureDevice bool `xml:"secure-device"`
	Language string `xml:"language"`
	Country string `xml:"country"`
	Locale string `xml:"locale"`
	TimeZoneAuto bool `xml:"time-zone-auto"`
	TimeZone string `xml:"time-zone"`
	TimeZoneName string `xml:"time-zone-name"`
	TimeZoneTZ string `xml:"time-zone-tz"`
	TimeZoneOffset string `xml:"time-zone-offset"`
	ClockFormat string `xml:"clock-format"`
	Uptime uint `xml:"uptime"`
	PowerMode string `xml:"power-mode"`
}

func (a *App) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range(start.Attr) {
		if attr.Name.Local == "id" {
			a.Id = attr.Value
		}
	}
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	a.Name = s
	return nil
}

func (c *Client) Get(path string) ([]byte, error) {
	url := c.address + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil { return []byte{}, err }
	resp, err := c.client.Do(req)
	if err != nil { return []byte{}, err }
	return ioutil.ReadAll(resp.Body)
}

func (c *Client) Post(path string) error {
	url := c.address + path
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	_, err = c.client.Do(req)
	return err
}

func (c *Client) Apps() []App {
	body, err := c.Get("query/apps")
	if err != nil { panic(err) }
	al := AppList{}
	xml.Unmarshal(body, &al)
	return al.Apps
}

func (c *Client) ActiveApp() App {
	body, err := c.Get("query/active-app")
	if err != nil { panic(err) }
	aa := ActiveApp{}
	xml.Unmarshal(body, &aa)
	return aa.App
}

func (c *Client) DeviceInfo() RokuInfo {
	body, err := c.Get("query/device-info")
	if err != nil { panic(err) }
	info := RokuInfo{}
	fmt.Println(string(body))
	xml.Unmarshal(body, &info)
	return info
}

func (c *Client) Keyup(key string) error {
	return c.Post("keyup/" + key)
}

func (c *Client) Keydown(key string) error {
	return c.Post("keydown/" + key)
}

func (c *Client) Keypress(key string) error {
	return c.Post("keypress/" + key)
}

func (c *Client) Launch(appId int) error {
	return c.Post(fmt.Sprintf("launch/%d", appId))
}

func (c *Client) Install(appId int) error {
	return c.Post(fmt.Sprintf("install/%d", appId))
}

const ssdpAddr = "239.255.255.250:1900"
const searchRequest = "M-SEARCH * HTTP/1.1\nHost: 239.255.255.250:1900\nMan: \"ssdp:discover\"\nST: roku:ecp\n\n"

func Connect(addr string) (Client, error) {
	return Client{addr, &http.Client{}}, nil
}

func Search() (Client, error) {

	// Bind a local UDP socket to listen for responses
	sock, err := net.ListenPacket("udp", ":0")
	if err != nil { panic(err) }
	defer sock.Close()

	// Write the search request to the SSDP address
	destAddr, err := net.ResolveUDPAddr("udp", ssdpAddr)
	sock.WriteTo([]byte(searchRequest), destAddr)

	for {

		// Read in responses (UDP datagrams typically 8kb)
		buf := make([]byte, 8192)
		_, _, err := sock.ReadFrom(buf)
		if err != nil { panic(err) }

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
