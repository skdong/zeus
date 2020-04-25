package parser

import (
	"testing"

	"github.com/astaxie/beego/logs"
	"github.com/skdong/zeus/pkg/models"
)

func TestParser(t *testing.T) {
	data := `<STX>Q,229,002.74,M,00,<ETX>16
			 <STX>Q,22,002.74,M,00,<ETX>16
			 <STX>Q,229,002.74,M,00,<ETX>16`
	p := NewParser()
	p.AddData(data)
	desire := []*models.Wind{
		models.NewWind("Q", 229, 2.74, "M"),
		models.NewWind("Q", 229, 2.74, "M"),
		models.NewWind("Q", 229, 2.74, "M"),
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
