
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

## Usage

Help/Usage is printed out when no parameter is given to the binary.

```bash
/usr/bin/i3-brightnessctl; # shows help
```


## License

GPL-3.0

