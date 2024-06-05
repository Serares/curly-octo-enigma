package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	awslambdago "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func AppLambda(scope constructs.Construct, id string) {
	stack := awscdk.NewStack(scope, &id, nil)
	appLambda := awslambdago.NewGoFunction(stack, jsii.Sprintf("UndertownAdmin-%s", "DEV"), &awslambdago.GoFunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		MemorySize:   jsii.Number(1024),
		Architecture: awslambda.Architecture_ARM_64(),
		Entry:        jsii.String("../services/ssr/admin/lambda"),
		Bundling:     BundlingOptions,
		Environment:  &envVars,
		Role:         s3BucketAccessRole,
		Timeout:      awscdk.Duration_Seconds(jsii.Number(30)),
	})
}
