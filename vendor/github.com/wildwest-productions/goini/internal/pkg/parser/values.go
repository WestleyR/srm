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

package parser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func setValue(s reflect.Value, typeT, value string) {
	switch typeT {
	case "string":
		s.SetString(value)
	case "int":
		val, err := strconv.Atoi(value)
		if err == nil {
			s.SetInt(int64(val))
		}
	case "float64":
		val, err := strconv.ParseFloat(value, 64)
		if err == nil {
			s.SetFloat(val)
		}
	case "float32":
		val, err := strconv.ParseFloat(value, 32)
		if err == nil {
			s.SetFloat(val)
		}
	case "bool":
		switch strings.ToLower(value) {
		case "1", "true", "yes", "t", "y", "on":
			s.SetBool(true)
		default:
			s.SetBool(false)
		}
	default:
		panic(fmt.Sprintf("goini: unknown type: %s (please report this if you need this type)", typeT))
	}
}
