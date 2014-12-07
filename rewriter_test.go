package sqlparser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRewriteQuery(t *testing.T) {
	sql := "select distinct table1.* from table1 as t1"
	tree, _ := Parse(sql)

	rewriter := func(origin []byte) []byte {
		s := string(origin)
		if s == "table1" {
			s = fmt.Sprintf("%s%s%s", "_", s, "_")
		}
		return []byte(s)
	}

	Rewrite(tree, rewriter)

	expected := "select distinct _table1_.* from _table1_ as t1"
	actual := String(tree)

	assert.Equal(t, expected, actual)
}

func TestParseDDL(t *testing.T) {
	sql := "create table table1 (c1 integer primary, c2 char(8), c3 text)"
	tree, err := Parse(sql)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(String(tree))
}
