package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ipld/go-ipld-schema/parser"
	"github.com/kr/pretty"
)

func parse(path string) error {
	fi, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fi.Close()

	s := bufio.NewScanner(fi)
	sch, err := parser.ParseSchema(s)
	if err != nil {
		return err
	}

	fmt.Printf("%# v", pretty.Formatter(sch))

	return nil
}

func main() {
	parse("selector.ipldsch")

}
