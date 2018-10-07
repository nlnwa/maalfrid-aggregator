// Copyright 2018 National Library of Norway
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aggregator

import (
	"github.com/nlnwa/maalfrid-aggregator/pkg/database"
)

type Store struct {
	*database.Rethink
}

type StoreOption func(s *Store) error

func WithDatabase(host string, port int, name string, user string, password string) StoreOption {
	return func(s *Store) error {
		db := database.New(
			database.WithAddress(host, port),
			database.WithName(name),
			database.WithCredentials(user, password))
		db.SetTags("json")
		return nil
	}
}

func (s *Store) SetOption(options ...StoreOption) error {
	for _, opt := range options {
		if err := opt(s); err != nil {
			return err
		}
	}
	return nil
}

func NewStore(options ...StoreOption) (*Store, error) {
	s := &Store{}
	err := s.SetOption(options...)
	if err != nil {
		return nil, err
	} else {
		return s, nil
	}
}

func (s *Store) SomeDbCall() (string, error) {
	if err := s.Connect(); err != nil {
		return "", err
	} else {
		defer s.Disconnect()
	}
	return "", nil
	//	if id, err := ws.Insert("result", value); err != nil {
	//		return "", err
	//	} else {
	//		return id, nil
	//	}
}
