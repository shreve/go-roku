package roku

import (
	"encoding/xml"
)

const (
	Home          = "Home"
	Rev           = "Rev"
	Fwd           = "Fwd"
	Play          = "Play"
	Select        = "Select"
	Left          = "Left"
	Right         = "Right"
	Down          = "Down"
	Up            = "Up"
	Back          = "Back"
	InstantReplay = "InstantReplay"
	Info          = "Info"
	Backspace     = "Backspace"
	Search        = "Search"
	Enter         = "Enter"
	FindRemote    = "FindRemote"
	VolumeDown    = "VolumeDown"
	VolumeMute    = "VolumeMute"
	VolumeUp      = "VolumeUp"
	PowerOff      = "PowerOff"
	ChannelUp     = "ChannelUp"
	ChannelDown   = "ChannelDown"
	InputTuner    = "InputTuner"
	InputHDMI1    = "InputHDMI1"
	InputHDMI2    = "InputHDMI2"
	InputHDMI3    = "InputHDMI3"
	InputHDMI4    = "InputHDMI4"
	InputAV1      = "InputAV1"
)

type DeviceInfo struct {
	AdvertisingId               string `xml:"advertising-id"`
	BuildNumber                 string `xml:"build-number"`
	CanUseWifiExtender          bool   `xml:"can-use-wifi-extender"`
	ClockFormat                 string `xml:"clock-format"`
	Country                     string `xml:"country"`
	DavinciVersion              string `xml:"davinci-version"`
	DefaultDeviceName           string `xml:"default-device-name"`
	DeveloperEnabled            bool   `xml:"developer-enabled"`
	DeviceId                    string `xml:"device-id"`
	EthernetMAC                 string `xml:"ethernet-mac"`
	FindRemoteIsPossible        bool   `xml:"find-remote-is-possible"`
	FriendlyDeviceName          string `xml:"friendly-device-name"`
	FriendlyModelName           string `xml:"friendly-model-name"`
	GrandcentralVersion         string `xml:"grandcentral-version"`
	HasMobileScreensaver        bool   `xml:"has-mobile-screensaver"`
	HasPlayOnRoku               bool   `xml:"has-play-on-roku"`
	HasWifi5GSupport            bool   `xml:"has-wifi-5G-support"`
	HasWifiExtender             bool   `xml:"has-wifi-extender"`
	HeadphonesConnected         bool   `xml:"headphones-connected"`
	IsStick                     bool   `xml:"is-stick"`
	IsTV                        bool   `xml:"is-tv"`
	Language                    string `xml:"language"`
	Locale                      string `xml:"locale"`
	ModelName                   string `xml:"model-name"`
	ModelNumber                 string `xml:"model-number"`
	ModelRegion                 string `xml:"model-region"`
	NetworkName                 string `xml:"network-name"`
	NetworkType                 string `xml:"network-type"`
	NotificationsEnabled        bool   `xml:"notifications-enabled"`
	NotificationsFirstUse       bool   `xml:"notifications-first-use"`
	PanelId                     uint   `xml:"panel-id"`
	PowerMode                   string `xml:"power-mode"`
	ScreenSize                  uint   `xml:"screen-size"`
	SearchChannelsEnabled       bool   `xml:"search-channels-enabled"`
	SearchEnabled               bool   `xml:"search-enabled"`
	SecureDevice                bool   `xml:"secure-device"`
	SerialNumber                string `xml:"serial-number"`
	SoftwareBuild               string `xml:"software-build"`
	SoftwareVersion             string `xml:"software-version"`
	SupportURL                  string `xml:"support-url"`
	SupportsAudioGuide          bool   `xml:"supports-audio-guide"`
	SupportsECSMicrophone       bool   `xml:"supports-ecs-microphone"`
	SupportsECSTextedit         bool   `xml:"supports-ecs-textedit"`
	SupportsEthernet            string `xml:"supports-ethernet"`
	SupportsFindRemote          bool   `xml:"supports-find-remote"`
	SupportsPrivateListening    bool   `xml:"supports-private-listening"`
	SupportsPrivateListeningDTV bool   `xml:"supports-private-listening-dtv"`
	SupportsRVA                 bool   `xml:"supports-rva"`
	SupportsSuspend             bool   `xml:"supports-suspend"`
	SupportsWakeOnWLAN          bool   `xml:"supports-wake-on-wlan"`
	SupportsWarmStandby         bool   `xml:"supports-warm-standby"`
	TRCChannelVersion           string `xml:"trc-channel-version"`
	TRCVersion                  string `xml:"trc-version"`
	TimeZone                    string `xml:"time-zone"`
	TimeZoneAuto                bool   `xml:"time-zone-auto"`
	TimeZoneName                string `xml:"time-zone-name"`
	TimeZoneOffset              string `xml:"time-zone-offset"`
	TimeZoneTZ                  string `xml:"time-zone-tz"`
	TunerType                   string `xml:"tuner-type"`
	UDN                         string `xml:"udn"`
	Uptime                      uint   `xml:"uptime"`
	UserDeviceLocation          string `xml:"user-device-location"`
	UserDeviceName              string `xml:"user-device-name"`
	VendorName                  string `xml:"vendor-name"`
	VoiceSearchEnabled          bool   `xml:"voice-search-enabled"`
	WiFiDriver                  string `xml:"wifi-driver"`
	WiFiMAC                     string `xml:"wifi-mac"`
	// <keyed-developer-id/>
	// <expert-pq-enabled>1.0</expert-pq-enabled>
}

type AppList struct {
	Apps []App `xml:"app"`
}

type ActiveApp struct {
	XMLName xml.Name `xml:"active-app"`
	App     App      `xml:"app"`
}

type App struct {
	Id      string `xml:"id,attr"`
	Type    string `xml:"type,attr"`
	Version string `xml:"version,attr"`
	Name    string
}

func (a *App) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
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

type Player struct {
	XMLName  xml.Name `xml:"player"`
	Error    bool     `xml:"error,attr"`
	State    string   `xml:"state,attr"`
	Position string   `xml:"position"`
	IsLive   bool     `xml:"is_live"`
	Plugin   Plugin   `xml:"plugin"`
	Format   Format   `xml:"format"`
}

type Plugin struct {
	XMLName   xml.Name `xml:"plugin"`
	Bandwidth string   `xml:"bandwidth,attr"`
	Id        string   `xml:"id,attr"`
	Name      string   `xml:"name,attr"`
}

type Format struct {
	XMLName  xml.Name `xml:"format"`
	Audio    string   `xml:"audio,attr"`
	Captions string   `xml:"captions,attr"`
	Drm      string   `xml:"drm,attr"`
	Video    string   `xml:"video,attr"`
}
