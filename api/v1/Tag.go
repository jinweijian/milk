package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"mlik/internal/model/entity"
)

type list struct {
	Status int
	Total  int
	Data   []entity.Tags
}

// TagCreateReq 创建Tag请求
type TagCreateReq struct {
	g.Meta `path:"/" tags:"tags" method:"post" summary:"add tag api"`
	Name   string `v:"required"`
}

// TagCreateRes 创建Tag返回
type TagCreateRes struct {
	g.Meta `mime:"application/json" example:"string"`
	entity.Tags
}

// TagUpdateReq 修改tag请求
type TagUpdateReq struct {
	g.Meta `path:"/" tags:"tags" method:"put" summary:"update tag api"`
	Id     int    `v:"required"`
	Name   string `v:"required"`
}

// TagUpdateRes 修改tag返回
type TagUpdateRes struct {
	g.Meta `mime:"application/json" example:"string"`
	entity.Tags
}

// TagDeleteReq 删除tag请求
type TagDeleteReq struct {
	g.Meta `path:"/{id}" tags:"tags" method:"delete" summary:"show tag api"`
}

// TagDeleteRes 删除tag返回
type TagDeleteRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Status int
}

type TagShowReq struct {
	g.Meta `path:"/{id}" tags:"tags" method:"get" summary:"show tag api"`
}

type TagShowRes struct {
	g.Meta `mime:"application/json" example:"string"`
	entity.Tags
}

type TagListReq struct {
	g.Meta `path:"/" tags:"tags" method:"get" summary:"tag list api"`
}

type TagListRes struct {
	g.Meta `mime:"application/json" example:"string"`
	list
}
