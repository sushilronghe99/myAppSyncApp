export VERSION=`git rev-list --count HEAD`
echo Building Serverless version $VERSION


echo Building Serverless version 


env GOOS=linux go build  -o bin/resolver appsync/lambda-resolver/main.go