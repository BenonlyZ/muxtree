package service

import (
	"fmt"
	web "muxtree/pkg/v1"
)

func Home(ctx *web.Context) {
	fmt.Fprintf(ctx.W, "这是主页")
}

func User(ctx *web.Context) {
	fmt.Fprintf(ctx.W, "这是用户")
}

func CreateUser(ctx *web.Context) {
	fmt.Fprintf(ctx.W, "这是创建用户")
}

func Order(ctx *web.Context) {
	fmt.Fprintf(ctx.W, "这是订单")
}
