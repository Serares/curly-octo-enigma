package stacks

import (
	awslambdago "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/jsii-runtime-go"
)

var BundlingOptions = &awslambdago.BundlingOptions{
	GoBuildFlags: &[]*string{jsii.String(`-ldflags "-s -w" -tags lambda.norpc`)},
}
