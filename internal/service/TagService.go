package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"mlik/internal/model"
	"mlik/internal/model/entity"
	"mlik/internal/service/internal/dao"
	"mlik/internal/service/internal/do"
)

type tagService struct {
}

var Tag = tagService{}

func (s tagService) Create(ctx context.Context, input model.TagCreateInput) (tag entity.Tags, err error) {

	id, daoError := dao.Tags.Ctx(ctx).Data(do.Tags{
		Id:      input.Id,
		TagName: input.TagName,
	}).InsertAndGetId()

	if daoError != nil {
		err = daoError
		g.Log("")
		return
	}
	g.Log().Debug(ctx, "install tag", id)

	tag.Id = id
	tag.TagName = input.TagName
	return
}
