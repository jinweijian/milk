// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SubscriptionPeriods is the golang structure of table subscription_periods for DAO operations like Where/Data.
type SubscriptionPeriods struct {
	g.Meta           `orm:"table:subscription_periods, do:true"`
	Id               interface{} //
	SubscriptionNo   interface{} // 订阅单号
	Uid              interface{} // 用户ID
	App              interface{} // 应用标识
	Package          interface{} // app包名称
	Channel          interface{} // 支付渠道
	Version          interface{} // 订阅时app版本号
	Sku              interface{} // 商品ID
	PurchaseToken    interface{} // 用户订阅的token
	TransactionNo    interface{} // 续订交易号
	Amount           interface{} // 订单金额
	CreatedBy        interface{} // 创建资源的用户标识
	NotificationType interface{} // 通知状态
	Type             interface{} // 状态 1. 新订阅； 2.订阅成功续费 3. 订阅已过期； 4. 订阅已暂停； 5： 订阅已撤销
	Credential       interface{} // json交易凭证
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
	DeletedAt        *gtime.Time // 删除时间
	StartedAt        *gtime.Time // 订阅开始时间
	ExpiredAt        *gtime.Time // 订阅到期时间
}