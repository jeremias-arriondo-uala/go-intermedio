.PHONY: build, clean

LAMBDA_NAME := $(shell grep -A3 'artifactId:' lambda.yaml | tail -n1 | awk '{ print $$2}')

build:
	mkdir -p .build
	cd .build && \
	GOOS=linux GOARCH=amd64 go build -o main ../cmd/main.go && \
	zip main.zip main

update: build
	cd .build && \
	aws lambda update-function-code --function-name ${LAMBDA_NAME} --zip-file fileb://main.zip --region us-east-1 --profile $(profile)

clean:
	rm -rf ./bin