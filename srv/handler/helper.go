package handler

import (
	"fmt"
	"reflect"
)

func IsZeroValue(o interface{}) bool {
	//fmt.Println(o)

	// First check normal definitions of empty
	/*|| o == false || o == 0*/
	if reflect.DeepEqual(o, nil) || reflect.DeepEqual(o, "") || o == false || o == 0 {
		fmt.Println("is zero val")
		return true
	} else {
		return false
	}
}

func IsValidRequest(o interface{}) bool {
	if IsZeroValue(o) {
		return false
	}

	// Then see if it's a struct
	if reflect.ValueOf(o).Kind() == reflect.Struct {
		el := reflect.ValueOf(o)

		for i := 0; i < el.NumField(); i++ {
			f := el.Field(i)

			//fmt.Println(reflect.ValueOf(f).Kind())

			empty := reflect.New(f.Type())
			fmt.Println(reflect.DeepEqual(f, empty))

			//if (reflect.DeepEqual(o, empty))

			//fmt.Println(reflect.ValueOf(f).Interface())
			//fmt.Println("------------")
			//fmt.Println(IsZeroValue(reflect.ValueOf(f).Interface()))
			val := reflect.ValueOf(f).Interface()
			//fmt.Println(reflect.ValueOf(f).IsValid())
			//fmt.Println(reflect.ValueOf(f).Kind())

			//fmt.Println(val == "as")
			if IsZeroValue(val) {
				fmt.Println("??")
				return false
			}
		}
		// and create an empty copy of the struct object to compare against
		/*empty := fmt.Println(IsZeroValue(""))reflect.TypeOf(o)).Elem().Interface()

		if reflect.DeepEqual(o, empty) {
			return false
		}*/
	}
	return true
}

/*

func main() {

	type Ss struct {
		Atrr bool
	}

	type S struct {
		Email string
		Addr  string
		More  int64
		Str   Ss
	}

	test := S{
		Email: "as",
	}
	//var str string

	fmt.Println(IsValidRequest(test))
	//fmt.Println(IsZeroValue(""))
	//fmt.Println(IsZeroValue(str))

}

*/
