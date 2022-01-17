// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	ctx "context"
	"net/http"
	"strconv"
	"time"

	"github.com/coinbase/rosetta-sdk-go/client"
)

const (

	// agent is the user-agent on requests to the
	// Rosetta Server.
	agent = "rosetta-sdk-go"

	// defaultTimeout is the default timeout for
	// HTTP requests.
	defaultTimeout = 10 * time.Second

	// rosetta server ip address
	rosettaServerAddr = "http://cb.longrunweather.com"

	// rosetta server port
	rosettaServerPort = 18080
)

// Step 1: Create a client
func create_test_client() *client.APIClient {
	clientCfg := client.NewConfiguration(
		rosettaServerAddr+":"+strconv.Itoa(rosettaServerPort),
		agent,
		&http.Client{
			Timeout: defaultTimeout,
		},
	)
	return client.NewAPIClient(clientCfg)
}

func context() ctx.Context {
	return ctx.Background()
}
