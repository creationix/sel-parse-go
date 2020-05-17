package parse

import (
	"testing"

	. "github.com/warpfork/go-wish"

	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/traversal/selector/builder"
)

func TestParsing(t *testing.T) {
	ns := basicnode.Style__Any{}
	ssb := builder.NewSelectorSpecBuilder(ns)

	t.Run("Parse basic selector", func(t *testing.T) {
		input := "*."
		Wish(t, parseSelector(input), ShouldEqual, ssb.ExploreAll(ssb.Matcher()).Node())

	})

}
