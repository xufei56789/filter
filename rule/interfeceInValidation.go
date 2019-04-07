package rule

import (
	"reflect"
)

type InterfaceInValidator func(interface{}, interface{}) bool

//type StringValueValidator func(v string, slice []string) bool

func NewInterfaceInValidator(v interface{}, validator InterfaceInValidator, message string) Rule {
	return Rule{
		Message:           message,
		InterfaceValidate: validator,
		InterfaceValue:    v,
	}
}


func contains(v interface{}, i interface{}) bool {
	if index(v, i) > -1 {
		return true
	}
	return false
}

// 反射值遍历查询
// iki: 5,[0,1,2,3,4,5,6,7,8,9] => 5
func index(v interface{}, i interface{}) int {
	index := -1
	switch reflect.TypeOf(i).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(i)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(v, s.Index(i).Interface()) {
				index = i
				break
			}
		}
	}
	return index
}
