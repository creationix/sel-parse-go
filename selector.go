package selector

import (
	"github.com/ipld/go-ipld-prime"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/traversal/selector/builder"
)

// ParseSyntax an IPLD selector syntax string into a node
func ParseSyntax(selector string) ipld.Node {
	ns := basicnode.Style__Any{}
	ssb := builder.NewSelectorSpecBuilder(ns)

	if selector == "*." {
		return ssb.ExploreAll(ssb.Matcher()).Node()
	}
	return ipld.Null
}
