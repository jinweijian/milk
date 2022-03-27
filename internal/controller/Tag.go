package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "mlik/api/v1"
	"mlik/internal/model"
	"mlik/internal/service"
)

type cShow struct {
}

type cCreate struct {
}

type cUpdate struct {
}

type cDel struct {
}

type cList struct {
}

var (
	CreateTag = cCreate{}
	UpdateTag = cUpdate{}
	DeleteTag = cDel{}
	ShowTag   = cShow{}
	ListTag   = cList{}
)

func (c cCreate) Create(ctx context.Context, req *v1.TagCreateReq) (res v1.TagCreateRes, err error) {

	tag, createErr := service.Tag.Create(ctx, model.TagCreateInput{
		Id:      2,
		TagName: req.Name,
	})
	if createErr != nil {
		err = createErr
		g.Log().Debug(ctx, err.Error())
	}
	res.Tags = tag

	return
}
