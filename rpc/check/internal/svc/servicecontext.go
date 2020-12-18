package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"

	"go-zero-demo/rpc/check/internal/config"
	"go-zero-demo/rpc/model"
)

type ServiceContext struct {
	c     config.Config
	Model model.BookModel // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:     c,
		Model: model.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动代码
	}
}
