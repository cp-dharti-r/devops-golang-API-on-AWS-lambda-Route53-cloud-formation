# Golang API Deployment on AWS Lambda with Route 53

This repository contains the infrastructure and deployment pipeline for a Golang API. The API is deployed to AWS Lambda and made accessible through a custom domain using API Gateway and Route 53. The project uses CloudFormation for infrastructure management and GitHub Actions for CI/CD.

---

## Features

- **Golang API**: Uses the Gin framework for the API implementation.
- **AWS Lambda**: Serverless deployment for scalability and cost efficiency.
- **API Gateway**: Acts as a gateway for the Lambda function.
- **Route 53**: Custom domain configuration for the API endpoint.
- **CloudFormation**: Infrastructure as code for AWS resources.
- **GitHub Actions**: CI/CD pipeline for automated deployment.

---

## Architecture

1. **Lambda Function**: Hosts the Golang API logic.
2. **API Gateway**: Maps API routes to the Lambda function.
3. **Custom Domain**: Configured in API Gateway and linked to Route 53.
4. **Route 53**: Provides DNS resolution for the custom domain.
5. **GitHub Actions**: Automates deployment using CloudFormation.

---

## Prerequisites

### Tools and Services

- **Golang**: Install [Go](https://go.dev/dl/).
- **AWS CLI**: Configure AWS CLI with your credentials.
- **GitHub Secrets**: Store sensitive information (e.g., AWS keys, domain details).
- **S3 Bucket**: Pre-created bucket for Lambda deployment packages.

### AWS Services

- API Gateway
- Lambda
- Route 53
- CloudFormation
- IAM

---

### Set Up GitHub Secrets

Add the following secrets to your repository:

- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY
- S3_BUCKET_NAME
- CUSTOM_DOMAIN (e.g., api.yourdomain.com)
- HOSTED_ZONE_ID
- CERTIFICATE_ARN

### Build and Deploy the API

Push your changes to the main branch to trigger the GitHub Actions pipeline:

```bash
git add .
git commit -m "Deploy API"
git push origin main
```

### Test the API

Once deployed, access the API at:

```bash
https://api.yourdomain.com/hello
```

#### Response

```bash
{
  "message": "Hello from AWS Lambda!"
}
```

### Development and Testing

#### Run Locally

Install dependencies:

```bash
go mod tidy
```

Run the API locally:

```bash
go run main.go
```

Access the API at http://localhost:8080/hello.

#### Build for Deployment

Prepare the binary for Lambda:

```bash
GOOS=linux GOARCH=amd64 go build -o main
zip deployment.zip main
```
