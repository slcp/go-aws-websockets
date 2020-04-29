run-dynamodb:
	docker kill dynamodb-lookup || true
	docker run --rm -d --name dynamodb-lookup -p 8000:8000 amazon/dynamodb-local

build:
	sh ./build.sh

deploy-local:
	AWS_PROFILE=personal sls deploy -s dev