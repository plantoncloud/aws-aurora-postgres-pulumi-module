package pkg

import (
	awsrdsclusterv1 "buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/provider/aws/awsrdscluster/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	AwsRdsCluster *awsrdsclusterv1.AwsRdsCluster
	Labels        map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *awsrdsclusterv1.AwsRdsClusterStackInput) *Locals {
	locals := &Locals{}

	//assign value for the locals variable to make it available across the project
	locals.AwsRdsCluster = stackInput.Target

	return locals
}
