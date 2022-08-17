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
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

// GetIniTags returns a key-value map of the [sector#key] = value.
func GetIniTags(iniData []byte) map[string]string {
	tags := make(map[string]string)

	sector := ""

	scanner := bufio.NewScanner(bytes.NewReader(iniData))
	for scanner.Scan() {
		if scanner.Text() == "" || scanner.Text()[0] == '#' || scanner.Text()[0] == ';' {
			// Skip empty lines, and comments
			continue
		}
		if scanner.Text()[0] == '[' && scanner.Text()[len(scanner.Text())-1] == ']' {
			sector = scanner.Text()
			sector = strings.Replace(sector, "[", "", 1)
			sector = strings.Replace(sector, "]", "", 1)
			continue
		}

		values := strings.Split(scanner.Text(), " = ")
		if len(values) == 1 {
			values = strings.Split(scanner.Text(), "=")
		}

		if len(values) == 2 {
			tags[fmt.Sprintf("%s#%s", sector, values[0])] = values[1]
		}
	}

	return tags
}
