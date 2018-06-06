package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Goods struct {
	Id					int
	Uid  				int
	Name 				string
	Price 				float64
	CreateTime 			time.Time
	UpdateTime 			time.Time
}

func init()  {
	orm.RegisterModel(new(Goods))
}

//新增商品
func (g *Goods) Add(o orm.Ormer, good Goods) (int64, error) {
	good.CreateTime = time.Now()
	good.UpdateTime = time.Now()
	id, err := o.Insert(good)
	return id, err
}

//编辑商品
func (g *Goods) Edit(o orm.Ormer, good Goods) (int64, error) {
	good.UpdateTime = time.Now()
	num, err := o.Update(good)
	return num, err
}