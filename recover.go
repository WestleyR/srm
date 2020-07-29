// Created by: WestleyR
// Email: westleyr@nym.hush.com
// Url: https://github.com/WestleyR/srm
// Last modified date: 2020-07-29
//
// This file is licensed under the terms of
//
// The Clear BSD License
//
// Copyright (c) 2019-2020 WestleyR
// All rights reserved.
//
// This software is licensed under a Clear BSD License.
//

package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "path/filepath"
  "strconv"
)

func recoverFileFromTrashIndex(trashIndex int) error {
  path := getCachePath()

  trashPath := filepath.Join(path, strconv.Itoa(trashIndex))
  files, err := ioutil.ReadDir(trashPath)
  if err != nil {
    fmt.Println(err)
    return err
  }

  if (len(files) == 0) {
    return fmt.Errorf("no files at index: %d", trashIndex)
  }
  
  fullTrashFile := filepath.Join(trashPath, files[0].Name())

  err = os.Rename(fullTrashFile, files[0].Name())
  if err != nil {
  	return err
  }

  fmt.Printf("File %s has been recovered to the current directory\n", files[0].Name())

  return nil
}

