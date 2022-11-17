/*
	Copyright 2022 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

// Package storage is used to store and retrieve built Scale Functions
package storage

import (
	"fmt"
	"github.com/loopholelabs/scale/scalefunc"
	"os"
	"path"
	"strings"
)

var (
	Default *Storage
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	Default, err = New(path.Join(homeDir, ".config", "scale", "functions"))
	if err != nil {
		panic(err)
	}
}

// Storage is used to store and retrieve built Scale Functions
type Storage struct {
	BaseDirectory string
}

func New(baseDirectory string) (*Storage, error) {
	err := os.MkdirAll(baseDirectory, 0755)
	if err != nil {
		return nil, err
	}

	return &Storage{
		BaseDirectory: baseDirectory,
	}, nil
}

func (s *Storage) path(name string) string {
	return fmt.Sprintf("%s.scale", path.Join(s.BaseDirectory, name))
}

// Get returns the Scale Function with the given name
func (s *Storage) Get(name string) (*scalefunc.ScaleFunc, error) {
	data, err := os.ReadFile(s.path(name))
	if err != nil {
		return nil, err
	}

	sf := new(scalefunc.ScaleFunc)
	err = sf.Decode(data)
	if err != nil {
		return nil, err
	}

	return sf, nil
}

func (s *Storage) get(name string) (*scalefunc.ScaleFunc, error) {
	data, err := os.ReadFile(path.Join(s.BaseDirectory, name))
	if err != nil {
		return nil, err
	}

	sf := new(scalefunc.ScaleFunc)
	err = sf.Decode(data)
	if err != nil {
		return nil, err
	}

	return sf, nil
}

// Copy copies the Scale Function with the given name to the given destination
func (s *Storage) Copy(name string, destination string) (string, error) {
	data, err := os.ReadFile(s.path(name))
	if err != nil {
		return "", err
	}

	if !strings.HasSuffix(destination, ".scale") {
		destination = fmt.Sprintf("%s.scale", destination)
	}

	err = os.WriteFile(destination, data, 0644)
	if err != nil {
		return "", err
	}

	return destination, nil
}

// List returns all the Scale Functions stored in the storage
func (s *Storage) List() ([]*scalefunc.ScaleFunc, error) {
	entries, err := os.ReadDir(s.BaseDirectory)
	if err != nil {
		return nil, fmt.Errorf("failed to read base directory: %w", err)
	}
	var scaleFuncEntries []*scalefunc.ScaleFunc
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		scaleFunc, err := s.get(entry.Name())
		if err != nil {
			return nil, fmt.Errorf("failed to get scale function %s: %w", entry.Name(), err)
		}
		scaleFuncEntries = append(scaleFuncEntries, scaleFunc)
	}
	return scaleFuncEntries, nil
}

// Put stores the Scale Function with the given name and optional tag
func (s *Storage) Put(name string, sf *scalefunc.ScaleFunc) error {
	return os.WriteFile(s.path(name), sf.Encode(), 0644)
}

// Delete removes the Scale Function with the given name and optional tag
func (s *Storage) Delete(name string) error {
	return os.Remove(s.path(name))
}
