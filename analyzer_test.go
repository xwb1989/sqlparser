/*
Tests for analyzer.go
*/
package sqlparser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrimaryKey(t *testing.T) {
	sql := `create table t1 (
	LastName varchar(255),
	FirstName varchar(255),
	ID int primary key
)`
	tree, err := Parse(sql)
	assert.Nil(t, err)
	primary_key, err := GetPrimaryKey(tree)
	assert.Nil(t, err)
	assert.Equal(t, "ID", primary_key)

	sql = `create table t1 (
	LastName varchar(255),
	FirstName varchar(255),
	ID int unique key
)`
	tree, err = Parse(sql)
	assert.Nil(t, err)
	_, err = GetPrimaryKey(tree)
	assert.NotNil(t, err)
}
