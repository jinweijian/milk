package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"mlik/api/v1"
	"mlik/internal/service"
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
	sup, supErr := service.NewSubscription().Get(ctx, 2)
	if supErr != nil {
		fmt.Println("get supErr data error :", supErr.Error())
	}
	fmt.Println("supsupsupsupsupsup :", sup.Json())
	databaseConfig := g.DB("master").GetConfig().String()
	fmt.Println("databaseConfig :", databaseConfig)
	allRes, allErr := g.DB("master").Model("charges").Sum("id")
	if allErr != nil {
		fmt.Println("get model data error :", allErr.Error())
	}
	//fmt.Println("get value: ", allRes.Json())
	fmt.Println("get value: ", cacheValue.String())
	fmt.Println(g.Cfg().Get(ctx, "master"))
	glog.Debug(ctx, "这个是一个log")
	fmt.Println("path : ", glog.GetPath())
	fmt.Println("config", glog.GetLevel())
	g.RequestFromCtx(ctx).Response.Writeln(allRes)
	return
}
