# Tobi's blog Service
Blog servie written on Golang and deployment on AWS lambda.

## Getting Started
### Installing
Build local package
```make build-local```

Run the binary
```./bin/main```

## Deployment
Please read https://docs.aws.amazon.com/lambda/latest/dg/lambda-go-how-to-create-deployment-package.html for 
the deployment details.

### Build package for AWS lambda
Enable algnhsa AWS Lambda Go net/http server adapter (https://github.com/akrylysov/algnhsa)

Build package
```make build```

### Create a AWS lambda function (for the first time)
```
aws lambda create-function --function-name blog-post-function --runtime go1.x \
  --zip-file fileb://main.zip --handler main \
  --role arn:aws:iam::705179926964:role/lambda-role
 ```
 
### Deploy
```make deploy```

### Create a API Gateway to trigger the function
Please read https://docs.aws.amazon.com/apigateway/latest/developerguide/getting-started-with-lambda-integration.html

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* **Nguyen Hoang Anh** - *Initial work* - [AnhNguyen](https://github.com/anhnguyenbk)
See also the list of [contributors](https://github.com/anhnguyenbk/blog-service/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
