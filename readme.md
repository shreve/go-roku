Roku ECP Client
===============

This is a go port of the [Roku ECP Client](https://github.com/shreve/roku-ecp) I wrote in Ruby.

## Usage

Create a client for a Roku device by passing it's IP to `Connect`. If you don't know the IP, you can search the network for it with `Discover`

```go
import "github.com/shreve/go-roku/roku"

var client roku.Client
host := os.Getenv("ROKU_HOST") // 192.168.0.100
if host == "" {
    client, err = roku.Connect(host)
} else {
    client, err = roku.Discover()
}
```

Once you have a client for a device on the network, you can access the ECP API via various helper functions.

```go
// Turn up the volume for one second
client.Keydown(roku.VolumeUp)
time.Sleep(time.Second)
client.Keyup(roku.VolumeUp)

// Press the power button
client.Keypress(roku.PowerOff)

// Print the name of the device
info := client.DeviceInfo()
fmt.Println(info.FriendlyDeviceName)

// Print the name of the currently running app
app := client.ActiveApp()
fmt.Println(app.Name)

// Launch last app on the list
apps := client.Apps()
client.Launch(apps[len(apps)-1].Id)
```
