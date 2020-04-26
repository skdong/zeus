package parser

import (
	"testing"

	"github.com/astaxie/beego/logs"
	"github.com/skdong/zeus/pkg/models"
)

func TestParser(t *testing.T) {
	raw := []rune{0x02, 0x51, 0x2C, 0x2C,
		0x30, 0x30, 0x30, 0x2E, 0x30,
		0x32, 0x2C, 0x4D, 0x2C, 0x30, 0x30, 0x2C,
		0x03, 0x32, 0x43, 0x0D, 0x0A}
	data := string(raw) + string(raw) + string(raw)
	p := NewParser()
	p.AddData(data)
	desire := []*models.Wind{
		models.NewWind("Q", 0, 0.02, "M"),
		models.NewWind("Q", 0, 0.02, "M"),
		models.NewWind("Q", 0, 0.02, "M"),
	}
	real, err := p.GetAll()
	if err != nil || len(real) != len(desire) {
		for _, w := range real {
			logs.Warn(w.ToString())
		}
		t.Fatalf("real lean %d, desire len %d", len(real), len(desire))
	}
	for i := range real {
		if !real[i].Equal(desire[i]) {
			logs.Warn("desire: ", desire[i].ToString())
			logs.Warn("real :", real[i].ToString())
			t.Fatal("Parser err")

		}
	}

}
