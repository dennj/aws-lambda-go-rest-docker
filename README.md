# AWS Lambda Go REST Docker

A template for creating AWS Lambda REST APIs with Go, deployed using Docker. This setup supports local testing without requiring AWS API Gateway for routing.

## Features

- **AWS Lambda:** Run your Go applications in AWS Lambda.
- **Docker:** Deploy using Docker containers.
- **Local Testing:** Test your API locally without AWS API Gateway.
- **Routing:** Handle routing within the Go application.
- **MIT License:** Open-source and free to use.

### Author

# Dennj Osele - dennj.osele@gmail.com

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [AWS CLI](https://aws.amazon.com/cli/)
- [Go](https://golang.org/dl/)

### Clone the Repository

```bash
git clone https://github.com/dennj/aws-lambda-go-rest-docker.git
cd aws-lambda-go-rest-docker
```

### Build and Run Locally

## Build the Docker Image

```bash
docker build --platform linux/amd64 -t docker-image:test .
```

## Run the Docker Container

```bash
docker run -d -p 9000:9000 --name lambda-go docker-image:test
```

## Test the Endpoints

```bash
curl -X GET "http://localhost:9000/v1/getinfo/12345"
```

```bash
curl -X POST "http://localhost:9000/v1/set/12345/of/67890"
```

### Stop the Docker Container

## List Running Containers

```bash
docker ps
```

## Stop the Container

```bash
docker stop lambda-go
```

## Remove the Container

```bash
docker rm lambda-go
```

### Deploy to AWS Lambda

## Push Docker Image to Amazon ECR

```bash
aws ecr create-repository --repository-name lambda-go-rest
aws ecr get-login-password --region <region> | docker login --username AWS --password-stdin <account-id>.dkr.ecr.<region>.amazonaws.com
docker tag lambda-go-rest <account-id>.dkr.ecr.<region>.amazonaws.com/lambda-go-rest:latest
docker push <account-id>.dkr.ecr.<region>.amazonaws.com/lambda-go-rest:latest
```

## Create Lambda Function

```bash
aws lambda create-function --function-name lambda-go \
--package-type Image \
--code ImageUri=<account-id>.dkr.ecr.<region>.amazonaws.com/lambda-go-rest:latest \
--role <your-execution-role-arn>
```

## Set Up Amazon API Gateway HTTP API

```bash
aws apigatewayv2 create-api --name "LambdaGoAPI" --protocol-type HTTP
```

## Integrate Lambda

```bash
aws apigatewayv2 create-integration \
--api-id <api-id> \
--integration-type AWS_PROXY \
--integration-uri arn:aws:apigateway:<region>:lambda:path/2015-03-31/functions/arn:aws:lambda:<region>:<account-id>:function:lambda-go/invocations \
--payload-format-version 2.0
```

## Catch all routes

```bash
aws apigatewayv2 create-route \
--api-id <api-id> \
--route-key "ANY /{proxy+}" \
--target "integrations/<integration-id>"
```

## Deploy API

```bash
aws apigatewayv2 create-deployment --api-id <api-id> --stage-name prod
```

## Add Permissions for API Gateway to Invoke Lambda:

```bash
aws lambda add-permission --function-name lambda-go \
--statement-id apigateway-access \
--action lambda:InvokeFunction \
--principal apigateway.amazonaws.com \
--source-arn arn:aws:execute-api:<region>:<account-id>:<api-id>/*/*/*
```
