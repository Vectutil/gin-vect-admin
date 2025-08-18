package metadata

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"strconv"
	"time"
)

func SetMetadataForUserInfo(c *gin.Context) {

	value, _ := c.Get("user")
	userInfo := value.(*MetaData)

	// 将用户Id和租户Id添加到请求上下文 (按需添加)
	c.Request = c.Request.WithContext(SetMDToIncoming(c.Request.Context(), UserName, userInfo.Username))
	c.Request = c.Request.WithContext(SetMDToIncoming(c.Request.Context(), UserId, strconv.FormatInt(userInfo.Id, 10)))
	c.Request = c.Request.WithContext(SetMDToIncoming(c.Request.Context(), TenantId, strconv.FormatInt(userInfo.TenantId, 10)))
}

func SetMDToIncoming(ctx context.Context, k, v string) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)
	newMd := metadata.MD{}
	for key, value := range md {
		newMd[key] = value
	}
	newMd.Set(k, v)
	return metadata.NewIncomingContext(ctx, newMd)
}

func GetMD(ctx context.Context, k string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	v := md.Get(k)
	if len(v) <= 0 {
		return ""
	}

	return v[0]
}

type MetaData struct {
	Id          int64     `json:"id"`          // 主键
	Username    string    `json:"username"`    // 用户名
	Password    string    `json:"-"`           // 密码
	FullName    string    `json:"fullName"`    // 全名
	Avatar      string    `json:"avatar"`      // 头像URL
	Email       string    `json:"email"`       // 邮箱
	Phone       string    `json:"phone"`       // 手机号
	DeptId      int64     `json:"deptId"`      // 所属主部门Id
	Status      int       `json:"status"`      // 状态：1启用 0禁用
	LoginCount  int       `json:"loginCount"`  // 登录次数
	LastLoginAt int64     `json:"lastLoginAt"` // 最后登录时间
	LastLoginIP string    `json:"lastLoginIp"` // 最后登录IP地址
	TenantId    int64     `json:"tenantId"`    // 租户Id
	OrgId       int64     `json:"orgId"`       // 组织Id
	Remark      string    `json:"remark"`      // 备注信息
	CreatedAt   time.Time `json:"createdAt"`   // 创建时间
	CreatedBy   int64     `json:"createdBy"`   // 创建人Id
	UpdatedAt   time.Time `json:"updatedAt"`   // 更新时间
	UpdatedBy   int64     `json:"updatedBy"`   // 更新人Id
}
