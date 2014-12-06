package sqlparser

import (
	"fmt"
	"testing"
)

func TestRewrite(t *testing.T) {
	sql := "select distinct Table1.* from table1 as t1"
	tree, _ := Parse(sql)
	Rewrite(tree)
	fmt.Printf("%s\n", String(tree))
}
