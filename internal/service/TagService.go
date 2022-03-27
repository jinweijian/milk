package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"mlik/internal/model"
	"mlik/internal/model/entity"
	"mlik/internal/service/internal/dao"
)

type tagService struct {
}

var Tag = tagService{}

func (s tagService) Create(ctx context.Context, input model.TagCreateInput) (tag entity.Tags, err error) {

	id, daoError := dao.Tags.DB().Ctx(ctx).Model(dao.Tags.Table()).InsertAndGetId(input)

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
