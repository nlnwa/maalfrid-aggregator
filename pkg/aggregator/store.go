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
	"time"
	"context"

	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
	"github.com/pkg/errors"

	"github.com/nlnwa/maalfrid-aggregator/pkg/database"
	"github.com/nlnwa/maalfrid-aggregator/pkg/types"
)

type Store struct {
	db *database.Rethink
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

func (s *Store) OpenStore() error {
	return s.db.Connect()
}

func (s *Store) CloseStore() error {
	return s.db.Disconnect()
}

func (s *Store) StartLanguageDetection() (string, error) {
	logEntry := &types.LanguageDetectionLogEntry{
		StartTime: time.Now(),
		Type:      "languageDetection",
	}
	if id, err := s.db.Insert("system", logEntry); err != nil {
		return "", errors.Wrap(err, "failed to insert language detection log entry")
	} else {
		return id, nil
	}
}

func (s *Store) EndLanguageDetection(id string) error {
	logEntry := &types.LanguageDetectionLogEntry{
		EndTime: time.Now(),
	}
	if err := s.db.Update("system", id, logEntry); err != nil {
		return errors.Wrap(err, "failed to update language detection log entry")
	} else {
		return nil
	}
}

func (s *Store) IsLanguageDetectionInProgress() (bool, error) {
	table := "system"
	cursor, err := r.
		Table(table).
		Filter(func(doc r.Term) r.Term {
			return doc.Field("type").Eq("languageDetection")
		}).
		HasFields("endTime").Not().
		Run(s.db.Session)
	if err != nil {
		return false, errors.Wrap(err, "failed to get cursor to language detection log entries without endTime")
	}
	logEntry := &types.LanguageDetectionLogEntry{}
	return cursor.Peek(logEntry)
}

func (s *Store) GetCursorToTextsMissingLanguageField(ctx context.Context) (*r.Cursor, error) {

	database := "veidemann"
	table := "extracted_text"
	cursor, err := r.DB(database).Table(table).
		Filter(func(text r.Term) r.Term {
			return text.HasFields("language").Not()
		}).
		Run(s.db)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to get cursor to: %s.%s", database, table)
	} else {
		return cursor, nil
	}
}
