package main

import (
	gg "github.com/ar4mirez/pulumi-github-governance/sdk/go/github-governance"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		_, err := gg.NewProject(ctx, "myawesomeproject", &gg.ProjectArgs{}, nil)
		if err != nil {
			return errors.Wrap(err, "creating my project :(")
		}

		return nil
	})
}
