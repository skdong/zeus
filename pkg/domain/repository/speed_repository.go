package repository

import (
	"time"

	"github.com/skdong/zeus/pkg/domain/entity"
)

type SpeedRepository interface {
	GetDevicesSpeeds(start, end time.Time, interval time.Duration) entity.DevicesSpeeds
}
