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

	api "github.com/nlnwa/maalfrid-aggregator/maalfrid/aggregator"
)

type Api struct{
	store aggregatorStore
}

api.RegisterAggregatorServer(server, aggregator.NewApi(options)

func (a *Api) SetOption(options ...func(*Api) error) error {
	for _, opt := range options {
		if err := opt(a); err != nil {
			return err
		}
	}
	return nil
}

func NewApi(options ...func(*Api) error) (*api.AggregatorServer, error) {
	a := Api{}

}

func (a *Api) RunLanguageDetection(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {

	return new(pb.Empty), nil
}

func (a *Api) RunAggregation(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return new(pb.Empty), nil
}

func (a *Api) SyncEntities(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return new(pb.Empty), nil
}

func (a *Api) SyncSeeds(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return new(pb.Empty), nil
}
