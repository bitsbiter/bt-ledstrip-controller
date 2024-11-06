package main

import (
    "errors"
    "image/color"
    "strconv"
    "strings"
    "flag"
)

// parses a hex color string or RGB values into a color.RGBA struct
func ParseColor(s string) (color.RGBA, error) {
    if strings.HasPrefix(s, "#") {
        s = s[1:]
    }
    if len(s) == 6 {
        // Hex color
        r, err := strconv.ParseUint(s[0:2], 16, 8)
        if err != nil {
            return color.RGBA{}, err
        }
        g, err := strconv.ParseUint(s[2:4], 16, 8)
        if err != nil {
            return color.RGBA{}, err
        }
        b, err := strconv.ParseUint(s[4:6], 16, 8)
        if err != nil {
            return color.RGBA{}, err
        }
        return color.RGBA{uint8(r), uint8(g), uint8(b), 0xFF}, nil
    } else if strings.Contains(s, ",") {
        // RGB values
        parts := strings.Split(s, ",")
        if len(parts) != 3 {
            return color.RGBA{}, errors.New("invalid RGB format")
        }
        r, err := strconv.Atoi(strings.TrimSpace(parts[0]))
        if err != nil {
            return color.RGBA{}, err
        }
        g, err := strconv.Atoi(strings.TrimSpace(parts[1]))
        if err != nil {
            return color.RGBA{}, err
        }
        b, err := strconv.Atoi(strings.TrimSpace(parts[2]))
        if err != nil {
            return color.RGBA{}, err
        }
        return color.RGBA{uint8(r), uint8(g), uint8(b), 0xFF}, nil
    } else {
        return color.RGBA{}, errors.New("invalid color format")
    }
}

// validate combination of arguments
func ValidateArguments(args Arguments) error {
    if args.Power != "" && args.Pattern != "" && args.Color != "" {
        return errors.New("cannot use 'color' and 'pattern' together")
    }
    if args.Speed != -1 && args.Pattern == "" {
        return errors.New("'speed' can only be used with 'pattern'")
    }
    if args.Brightness != -1 && args.Power == "" && args.Color == "" && args.Pattern == "" {
        return errors.New("'brightness' must be used with 'power', 'color', or 'pattern'")
    }
    return nil
}

type Arguments struct {
    Power      string
    Color      string
    Brightness int
    Pattern    string
    Speed      int
    Command    string
    DeviceMAC  string
    DeviceName string
    List       bool
}

// parse command-line arguments into the Arguments struct
func ParseArguments() Arguments {
    var args Arguments
    args.Brightness = -1
    args.Speed = -1
    args.DeviceMAC = "C0:00:00:00:01:01"
    args.DeviceName = "LEDDMX-00-000101"

    // Use the flag package to parse command-line arguments
    flag.StringVar(&args.Power, "power", "", "Set power state: on or off")
    flag.StringVar(&args.Color, "color", "", "Set color (hex code or R,G,B values)")
    flag.IntVar(&args.Brightness, "brightness", -1, "Set brightness (0-100)")
    flag.StringVar(&args.Pattern, "pattern", "", "Set pattern (index or name)")
    flag.IntVar(&args.Speed, "speed", -1, "Set speed (0-100)")
    flag.StringVar(&args.Command, "command", "", "Command: check, setmac, setname, listpatterns")
    flag.StringVar(&args.DeviceMAC, "mac", "C0:00:00:00:01:01", "Device MAC address")
    flag.StringVar(&args.DeviceName, "name", "LEDDMX-00-000101", "Device name")
    flag.BoolVar(&args.List, "list", false, "List available patterns")
    flag.Parse()

    return args
}
