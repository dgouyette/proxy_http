ls
# Context

This tool allows you to check automation status

It can be executed on your personnal computer, but also on Google Cloud Platform.

Network flows have been opened between Solocal Kubernetes Cluster and marketing cloud.


# Build

Golang need to be installed https://golang.org/doc/install

If you want to build it for Windows, you'll have to add an environnement variable :

```bash
export GOOS=windows
```

If you want to build it with a specified name : 

```bash
go build -o automation  .
```

# Deployment

It will deployed as a docker image using Gitlab and Kaniko

```yaml
build:
  stage: build
  script:
    - mkdir -p /kaniko/.docker/
    - echo "{\"auths\":{\"$CI_REGISTRY\":{\"username\":\"$CI_REGISTRY_USER\",\"password\":\"$CI_REGISTRY_PASSWORD\"}}}" > /kaniko/.docker/config.json
    - /kaniko/executor --reproducible --single-snapshot --cache=true --context $CI_PROJECT_DIR --destination $CI_REGISTRY_IMAGE:${BRANCH_TAG} --build-arg KEY_FILE_PATH=$CI_PIPELINE_ID
```      

Dockerfile
```Dockerfile
FROM golang:1.13.3

WORKDIR /build

COPY src/* /build/
COPY go.mod . 
RUN go mod download

RUN go build -o /build/automation  .
ENTRYPOINT ["/build/automation"]
```
