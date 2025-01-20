
# i3-brightnessctl

Standalone and dead-simple brightness controller that works for both Linux-supported `backlight`
controllers and via `xrandr` connected monitors.


## Installation

```bash
git clone git@github.com:cookiengineer/i3-brightnessctl.git;
cd ./i3-brightnessctl;

go build cmds/i3-brightnessctl/main.go -o i3-brightnessctl;

sudo cp i3-brightnessctl /usr/bin/i3-brightnessctl;
sudo chmod +x /usr/bin/i3-brightnessctl;
```


## Non-Root Usage

After installation, it might be necessary to add a `udev` rule to be able to modify the
`/sys/class/backlight/*/brightness` files as a `wheel` user without `root` rights.

Store the following content as `/etc/udev/rules.d/90-backlight.rules`:

```udev
SUBSYSTEM=="backlight", ACTION=="add", RUN+="/bin/chgrp wheel /sys/class/backlight/%k/brightness", RUN+="/bin/chmod g+w /sys/class/backlight/%k/brightness"
```


## Usage

Help/Usage is printed out when no parameter is given to the binary.

```bash
/usr/bin/i3-brightnessctl; # shows help
```

**CLI Parameters**:

- `i3-brightnessctl increase-all` increases the brightness of all connected monitors.
- `i3-brightnessctl decrease-all` decreases the brightness of all connected monitors.
- `i3-brightnessctl increase` increases the brightness of the primary connected monitor.
- `i3-brightnessctl decrease` decreases the brightness of the primary connected monitor.
- `i3-brightnessctl status` prints out the brightness percentage of the primary connected monitor (for use with `i3-status`).

**Usage with i3**:

Add these lines to the `~/.config/i3/config` file to be able to change the primary monitor's
brightness with the corresponding keys:

```i3conf
bindsym XF86MonBrightnessUp exec --no-startup-id i3-brightnessctl increase
bindsym XF86MonBrightnessDown exec --no-startup-id i3-brightnessctl decrease
```

**Usage with i3status**:

As i3status does not have any external program execution, the author wants us to use a
wrapper shell script that concats together all status information. For more details, check
the [official i3status documentation](https://i3wm.org/docs/i3status.html).

This example will prepend the brightness percentage to i3status:

```bash
#!/bin/bash
# stored as ~/.bin/i3status.sh

brightness=`i3-brightnessctl status`;

i3status | while :
do
        read line
        echo "Brightness: ${brightness} | $line" || exit 1
done
```


## License

GPL-3.0

