// Copyright 2020 FastWeGo
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

package job_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/guide/job"
)

func ExampleAddGuideMassendJob() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := job.AddGuideMassendJob(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGuideMassendJobList() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := job.GetGuideMassendJobList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGuideMassendJob() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := job.GetGuideMassendJob(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateGuideMassendJob() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := job.UpdateGuideMassendJob(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCancelGuideMassendJob() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := job.CancelGuideMassendJob(ctx, payload)

	fmt.Println(resp, err)
}
