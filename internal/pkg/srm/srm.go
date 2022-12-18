// Created by WestleyR <westleyr@nym.hush.com> on July 28, 2020
// Source code: https://github.com/WestleyR/srm
//
// Copyright (c) 2020-2022 WestleyR. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
//

package srm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"time"

	"github.com/WestleyR/srm/internal/pkg/paths"
	"github.com/c2h5oh/datasize"
	"golang.org/x/sys/unix"
)

// maxEntries is the max number of cached trash items.
const maxEntries = 3000

type SrmData struct {
	Entries        []*Entry `json:"entry"`
	LastTrashIndex int      `json:"last_index"`
}

type Entry struct {
	Path  string            `json:"path"`
	Index int               `json:"index"`
	IsDir bool              `json:"is_dir"`
	Size  datasize.ByteSize `json:"size"`
	Date  time.Time         `json:"time"`
}

type Manager struct {
	options        *Options
	Data           *SrmData
	nextTrashIndex int
}

// Options is the remove options for the file
type Options struct {
	Force          bool
	Recursive      bool
	RemoveEmptyDir bool // TODO: not used yet
}

func New(options *Options) (*Manager, error) {
	m := &Manager{}

	m.options = options
	m.Data = &SrmData{}

	b, err := os.ReadFile(paths.GetDataCacheFile())
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	if err == nil {
		err = json.Unmarshal(b, &m.Data)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal data: %w", err)
		}
		sort.Slice(m.Data.Entries, func(i, j int) bool {
			return m.Data.Entries[i].Date.Unix() < m.Data.Entries[j].Date.Unix()
		})
	}

	m.nextTrashIndex = m.getNextTrashIndex(m.Data.LastTrashIndex)

	return m, nil
}

func (m *Manager) Close() error {
	b, err := json.Marshal(m.Data)
	if err != nil {
		return err
	}

	err = os.WriteFile(paths.GetDataCacheFile(), b, 0655)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) RM(path string) error {
	if !doesPathExist(path) {
		return fmt.Errorf("%s: does not exist", path)
	}

	isDir := isPathADir(path)

	if isDir && !m.options.Recursive {
		return fmt.Errorf("%s: is a directory", path)
	}

	size := getFileSize(path)

	if isDir {
		// Is a directory
		size = getDirSize(path)

		if !m.options.Recursive {
			return fmt.Errorf("%s: is a directory", path)
		}

		if !m.options.Force {
			err := checkForWriteProtectedFileIn(path)
			if err != nil {
				return err
			}
		}
	} else {
		// Is a file
		if !m.options.Force && !checkIfFileIsWriteProtected(path) {
			return fmt.Errorf("%s: is write protected", path)
		}
	}

	trashPath := filepath.Join(paths.GetTrashDir(), fmt.Sprintf("%d", m.nextTrashIndex), filepath.Base(path))

	err := os.MkdirAll(filepath.Dir(trashPath), 0755)
	if err != nil {
		return fmt.Errorf("failed to create trash dir: %w", err)
	}

	entry := &Entry{
		Path:  trashPath,
		Index: m.nextTrashIndex,
		IsDir: isDir,
		Size:  size,
		Date:  time.Now(),
	}

	// Move the file to srm trash
	err = os.Rename(path, trashPath)
	if err != nil {
		cmd := exec.Command("mv", path, trashPath)
		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("%s: failed to move to trash directory %s", path, trashPath)
		}
	}

	m.Data.LastTrashIndex = m.nextTrashIndex

	m.Data.Entries = append(m.Data.Entries, entry)

	if len(m.Data.Entries) > maxEntries {
		//log.Printf("Warning: exeeded %v entries, removing old ones", maxEntries)
		for len(m.Data.Entries) > maxEntries {
			m.Data.Remove(0)
		}
	}

	m.nextTrashIndex = m.getNextTrashIndex(m.Data.LastTrashIndex)

	return nil
}

// Remove removes the entry from cache dir
func (d *SrmData) Remove(index int) error {
	//log.Printf("Removing entry: %s...", d.Entries[index].Path)

	err := os.RemoveAll(filepath.Dir(d.Entries[index].Path))
	if err != nil {
		return fmt.Errorf("failed to remove entry at index: %d: %w", index, err)
	}

	d.LastTrashIndex = d.Entries[index].Index

	d.Entries = append(d.Entries[:index], d.Entries[index+1:]...)

	return nil
}

// Eg. ~/.cache/srm/trash/{1,2...1000}
func (m *Manager) getNextTrashIndex(start int) int {
	start--
	if start < 0 {
		start = 0
	}
	trashNumber := start
	var i int
	for i = start; i < maxEntries+25; i++ {
		path := filepath.Join(paths.GetTrashDir(), fmt.Sprintf("%d", i))

		if isDirEmpty(path) {
			// Empty trash index, so use that
			trashNumber = i
			break
		}
	}

	if i == maxEntries+25 {
		log.Fatalf("Could not find a cache dir (please report)")
	}

	notIn := true
	for notIn {
		notIn = false
		for _, f := range m.Data.Entries {
			if f.Index == trashNumber {
				trashNumber++
				notIn = true
			}
		}
	}

	return trashNumber
}

func isDirEmpty(path string) bool {
	err := unix.Access(path, unix.F_OK)
	if errors.Is(err, os.ErrNotExist) {
		return true
	}

	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	c, err := f.Readdirnames(1)
	if errors.Is(err, io.EOF) {
		return true
	}
	return len(c) == 0
}
