package router

import "github.com/gin-gonic/gin"

type Group struct {
	groups []*gin.RouterGroup
}

func NewGroup(groups []*gin.RouterGroup) Group {
	return Group {
		groups,
	}
}

func (g *Group) handle(method string, path string, handler gin.HandlerFunc) {
	for _, group := range g.groups {
		group.Handle(method, path, handler)
	}
}