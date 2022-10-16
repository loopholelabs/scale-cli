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
	"github.com/loopholelabs/scale-go/scalefunc"
	"os"
	"path"
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

func (s *Storage) path(name string, tag ...string) string {
	if len(tag) > 0 {
		return path.Join(s.BaseDirectory, fmt.Sprintf("%s:%s", name, tag[0]))
	}
	return path.Join(s.BaseDirectory, name)
}

// Get returns the Scale Function with the given name and optional tag
func (s *Storage) Get(name string, tag ...string) (*scalefunc.ScaleFunc, error) {
	data, err := os.ReadFile(s.path(name, tag...))
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

// Put stores the Scale Function with the given name and optional tag
func (s *Storage) Put(name string, sf *scalefunc.ScaleFunc, tag ...string) error {
	return os.WriteFile(s.path(name, tag...), sf.Encode(), 0644)
}

// Delete removes the Scale Function with the given name and optional tag
func (s *Storage) Delete(name string, tag ...string) error {
	return os.Remove(s.path(name, tag...))
}
