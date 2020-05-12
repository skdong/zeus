package collector

import (
	"time"
)

type Buffer struct {
	Time  time.Time
	Datas []interface{}
}

func (b *Buffer) clear() {
	datas := b.Datas
	for _, d := range b.Datas {

	}
}
