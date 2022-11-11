package v1

import "github.com/gogf/gf/v2/frame/g"

type Ping struct {
	g.Meta `path:"/ping" method:"get" tags:"基础组件"`
}

type Pong struct {
	g.Meta `mime:"text/html" example:"string"`
}
