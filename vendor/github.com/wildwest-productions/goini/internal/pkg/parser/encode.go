//
// Created by WestleyR <westleyr@nym.hush.com> on Nov 22, 2021
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
	"bytes"
	"fmt"
	"reflect"
)

func marshaler(buf *bytes.Buffer, value reflect.Value, sector string) error {
	structFields := getStructTags(value.Interface())

	// Need to loop twice since we want to get all "default" fields first, then
	// sub-structs.

	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).Kind() != reflect.Struct && sector == "" {
			writeExp(buf, structFields.fieldTags[i], value.Field(i).Interface())
		}
	}

	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).Kind() != reflect.Struct {
			if sector != "" {
				writeExp(buf, structFields.fieldTags[i], value.Field(i).Interface())
			}
		} else {
			sectorT := value.FieldByName(structFields.fieldNames[i])
			if !sectorT.IsValid() {
				return fmt.Errorf("internal error: field not found: %s", structFields.fieldNames[i])
			}

			writeSection(buf, structFields.fieldTags[i])

			err := marshaler(buf, sectorT, structFields.fieldTags[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func writeExp(buf *bytes.Buffer, key string, val interface{}) {
	buf.WriteString(fmt.Sprintf("%s = %v", key, val))
	buf.WriteString("\n")
}

func writeSection(buf *bytes.Buffer, sect string) {
	// TODO: cat exps
	buf.WriteString("\n")
	buf.WriteString(fmt.Sprintf("[%s]", sect))
	buf.WriteString("\n")
}

func MarshalValue(value reflect.Value) ([]byte, error) {
	b := bytes.NewBuffer(nil)

	err := marshaler(b, value, "")

	return b.Bytes(), err
}
