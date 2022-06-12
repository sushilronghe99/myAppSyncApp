## myAppSyncApp

This is the sample project to explain how the data mapping resolver and graphql lambda resolver works. The app is developed in GO. The details steps are explained below. 

- Set up the serverless golang project
-	Set up the DynamoDB & define the app data model 
-	Set up Mapping template for data resolver 
-	Expose the graphQL resolver lambda 
-	Deploy the application to the AWS 


#How to build and run this application. 

Step 1 : Set up Serverless 
Install the serverless framework using npm and configure it to use your aws account. (Get the key and secret by logging to aws console.)

Npm install -g serverless 
serverless config credentials --provider aws --key <access key ID> --secret <secret access key>

Following  serverless plugins are required for appsync (more details on their github page). No need to install this plugin just run npm install. 

npm install serverless-appsync-plugin --save-dev

Go to serverless folder and modify the serverless.xml file to match your aws account ID and region. 
  
Run following commands. 

Go mod tidy (this should install all required modules in your local machine)
./build.sh  (this should build the go app) 
serverless deploy (this will deploy the app to your aws account. ) 
 
## More information About the App setup

Step 2 : Dynamo DB Setup 
Open the serverless.xml the database is defined under resources section  
```MyDB:
        Type: AWS::DynamoDB::Table
        Properties:
          TableName: MyAppSync-${opt:stage, self:provider.stage}
          AttributeDefinitions:
            - AttributeName: pk
              AttributeType: S
            - AttributeName: sk
              AttributeType: S
          
          KeySchema:
            - AttributeName: pk
              KeyType: HASH
            - AttributeName: sk
              KeyType: RANGE
          # Set the capacity to auto-scale
          BillingMode: PAY_PER_REQUEST 
  ```

 We have used a very basic setup of the database. Just a PK/SK will be used and no GSI is defined.
 
  ###GraphQL Schema 
GraphQL schema is as below. It represents the simple Todo application. 
  
``` 
type Todo {
  id: ID!
  text: String!
  done: Boolean!
  user: User!
}

type User {
  id: ID!
 
}

type Query {
  todos: [Todo!]!
  getTodoByID(id: String): Todo!
}

input NewTodo {
  text: String!
  userId: String
}

input UpdateTodoInput {
  id : ID,
  done: Boolean
}

type Mutation {
  createTodo(input: NewTodo!): Todo!
  updateTodo(input: UpdateTodoInput): Todo!
}

```
#DataBase resolver 
  
The Query GetTodo use the DynamoDB resolver. The request is as follows (request.vtl) 
```
  {
   "version" : "2017-02-28",
    "operation": "GetItem",
    "key": {
      "pk": {"S": "user_$ctx.identity.username" },
      "sk": { "S": "$ctx.args.id" }
    }
}
```
The Mutation Update Todo also use the DynamoDB resolver. The request is as follows. (request.vtl)
```
{
    "version" : "2018-05-29",
    "operation" : "UpdateItem",
    "key": {
        "pk" : {"S" : "user_$ctx.identity.username"},
        "sk" : {"S" : "$ctx.args.input.id" }, 
    },
    "condition": {
        "expression": "attribute_exists(pk)"
    },
    "update" : {
        "expression" : "SET #done = :done",
        "expressionNames" : {
           "#done" : "done"
       },
       "expressionValues" : {
           ":done" : { "BOOL": true}
       },
    },  
}
 ```

Create Todo is the lambda resolver and the request is processed by the lambda function. The lambda function is defined in appsync/lambda-resolver
```
{
    "version" : "2017-02-28",
    "operation": "Invoke",
    "payload": {
        "field": "createtodo",
        "identity": $util.toJson($context.identity.username),
        "arguments":{
            "input":$util.toJson($context.arguments.input),
            
        }
    } 
}
```
  
The mapping template configuration is defined in the serverless configuration under appsync header. It is as follows.
  
```
  mappingTemplates:
      # Queries
      - dataSource: My_dynamodb
        type: Query
        field: getTodoByID
        request: gettodo/request.vtl
        response: results.vtl
      - dataSource: My_dynamodb
        type: Mutation
        field: updateTodo
        request: updatetodo/request.vtl
        response: results.vtl
      - dataSource: lambda_resolver
        type: Mutation
        field: createTodo
        request: createtodo/request.vtl
        response: results.vtl
  ```
  
