package entity

import "time"

type Speed struct {
	Speed    float64
	CreateAt time.Time
}

type Speeds []Speed

type DevicesSpeeds map[string]Speeds
