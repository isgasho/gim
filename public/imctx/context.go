package imctx

import "context"

type Context struct {
	Ctx context.Context
}

func NewContext(ctx context.Context) *Context {
	return &Context{Ctx: ctx}
}
