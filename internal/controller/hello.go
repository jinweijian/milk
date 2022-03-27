package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"mlik/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	//g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("database.yaml")
	//gerror.New("zhege")
	// 设置一个缓存
	gCacheErr := gcache.Set(ctx, "thisKey", "value....", 0)
	if gCacheErr != nil {
		fmt.Println("cache set error :", gCacheErr.Error())
	}

	cacheValue, gCacheGetErr := gcache.Get(ctx, "thisKey")
	if gCacheGetErr != nil {
		fmt.Println("cache get error :", gCacheGetErr.Error())
	}
	databaseConfig := g.DB().GetConfig().String()
	fmt.Println("databaseConfig :", databaseConfig)
	allRes, allErr := g.DB().Model("users").Sum("id")
	if allErr != nil {
		fmt.Println("get model data error :", allErr.Error())
	}
	//fmt.Println("get value: ", allRes.Json())
	fmt.Println("get value: ", cacheValue.String())
	fmt.Println(g.Cfg().Get(ctx, "database.default"))
	glog.Debug(ctx, "这个是一个log")
	fmt.Println("path : ", glog.GetPath())
	fmt.Println("config", glog.GetLevel())
	g.RequestFromCtx(ctx).Response.Writeln(allRes)
	return
}
