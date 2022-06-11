# myAppSyncApp
AppSync GraphQL API - In Go

How to build and run this application. 

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




