package roku

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var ClientNotReady = errors.New("The client is not yet ready.")

func Connect(addr string) Client {
	if addr[0:4] != "http" {
		addr = "http://" + addr
	}
	url, _ := url.Parse(addr)
	url.Scheme = "http"
	url.Path = "/"
	if url.Port() == "" {
		url.Host += ":8060"
	}
	return Client{true, url.String(), &http.Client{}}
}

type Client struct {
	ready   bool
	Address string
	client  *http.Client
}

func (c *Client) Ready() bool {
	return c.ready
}

func (c *Client) Apps() ([]App, error) {
	body, err := c.get("query/apps")
	if err != nil {
		return []App{}, err
	}
	al := appList{}
	xml.Unmarshal(body, &al)
	return al.Apps, nil
}

func (c *Client) ActiveApp() (App, error) {
	body, err := c.get("query/active-app")
	if err != nil {
		return App{}, err
	}
	aa := activeApp{}
	xml.Unmarshal(body, &aa)
	return aa.App, nil
}

func (c *Client) DeviceInfo() (info DeviceInfo, err error) {
	body, err := c.get("query/device-info")
	if err != nil {
		return
	}
	xml.Unmarshal(body, &info)
	return
}

func (c *Client) MediaPlayer() (player Player, err error) {
	body, err := c.get("query/media-player")
	if err != nil {
		return
	}
	xml.Unmarshal(body, &player)
	return
}

func (c *Client) Keyup(key string) error {
	return c.post("keyup/" + key)
}

func (c *Client) Keydown(key string) error {
	return c.post("keydown/" + key)
}

func (c *Client) Keypress(key string) error {
	return c.post("keypress/" + key)
}

func (c *Client) Launch(appId string) error {
	return c.post(fmt.Sprintf("launch/%s", appId))
}

func (c *Client) Install(appId int) error {
	return c.post(fmt.Sprintf("install/%d", appId))
}

func (c *Client) get(path string) (body []byte, err error) {
	if !c.ready {
		return []byte{}, ClientNotReady
	}
	url := c.Address + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	return ioutil.ReadAll(resp.Body)
}

func (c *Client) post(path string) error {
	if !c.ready {
		return ClientNotReady
	}
	url := c.Address + path
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	_, err = c.client.Do(req)
	return err
}
