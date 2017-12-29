package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PostDec struct {
	name       string
	attributes map[string]interface{}
	variable   node.Node
}

func NewPostDec(variable node.Node) node.Node {
	return PostDec{
		"PostDec",
		map[string]interface{}{},
		variable,
	}
}

func (n PostDec) Name() string {
	return "PostDec"
}

func (n PostDec) Attributes() map[string]interface{} {
	return n.attributes
}

func (n PostDec) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	v.LeaveNode(n)
}