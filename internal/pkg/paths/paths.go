// Created by WestleyR <westleyr@nym.hush.com> on 2022-05-09
// Source code: https://github.com/WestleyR/srm
//
// Copyright (c) 2022 WestleyR. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
//

package paths

import (
	"fmt"
	"os"
	"path/filepath"
)

var homeDir string = ""

func GetHome() string {
	if homeDir == "" {
		var err error
		homeDir, err = os.UserHomeDir()
		if err != nil {
			panic(fmt.Sprintf("failed to get home dir: %s", err))
		}
	}

	return homeDir
}

func GetCachePath() string {
	return filepath.Join(GetHome(), ".cache", "srm")
}

func GetDataCacheFile() string {
	return filepath.Join(GetCachePath(), "data.json")
}

func GetTrashDir() string {
	return filepath.Join(GetCachePath(), "trash")
}
