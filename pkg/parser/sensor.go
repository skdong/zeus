package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/skdong/zeus/pkg/models"
)

//<STX>Q,229,002.74,M,00,<ETX>16
var TAGLEN = 1
var FIELDNUMS = 5
var STARTTAG = string(rune(0x02))
var ENDTAG = string(rune(0x03))
var SPLIT = ","

type Parser struct {
	raw string
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) AddData(add string) {
	p.raw = p.raw + add
}

func (p *Parser) getStart() int {
	i := strings.Index(p.raw, STARTTAG)
	if i == -1 {
		return i
	}
	return i + TAGLEN
}

func (p *Parser) getEnd() int {
	i := strings.Index(p.raw, ENDTAG)
	if i == -1 {
		return i
	}
	return i - 1
}

func (p *Parser) trimNosie() {
	s := p.getStart()
	if s == -1 {
		p.raw = ""
	} else {
		p.raw = p.raw[s-TAGLEN:]
	}
}

func (p *Parser) hasEntry() bool {
	s, e := p.getStart(), p.getEnd()
	if s >= 0 && e > s {
		return true
	}
	return false
}

func (p *Parser) GetAll() (ws []*models.Wind, err error) {
	logs.Debug(p.raw)
	p.trimNosie()
	logs.Debug(p.raw)
	for p.hasEntry() {
		w, err := p.getOne()
		logs.Debug(p.raw)
		if err != nil {
			logs.Warn(err)
			continue
		}
		ws = append(ws, w)
		p.trimNosie()
		logs.Debug(p.raw, p.hasEntry())
	}
	return
}

func (p *Parser) getOne() (w *models.Wind, err error) {
	s, e := p.getStart(), p.getEnd()
	v := p.raw[s:e]
	p.raw = p.raw[e+TAGLEN+1:]
	vs := strings.Split(v, SPLIT)
	if len(vs) != FIELDNUMS {
		msg := fmt.Sprintf("Field nums is %d desire is %d",
			len(vs), FIELDNUMS)
		logs.Warn(vs)
		return w, errors.New(msg)
	}
	logs.Info(vs)
	w = new(models.Wind)
	w.DeviceId = vs[0]
	if len(vs[1]) == 0 {
		w.Direction = 0
	} else {
		w.Direction, err = strconv.Atoi(vs[1])
		if err != nil {
			logs.Warn("Direction is Not valid")
			logs.Warn(err)
			return
		}
	}

	if len(vs[2]) == 0 {
		w.Speed = 0.0

	} else {
		w.Speed, err = strconv.ParseFloat(vs[2], 64)
		if err != nil {
			logs.Warn("Speed is Not valid")
			logs.Warn(err)
			return
		}

	}
	w.Unit = vs[3]
	return
}
