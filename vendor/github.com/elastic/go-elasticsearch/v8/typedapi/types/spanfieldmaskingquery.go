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
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// SpanFieldMaskingQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/363111664e81786557afe06e68221018847b3676/specification/_types/query_dsl/span.ts#L30-L33
type SpanFieldMaskingQuery struct {
	Boost      *float32   `json:"boost,omitempty"`
	Field      string     `json:"field"`
	Query      *SpanQuery `json:"query,omitempty"`
	QueryName_ *string    `json:"_name,omitempty"`
}

func (s *SpanFieldMaskingQuery) UnmarshalJSON(data []byte) error {

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

		case "boost":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return err
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.QueryName_ = &o

		}
	}
	return nil
}

// NewSpanFieldMaskingQuery returns a SpanFieldMaskingQuery.
func NewSpanFieldMaskingQuery() *SpanFieldMaskingQuery {
	r := &SpanFieldMaskingQuery{}

	return r
}