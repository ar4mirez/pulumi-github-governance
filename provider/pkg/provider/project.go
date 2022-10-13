package provider

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectArgs struct{}

type Project struct {
	pulumi.ResourceState
}

func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}

	component := &Project{}
	err := ctx.RegisterComponentResource("github-governance:index:Project", name, component,
		opts...)
	if err != nil {
		return nil, err
	}

	// Our code goes between these comments :D
	//_, err = github.NewRepository(ctx, name, &github.RepositoryArgs{
	//	Description: pulumi.String("Generated from automated test"),
	//	Visibility:  pulumi.String("private"),
	//})
	//if err != nil {
	//	return nil, err
	//}

	//

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{}); err != nil {
		return nil, err
	}

	return component, nil
}
