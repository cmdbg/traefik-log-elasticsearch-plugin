// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/363111664e81786557afe06e68221018847b3676

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// DeprecationIndexing type.
//
// https://github.com/elastic/elasticsearch-specification/blob/363111664e81786557afe06e68221018847b3676/specification/nodes/info/types.ts#L140-L142
type DeprecationIndexing struct {
	Enabled string `json:"enabled"`
}

func (s *DeprecationIndexing) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "enabled":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Enabled = o

		}
	}
	return nil
}

// NewDeprecationIndexing returns a DeprecationIndexing.
func NewDeprecationIndexing() *DeprecationIndexing {
	r := &DeprecationIndexing{}

	return r
}