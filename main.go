package main

import (
    "fmt"
    "os"
    "time"
    "strconv"
)

const commandDelay = 1000 * time.Millisecond

func main() {
    args := ParseArguments()

    // Handle special commands first
    if args.Command != "" {
        switch args.Command {
        case "check":
            handleCheckCommand(args)
        case "setmac":
            handleSetMacCommand(args)
        case "setname":
            handleSetNameCommand(args)
        case "listpatterns":
            ListPatterns()
        default:
            fmt.Println("Unknown command:", args.Command)
        }
        return
    }

    if args.List {
        ListPatterns()
        return
    }

    err := ValidateArguments(args)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    bm := NewBluetoothManager(args.DeviceMAC, args.DeviceName)

    enabled, err := bm.IsBluetoothEnabled()
    if err != nil || !enabled {
        fmt.Println("Bluetooth is not enabled...")
        os.Exit(1)
    }

    found, err := bm.FindDevice()
    if err != nil || !found {
        fmt.Println("Device not found in range. Please make sure the device is on and in range.")
        fmt.Println("Sometimes it helps to open your bluetooth GUI manually")
        os.Exit(1)
    }

    err = bm.Connect()
    if err != nil {
        fmt.Println("Failed to connect to the device:", err)
        fmt.Println("Sometimes it helps to open your bluetooth GUI manually")
        os.Exit(1)
    }
    defer bm.Disconnect()

    if args.Power != "" {
        powerOn := args.Power == "on"
        data := MakePowerData(powerOn)
        err = bm.SendCommand(data)
        if err != nil {
            fmt.Println("Failed to send power command:", err)
            fmt.Println("Sometimes it helps to open your bluetooth GUI manually")
            os.Exit(1)
        }
        time.Sleep(commandDelay)
    }

    if args.Color != "" {
        c, err := ParseColor(args.Color)
        if err != nil {
            fmt.Println("Invalid color value:", err)
            os.Exit(1)
        }
        data := MakeColorData(c)
        err = bm.SendCommand(data)
        if err != nil {
            fmt.Println("Failed to send color command:", err)
            os.Exit(1)
        }
        time.Sleep(commandDelay)
    }

    if args.Pattern != "" {
        patternIndex, err := strconv.Atoi(args.Pattern)
        if err != nil {
            patternIndex = GetPatternIndex(args.Pattern)
            if patternIndex == -1 {
                fmt.Println("Invalid pattern value:", args.Pattern)
                os.Exit(1)
            }
        }
        data := MakePatternData(patternIndex)
        err = bm.SendCommand(data)
        if err != nil {
            fmt.Println("Failed to send pattern command:", err)
            os.Exit(1)
        }
        time.Sleep(commandDelay)
    }

    if args.Brightness != -1 {
        data := MakeBrightnessData(args.Brightness)
        err = bm.SendCommand(data)
        if err != nil {
            fmt.Println("Failed to send brightness command:", err)
            os.Exit(1)
        }
        time.Sleep(commandDelay)
    }

    if args.Speed != -1 {
        // TODO: implement speed command
        fmt.Println("Speed command not implemented yet.")
    }

    fmt.Println("Commands sent successfully!")
}

func handleCheckCommand(args Arguments) {
    bm := NewBluetoothManager(args.DeviceMAC, args.DeviceName)
    found, err := bm.FindDevice()
    if err != nil || !found {
        fmt.Println("Device not found in range.")
    } else {
        fmt.Println("Device is in range.")
    }
}

func handleSetMacCommand(args Arguments) {
    fmt.Println("Set MAC address functionality not implemented.")
}

func handleSetNameCommand(args Arguments) {
    fmt.Println("Set device name functionality not implemented.")
}
