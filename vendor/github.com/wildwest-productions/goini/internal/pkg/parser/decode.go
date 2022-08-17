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
)

type values struct {
	interfaceType map[string]string
	fieldNames    []string // the array containing the fields in order
	fieldTags     []string // the array containing all the fields in order
}

func getStructTags(s interface{}) *values {
	ret := &values{}
	ret.interfaceType = make(map[string]string)

	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		ret.interfaceType[field.Tag.Get("ini")] = field.Type.String()
		ret.fieldNames = append(ret.fieldNames, field.Name)
		ret.fieldTags = append(ret.fieldTags, field.Tag.Get("ini"))
	}

	return ret
}

func Unmarshal(structRet interface{}, value reflect.Value, iniMap map[string]string, subSectorName string) error {
	structFields := getStructTags(value.Interface())

	for i := 0; i < value.NumField(); i++ {
		iniValue, ok := iniMap[subSectorName+"#"+structFields.fieldTags[i]]
		//fmt.Printf("StructFieldName=%s StructTag=%s StructType=%s Ini=%s\n\n", structFields.fieldNames[i], structFields.fieldTags[i], structFields.interfaceType[structFields.fieldTags[i]], iniValue)
		if value.Field(i).Kind() != reflect.Struct {
			// Only set the value if it is found in the ini
			if ok {
				setValue(value.Field(i), structFields.interfaceType[structFields.fieldTags[i]], iniValue)
			}
		} else {
			sectorT := value.FieldByName(structFields.fieldNames[i])
			if !sectorT.IsValid() {
				return fmt.Errorf("internal error: field not found: %s", structFields.fieldNames[i])
			}
			err := Unmarshal(structRet, sectorT, iniMap, structFields.fieldTags[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}
