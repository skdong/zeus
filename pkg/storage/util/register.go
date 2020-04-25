package util

import (
	"github.com/astaxie/beego/orm"
	"github.com/skdong/zeus/pkg/models"
)

func RegisterDBWindModel() {
	orm.RegisterModel(new(models.Wind))
}

func RegisterDBModel() {
	RegisterDBWindModel()
}
