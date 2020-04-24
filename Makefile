run-dynamodb:
	docker kill dynamodb-lookup || true
	docker run --rm -d --name dynamodb-lookup -p 8000:8000 amazon/dynamodb-local

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/connecntionmanager ./connecntionmanager

deploy-local:
	AWS_PROFILE=personal sls deploy -s dev