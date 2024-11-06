package main

import (
    "errors"
    "fmt"
    "time"

    "github.com/muka/go-bluetooth/api"
    "github.com/muka/go-bluetooth/bluez/profile/device"
)

// BluetoothManager handles Bluetooth interactions
type BluetoothManager struct {
    AdapterID     string
    DeviceAddress string
    DeviceName    string
    Device        *device.Device1
    CharUUID      string
}

// NewBluetoothManager creates a new BluetoothManager instance
func NewBluetoothManager(mac, name string) *BluetoothManager {
    return &BluetoothManager{
        AdapterID:     "hci0",
        DeviceAddress: mac,
        DeviceName:    name,
        CharUUID:      "0000ffe1-0000-1000-8000-00805f9b34fb", // UUID for characteristic
    }
}

func (bm *BluetoothManager) IsBluetoothEnabled() (bool, error) {
    a, err := api.GetAdapter(bm.AdapterID)
    if err != nil {
        return false, err
    }
    poweredVariant, err := a.GetProperty("Powered")
    if err != nil {
        return false, err
    }
    powered, ok := poweredVariant.Value().(bool)
    if !ok {
        return false, fmt.Errorf("failed to assert 'Powered' property to bool")
    }
    return powered, nil
}

func (bm *BluetoothManager) FindDevice() (bool, error) {
    a, err := api.GetAdapter(bm.AdapterID)
    if err != nil {
        return false, err
    }

    devices, err := a.GetDevices()
    if err != nil {
        return false, err
    }

    for _, dev := range devices {
        props, err := dev.GetProperties()
        if err != nil {
            continue
        }
        if props.Address == bm.DeviceAddress || props.Name == bm.DeviceName {
            bm.Device = dev
            fmt.Printf("Selected device: Address=%s Name=%s\n", props.Address, props.Name)
            return true, nil
        }
    }
    return false, nil
}


func (bm *BluetoothManager) Connect() error {
    if bm.Device == nil {
        return errors.New("device not found")
    }
    deviceName := bm.Device.Properties.Name
    if deviceName == "" {
        deviceName = bm.Device.Properties.Alias
    }
    fmt.Println("Attempting to connect to device:", deviceName)
    err := bm.Device.Connect()
    if err != nil {
        fmt.Println("Failed to connect to device:", err)
        return err
    }
    // Wait for services to be resolved
    err = bm.waitForServicesResolved(5 * time.Second)
    if err != nil {
        fmt.Println("Failed while waiting for services to be resolved:", err)
        return err
    }
    fmt.Println("Connected to device successfully.")
    return nil
}


func (bm *BluetoothManager) SendCommand(data []byte) error {
    fmt.Printf("Attempting to send data: % X\n", data)

    char, err := bm.Device.GetCharByUUID(bm.CharUUID)
    if err != nil {
        fmt.Println("Failed to get characteristic:", err)
        return err
    }

    // Dont pass the write-without-response-option.
    // The signature requires it but when it is passed the device wont respond
    err = char.WriteValue(data, nil)
    if err != nil {
        fmt.Println("Failed to write value:", err)
        return err
    }
    fmt.Println("Command sent successfully via Bluetooth.")
    return nil
}


func (bm *BluetoothManager) Disconnect() {
    if bm.Device != nil {
        bm.Device.Disconnect()
    }
}

func (bm *BluetoothManager) waitForServicesResolved(timeout time.Duration) error {
    start := time.Now()
    for {
        props, err := bm.Device.GetProperties()
        if err != nil {
            return err
        }
        if props.ServicesResolved {
            return nil
        }
        if time.Since(start) > timeout {
            return errors.New("timeout waiting for services to be resolved")
        }
        time.Sleep(1000 * time.Millisecond)
    }
}

