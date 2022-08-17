//
// Created by WestleyR <westleyr@nym.hush.com> on Nov 20, 2021
// Source code: https://github.com/WestleyR/goini
//              https://github.com/WildWest-Productions/goini
//
// Copyright (c) 2021 WestleyR. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
//

package goini

import (
	"fmt"
	"reflect"

	"github.com/wildwest-productions/goini/internal/pkg/parser"
)

func Unmarshal(iniData []byte, s interface{}) error {
	t := reflect.ValueOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("not a struct")
	}

	iniTags := parser.GetIniTags(iniData)

	return parser.Unmarshal(s, t, iniTags, "")
}

func Marshal(s interface{}) ([]byte, error) {
	t := reflect.ValueOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return []byte{}, fmt.Errorf("not a struct")
	}

	return parser.MarshalValue(t)
}
