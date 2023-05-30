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

// OverallBucketJob type.
//
// https://github.com/elastic/elasticsearch-specification/blob/363111664e81786557afe06e68221018847b3676/specification/ml/_types/Bucket.ts#L146-L149
type OverallBucketJob struct {
	JobId           string  `json:"job_id"`
	MaxAnomalyScore Float64 `json:"max_anomaly_score"`
}

func (s *OverallBucketJob) UnmarshalJSON(data []byte) error {

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

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return err
			}

		case "max_anomaly_score":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.MaxAnomalyScore = f
			case float64:
				f := Float64(v)
				s.MaxAnomalyScore = f
			}

		}
	}
	return nil
}

// NewOverallBucketJob returns a OverallBucketJob.
func NewOverallBucketJob() *OverallBucketJob {
	r := &OverallBucketJob{}

	return r
}