package main

import (
    "image/color"
)

// turn the lights on or off
func MakePowerData(powerOn bool) []byte {
    var powerByte byte
    if powerOn {
        powerByte = 0x03
    } else {
        powerByte = 0x02
    }
    return []byte{
        0x7B,
        0xFF,
        0x04,
        powerByte,
        0xFF,
        0xFF,
        0xFF,
        0xFF,
        0xBF,
    }
}

// set the light color
func MakeColorData(c color.RGBA) []byte {
    return []byte{
        0x7B,
        0xFF,
        0x07,
        c.R,
        c.G,
        c.B,
        0x00,
        0xFF,
        0xBF,
    }
}

// byte array to set the brightness
func MakeBrightnessData(brightness int) []byte {
    brightnessPercentage := clamp(brightness, 0, 100)
    adjustedPercentage := (brightnessPercentage * 32) / 100
    return []byte{
        0x7B,
        0xFF,
        0x01,
        byte(adjustedPercentage),
        byte(brightnessPercentage),
        0x00,
        0xFF,
        0xFF,
        0xBF,
    }
}

// create a byte array to set a pattern by its index
func MakePatternData(patternIndex int) []byte {
    index := clamp(patternIndex, 0, 210)
    return []byte{
        0x7B,
        0xFF,
        0x03,
        byte(index),
        0xFF,
        0xFF,
        0xFF,
        0xFF,
        0xBF,
    }
}

// restrict a value within a range
func clamp(value, min, max int) int {
    if value < min {
        return min
    }
    if value > max {
        return max
    }
    return value
}
