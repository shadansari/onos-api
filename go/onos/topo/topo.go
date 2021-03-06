// Copyright 2019-present Open Networking Foundation.
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

package topo

import (
	"google.golang.org/grpc"
)

// ID ...
type ID string

// NullID ...
const NullID = ""

// Revision is an object revision
type Revision uint64

// Attribute keys
const (
	Address     = "address"
	Target      = "target"
	Type        = "type"
	Version     = "version"
	Timeout     = "timeout"
	Role        = "role"
	Displayname = "displayname"
	User        = "user"
	Password    = "password"
	TLSCaCert   = "tls-ca-cert"
	TLSCert     = "tls-cert"
	TLSKey      = "tls-key"
	TLSPlain    = "tls-plain"
	TLSInsecure = "tls-insecure"
)

// TopoClientFactory : Default EntityServiceClient creation.
var TopoClientFactory = func(cc *grpc.ClientConn) TopoClient {
	return NewTopoClient(cc)
}

// CreateTopoClient creates and returns a new topo device client
func CreateTopoClient(cc *grpc.ClientConn) TopoClient {
	return TopoClientFactory(cc)
}
