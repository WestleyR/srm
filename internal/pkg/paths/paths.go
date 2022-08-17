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
	"time"
)

var homeDir string = ""

func GetHome() string {
	if homeDir == "" {
		var err error
		homeDir, err = os.UserHomeDir()
		if err != nil {
			panic(fmt.Sprintf("failed to get home dir (please report): %s", err))
		}
	}

	return homeDir
}

func GetCachePath() string {
	return filepath.Join(GetHome(), ".cache", "srm")
}

func GetDataCachePath() string {
	return filepath.Join(GetCachePath(), "data-cache.ini")
}

func getTrashDir() string {
	return filepath.Join(GetCachePath(), "trash")
}

func GetTrashDirPath() string {
	// TODO: may not need to get these values every time, could store them like homeDir.
	year, month, _ := time.Now().UTC().Date()

	return filepath.Join(getTrashDir(), fmt.Sprintf("%d", year), fmt.Sprintf("%d", month))
}
