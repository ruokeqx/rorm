package main

import (
	"fmt"
	"reflect"
)

type Account struct {
	Username string
	Password string
}

func reflect_test() {
	// reflect.ValueOf() 获取指针对应的反射值
	// reflect.Indirect() 获取指针指向的对象的反射值
	typ := reflect.Indirect(reflect.ValueOf(&Account{})).Type()
	// fmt.Println(typ.String())	main.Account
	fmt.Println(typ.Name()) // Account
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)   // StructField
		fmt.Println(field.Name) // Username Password
	}
}
