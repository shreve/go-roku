package roku

import (
	"encoding/xml"
)

const (
	Home = "Home"
	Rev = "Rev"
	Fwd = "Fwd"
	Play = "Play"
	Select = "Select"
	Left = "Left"
	Right = "Right"
	Down = "Down"
	Up = "Up"
	Back = "Back"
	InstantReplay = "InstantReplay"
	Info = "Info"
	Backspace = "Backspace"
	Search = "Search"
	Enter = "Enter"
	FindRemote = "FindRemote"
	VolumeDown = "VolumeDown"
	VolumeMute = "VolumeMute"
	VolumeUp = "VolumeUp"
	PowerOff = "PowerOff"
	ChannelUp = "ChannelUp"
	ChannelDown = "ChannelDown"
	InputTuner = "InputTuner"
	InputHDMI1 = "InputHDMI1"
	InputHDMI2 = "InputHDMI2"
	InputHDMI3 = "InputHDMI3"
	InputHDMI4 = "InputHDMI4"
	InputAV1 = "InputAV1"
)

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
