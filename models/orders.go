package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Orders struct {
	Id 						int
	Uid 					int
	GoodsId 				int
	Stage 					int
	ItemNum 				int
	UpdateTime 				time.Time
	CreateTime 				time.Time
}

func init()  {
	orm.RegisterModel(new(Orders))
}

//添加订单
func (t *Orders) Add(o orm.Ormer, order Orders) (int64, error) {
	order.UpdateTime = time.Now()
	order.CreateTime = time.Now()
	id, err := o.Insert(order)
	return id, err
}
