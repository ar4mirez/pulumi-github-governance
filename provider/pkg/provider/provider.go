// Copyright 2016-2021, Pulumi Corporation.
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

package provider

import (
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
)

func construct(ctx *pulumi.Context, typ, name string, inputs provider.ConstructInputs,
	options pulumi.ResourceOption) (*provider.ConstructResult, error) {
	// TODO: Add support for additional component resources here.
	switch typ {
	case "github-governance:index:Project":
		return constructProject(ctx, name, inputs, options)
	default:
		return nil, errors.Errorf("unknown resource type %s", typ)
	}
}

func constructProject(ctx *pulumi.Context, name string, inputs provider.ConstructInputs,
	options pulumi.ResourceOption) (*provider.ConstructResult, error) {

	args := &ProjectArgs{}
	if err := inputs.CopyTo(args); err != nil {
		return nil, errors.Wrap(err, "setting args")
	}

	project, err := NewProject(ctx, name, args, options)
	if err != nil {
		return nil, errors.Wrap(err, "creating component")
	}

	return provider.NewConstructResult(project)
}
