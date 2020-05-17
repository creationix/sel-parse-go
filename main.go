package main

import (
	"fmt"

	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/traversal/selector/builder"

	"github.com/kr/pretty"
)

func main() {
	ns := basicnode.Style__Any{}
	ssb := builder.NewSelectorSpecBuilder(ns)
	node := ssb.ExploreAll(ssb.Matcher()).Node()
	fmt.Printf("%# v\n", pretty.Formatter(node))
}
