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
	"context"

	pb "github.com/golang/protobuf/ptypes/empty"
)

type AggregatorApi struct {
	store *Store
}

type AggregatorOption func (a* AggregatorApi) error

func WithStore(store *Store) AggregatorOption {
	return func(a *AggregatorApi) error {
		a.store = store
		return nil
	}
}

func (a *AggregatorApi) SetOption(options ...AggregatorOption) error {
	for _, opt := range options {
		if err := opt(a); err != nil {
			return err
		}
	}
	return nil
}

func NewApi(options ...AggregatorOption) (*AggregatorApi, error) {
	a := &AggregatorApi{}
	err := a.SetOption(options...)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AggregatorApi) RunLanguageDetection(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return new(pb.Empty), nil
}

func (a *AggregatorApi) RunAggregation(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return new(pb.Empty), nil
}

func (a *AggregatorApi) SyncEntities(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return new(pb.Empty), nil
}

func (a *AggregatorApi) SyncSeeds(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return new(pb.Empty), nil
}
