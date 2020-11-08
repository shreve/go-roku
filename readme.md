Roku ECP Client
===============

This is a go port of the [Roku ECP Client](https://github.com/shreve/roku-ecp) I wrote in Ruby.

## TUI Client Usage

Install then run the client

```
$ go install github.com/shreve/go-roku/cmd/roku

$ roku
```

This runs an interactive remote control program. By default, it scans the
network and connects to the first Roku device it finds. If you'd prefer to
hard-code the address, use the environment variable `$ROKU_HOST`.

Here are some of the key bindings

| Key             | Command                                           |
|-----------------|---------------------------------------------------|
| arrow keys      | navigation                                        |
| h/j/k/l         | vi-style navigation                               |
| ctrl+up/down    | volume up/down                                    |
| ctrl+left/right | toggle mute                                       |
| space bar       | play/pause                                        |
| asterisk        | info/options (asterisk button on physical remote) |
| esc             | back/exit                                         |
| backspace       | back/exit                                         |
| enter           | select                                            |
| ctrl+q          | power off                                         |
| q               | quit the remote app                               |
| ctrl+c          | quit the remote app                               |
| o               | open app (search interface)                       |


## Library Usage

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

## Todo

* Allow retrying discovery rather than quitting in place

## Wishlist

* Add help screen to app
* Better reporting in app
* Onscreen keyboard tool
* Create MPRIS daemon to allow controlling remote device with media keys
