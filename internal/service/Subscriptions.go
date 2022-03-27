package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"mlik/internal/service/internal/dao"
)

type sSubscription struct{}

var insSubscriptions = sSubscription{}

func NewSubscription() *sSubscription {

	return &insSubscriptions
}

func (s sSubscription) Get(ctx context.Context, id int) (res gdb.Record, err error) {
	res, err = dao.Subscriptions.Ctx(ctx).Where("id", id).One()
	return
}
