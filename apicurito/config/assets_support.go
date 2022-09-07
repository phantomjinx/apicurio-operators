/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// AssetAsString returns the named resource content as string.
func AssetAsString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// Asset provides an easy way to access to embedded assets.
func Asset(name string) ([]byte, error) {
	name = strings.Trim(name, " ")
	if strings.HasPrefix(name, "/") {
		name = strings.TrimPrefix(name, "/")
	}

	file, err := openAsset(name)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot access resource file %s", name)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		_ = file.Close()
		return nil, errors.Wrapf(err, "cannot access resource file %s", name)
	}

	return data, file.Close()
}

// DirExists tells if a directory exists and can be listed for files.
func DirExists(dirName string) bool {
	if _, err := openAsset(dirName); err != nil {
		return false
	}
	return true
}

// CloseQuietly unconditionally close an io.Closer
// It should not be used to replace the Close statement(s).
func closeQuietly(closer io.Closer) {
	_ = closer.Close()
}

// Assets lists all file names in the given path (starts with '/').
func Assets(dirName string) ([]string, error) {
	dirName = strings.Trim(dirName, " ")
	if strings.HasPrefix(dirName, "/") {
		dirName = strings.TrimPrefix(dirName, "/")
	}

	dir, err := openAsset(dirName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}

		return nil, errors.Wrapf(err, "error while listing resource files %s", dirName)
	}

	info, err := dir.Stat()
	if err != nil {
		return nil, dir.Close()
	}
	if !info.IsDir() {
		closeQuietly(dir)
		return nil, errors.Wrapf(err, "location %s is not a directory", dirName)
	}

	files, err := EFS.ReadDir(dirName)
	if err != nil {
		closeQuietly(dir)
		return nil, errors.Wrapf(err, "error while listing files on directory %s", dirName)
	}

	var res []string
	for _, f := range files {
		if !f.IsDir() {
			res = append(res, filepath.Join(dirName, f.Name()))
		}
	}

	return res, dir.Close()
}

func openAsset(path string) (fs.File, error) {
	return EFS.Open(filepath.ToSlash(path))
}
