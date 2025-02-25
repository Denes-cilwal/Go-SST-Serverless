import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import {RestApi, LambdaIntegration} from "aws-cdk-lib/aws-apigateway";


export class GoServerlessSstStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);



    const myFunc = new cdk.aws_lambda.Function(this, 'MyFunction', {
      // name of folder in lambdas/ that contains the go code main.go
      code: cdk.aws_lambda.Code.fromAsset('lambdas'),
      runtime: cdk.aws_lambda.Runtime.GO_1_X,
      handler: 'main',

    });

    // cdk stack that will deploy our lambda function and it's going to deploy the
    // api gateway and bind the lambda function to the api gateway

    const api = new RestApi(this, 'MyApi', {
      defaultCorsPreflightOptions: {
        allowOrigins: ['*'],
        allowMethods: ['GET', 'POST', 'PUT', 'DELETE', 'OPTIONS'],
        allowHeaders: ['*'],
      }
    });

    const integration = new LambdaIntegration(myFunc);
    const testResource = api.root.addResource('test');
    testResource.addMethod('GET', integration);
  }
}
