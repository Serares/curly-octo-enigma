package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/coreos/go-oidc"
)

var (
	clientID     = os.Getenv("GOOGLE_CLIENT_ID")
	oidcProvider *oidc.Provider
)

func init() {
	var err error
	ctx := context.Background()
	oidcProvider, err = oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		panic(err)
	}
}

func authorizerHandler(ctx context.Context, request events.APIGatewayCustomAuthorizerRequestTypeRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	token := request.Headers["Authorization"]
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).WithGroup("CRUD AUTHORIZER")
	log.Info("The event object", "event", request)

	idToken, err := oidcProvider.Verifier(&oidc.Config{ClientID: clientID}).Verify(ctx, token)
	if err != nil {
		return generatePolicy("user", "Deny", request.MethodArn, ""), fmt.Errorf("Unauthorized: %v", err)
	}

	var claims struct {
		Email string `json:"email"`
		Sub   string `json:"sub"`
		Name  string `json:"name"`
	}

	if err := idToken.Claims(&claims); err != nil {
		return generatePolicy("user", "Deny", request.MethodArn, ""), fmt.Errorf("Unauthorized: %v", err)
	}

	// Include the Authorization token in the context
	return generatePolicy(claims.Sub, "Allow", request.MethodArn, token), nil
}

func generatePolicy(
	principalID,
	effect,
	resource,
	token string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalID}
	authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
		Version: "2012-10-17",
		Statement: []events.IAMPolicyStatement{
			{
				Action:   []string{"execute-api:Invoke"},
				Effect:   effect,
				Resource: []string{resource},
			},
		},
	}
	authResponse.Context = map[string]interface{}{
		"Authorization": token,
	}
	return authResponse
}

func main() {
	lambda.Start(authorizerHandler)
}
