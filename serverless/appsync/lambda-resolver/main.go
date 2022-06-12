package main

import (
	"myAppSyncApp-API/serverless/appsync/lambda-resolver/graphqlrouter"
	"os"

	dynamo "myAppSyncApp-API/serverless/storage"

	"github.com/andreaperizzato/go-config"
	"github.com/aws/aws-lambda-go/lambda"
)

type configuration struct {
	TableName string `env:"DATA_TABLE_NAME"`
}

func main() {

	stage := os.Getenv("STAGE")

	var envVars configuration
	ssm := config.NewSSMSourceWithSubstitutions(map[string]string{
		"stage": stage,
	})
	err := config.NewLoader(ssm, config.NewEnvSource()).Load(&envVars)
	if err != nil {

		os.Exit(-1)
		return
	}
	dynamoStore := dynamo.NewStore(envVars.TableName)

	routes := graphqlrouter.NewGraphQLRouter(dynamoStore)
	lambda.Start(routes.Route)

}
