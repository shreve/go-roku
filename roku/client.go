package roku

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"encoding/xml"
)

type Client struct {
	Ready bool
	address string
	client *http.Client
}

func Connect(addr string) (Client, error) {
	if addr[0:4] != "http" {
		addr = "http://" + addr
	}
	url, _ := url.Parse(addr)
	url.Scheme = "http"
	url.Path = "/"
	if url.Port() == "" {
		url.Host += ":8060"
	}
	return Client{true, url.String(), &http.Client{}}, nil
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
