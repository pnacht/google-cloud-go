// Copyright 2021 Google LLC
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

// [START bigquerystorage_generated_bigquery_storage_apiv1_BigQueryReadClient_SplitReadStream]

package main

import (
	"context"

	storage "cloud.google.com/go/bigquery/storage/apiv1"
	storagepb "google.golang.org/genproto/googleapis/cloud/bigquery/storage/v1"
)

func main() {
	// import storagepb "google.golang.org/genproto/googleapis/cloud/bigquery/storage/v1"

	ctx := context.Background()
	c, err := storage.NewBigQueryReadClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &storagepb.SplitReadStreamRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SplitReadStream(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END bigquerystorage_generated_bigquery_storage_apiv1_BigQueryReadClient_SplitReadStream]