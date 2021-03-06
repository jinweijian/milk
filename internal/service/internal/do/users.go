// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table t_users for DAO operations like Where/Data.
type Users struct {
	g.Meta    `orm:"table:t_users, do:true"`
	Id        interface{} //
	OpenId    interface{} // 用户的标识，对当前公众号唯一
	Nikename  interface{} // 用户名称
	Username  interface{} // 用户名称
	Sex       interface{} // 性别； 0 - 未知；1 - 男 ；2 - 女
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
