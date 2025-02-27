// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"context"

	"github.com/go-openapi/runtime/middleware"

	"github.com/goharbor/harbor/src/server/v2.0/restapi/operations/ping"
)

type pingAPI struct {
	BaseAPI
}

func newPingAPI() *pingAPI {
	return &pingAPI{}
}

func (p *pingAPI) GetPing(ctx context.Context, params ping.GetPingParams) middleware.Responder {
	return ping.NewGetPingOK().WithPayload("Pong")
}
