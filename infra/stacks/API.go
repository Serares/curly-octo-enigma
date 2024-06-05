package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type APIStackProps struct {
	awscdk.StackProps
	Env string
}

func API(scope constructs.Construct, id string, props *APIStackProps) {
	var authorizationHeader = "Authorization"
	stack := awscdk.NewStack(scope, &id, &props.StackProps)
	// Define the API Gateway
	rootResource := awsapigateway.NewRestApi(stack, jsii.Sprintf("BlogPageApi-%s", props.Env), &awsapigateway.RestApiProps{
		RestApiName:    jsii.Sprintf("BlogPageApi-%s", props.Env),
		Deploy:         jsii.Bool(false),
		CloudWatchRole: jsii.Bool(true),
	})

	apiResource := rootResource.Root().AddResource(jsii.String("api"), nil)

	v1Resource := apiResource.AddResource(jsii.String("v1"), nil)

	questionsResource := v1Resource.AddResource(jsii.String("questions"), nil)

	questionResource := questionsResource.AddResource(jsii.String("{slug}"), nil)

	// Define methods for the /api/v1/questions/{slug} resource
	questionResource.AddMethod(jsii.String("GET"), awsapigateway.NewMockIntegration(&awsapigateway.MockIntegrationProps{
		IntegrationResponses: &[]*awsapigateway.IntegrationResponse{{
			StatusCode: jsii.String("200"),
		}},
		RequestTemplates: &map[string]*string{
			"application/json": jsii.String("{\"statusCode\": 200}"),
		},
		PassthroughBehavior: awsapigateway.PassthroughBehavior_NEVER,
	}), nil)

	// Output the API URL
	awscdk.NewCfnOutput(stack, jsii.String("ApiUrl"), &awscdk.CfnOutputProps{
		Value: api.Url(),
	})

	return stack

}
