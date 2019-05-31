#!/usr/bin/env bash

########################
### USE WITH CAUTION ###
########################

name="siloqcrud"
version="v0.1.0"
tag="$name"
repository="tpalanisamy"
namespace="iapi"
tlnamespace="kube-system"
#registry="cloud.docker.com/repository/docker"
versionedImageName="$tag:$version"
#taggedImageName="$registry/$repository/$versionedImageName"
taggedImageName="$repository/$versionedImageName"
configDirectory="$(pwd)"
configFileName="config.json"

#set -e makes the script fail when there is an error or failed test
#set -e
#add package names to the test, as -run only runs the packages that you specify.
go test -run ./pkg/api ./pkg/config -v

echo Building ${repository}/${name}:build
docker build -t ${repository}/${name}:build . -f Dockerfile.build

docker container create --name extract ${repository}/${name}:build
docker container cp extract:/go/src/${name}/cmd/svr/${name} ./
docker container rm -f extract

echo Building ${taggedImageName}

docker build --no-cache -t ${taggedImageName} . -f Dockerfile.release
docker rmi ${repository}/${name}:build

docker push ${taggedImageName}

echo kubectl delete secret ${name} --namespace=${namespace} --v=10
kubectl delete secret ${name} --namespace=${namespace}
echo kubectl create secret generic ${name} --from-file=${configDirectory}/${configFileName} --namespace=${namespace} --v=10
kubectl create secret generic ${name} --from-file=${configDirectory}/${configFileName} --namespace=${namespace}

helm delete --purge ${name} --tiller-namespace=${tlnamespace}
helm install helm --values ./helm/values.yaml --name ${name} --tiller-namespace=${tlnamespace} --namespace=${namespace}

