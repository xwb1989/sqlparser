package sqlparser

import (
	"fmt"
	"reflect"
)

// Data Model
type Dish struct {
	Id      int
	Name    string
	Origin  string
	Table   TableExpr
	ColName ColName
	Query   func()
}

var typeOfBytes = reflect.TypeOf([]byte(nil))
var typeOfSQLNode = reflect.TypeOf((*SQLNode)(nil)).Elem()

func Rewrite(node SQLNode) {
	rewrite(reflect.ValueOf(node))
}
func rewrite(nodeVal reflect.Value) {
	if !nodeVal.IsValid() {
		return
	}
	nodeTyp := nodeVal.Type()
	switch nodeTyp.Kind() {
	case reflect.Slice:
		if nodeTyp == typeOfBytes {
			//the case to rewrite the val!
			fmt.Printf("Bytes: %s\n", nodeVal.Bytes())
			nodeVal.SetBytes([]byte("hey"))
		} else if nodeTyp.Implements(typeOfSQLNode) {
			for i := 0; i < nodeVal.Len(); i++ {
				m := nodeVal.Index(i)
				rewrite(m)
			}
		}
	case reflect.Struct:
		fmt.Println("Struct:", nodeTyp.Name())
		for i := 0; i < nodeVal.NumField(); i++ {
			f := nodeVal.Field(i)
			rewrite(f)
		}
	case reflect.Ptr, reflect.Interface:
		rewrite(nodeVal.Elem())
	}
}

// Example of how to use Go's reflection
// Print the attributes of a Data Model
func attributes(m interface{}) map[string]reflect.Type {
	typ := reflect.TypeOf(m)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// create an attribute data structure as a map of types keyed by a string.
	attrs := make(map[string]reflect.Type)
	// Only structs are supported so return an empty result if the passed object
	// isn't a struct
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
		return attrs
	}

	// loop through the struct's fields and set the map
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if p.Type.Kind() == reflect.Struct {
			subattr := attributes(p)
			for key, val := range subattr {
				attrs[key] = val
			}
		}
		if !p.Anonymous {
			attrs[p.Name] = p.Type
		}
	}

	return attrs
}
