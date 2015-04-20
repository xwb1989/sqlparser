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
	assert.Equal(t, "ID", primary_key.ColName)

}
func TestNoPrimaryKey(t *testing.T) {
	sql := `create table t1 (
	LastName varchar(255),
	FirstName varchar(255),
	ID int unique key
)`
	tree, err := Parse(sql)
	assert.Nil(t, err)
	_, err = GetPrimaryKey(tree)
	assert.NotNil(t, err)
}

func TestChangePrimaryKeyProperty(t *testing.T) {
	sql := `create table t1 (
	LastName varchar(255),
	FirstName varchar(255),
	ID int primary key
)`

	tree, err := Parse(sql)
	assert.Nil(t, err)
	primary_key, err := GetPrimaryKey(tree)
	assert.Nil(t, err)

	primary_key.ColumnAtts = append(primary_key.ColumnAtts, "auto_increment")

	sql_actual := String(tree)

	sql_expected := `create table t1 (
	LastName varchar(255),
	FirstName varchar(255),
	ID int primary key auto_increment
)`
	assert.Equal(t, sql_expected, sql_actual)
}
