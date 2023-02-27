[AWS Lambda function handler in Go](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)

[context pack](https://pkg.go.dev/context)

[AWS Lambda context object in Go](https://docs.aws.amazon.com/lambda/latest/dg/golang-context.html)

When Lambda runs your function, it passes a context object to the handler. 
This object provides methods and properties with information about the invocation, function, and execution environment.

`func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
    return fmt.Sprintf("Hello %s!", name.Name), nil
}`

Build `GOOS=linux GOARCH=amd64 go build -o main`

Zip `zip main.zip main`  

