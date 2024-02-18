# Introduction

## copy

1. copy repo `cp -rf koffloader/*  YourRepoName  && cp koffloader/.gitignore YourRepoName  && cp koffloader/.github  YourRepoName `

   replace all 'koffloader' to 'YourRepoName'

   replace all 'koffloader-io' and 'koffloader.io' to 'YourOrigin'

   replace all 'Copyright 2024' to be the right time

2. grep "====modify====" * -RHn --colour  and modify all of them

3. in a linux machine, update api/v1/openapi.yaml and `make update_openapi_sdk`

4. redefine CRD in pkg/k8s/v1
    rename directory name 'pkg/k8s/apis/koffloader.koffloader.io' 
    replace all 'kcluster' to 'YourCRDName'
    and `make update_crd_sdk`, and write code in pkg/kclusterManager

    rename pkg/kclusterManager and replace all 'kcluster' with your CRD name in this directory

    rm charts/crds/koffloader.koffloader.io_kclusters.yaml 

    # in repo: replace all "github.com/koffloader-io/spiderdoctor/pkg/kclusterManager" to "github.com/koffloader-io/spiderdoctor/pkg/${crdName}Manager"
    # in repo: find and replace all "kcluster" to YourCrd

5. update charts/ , and images/ , and CODEOWNERS

6. `go mod tidy` , `go mod vendor` , `go vet ./...` , double check all is ok

7. `go get -u` , `go mod tidy` , `go mod vendor` , `go vet ./...`  , update all vendor

8. create an empty branch 'github_pages' and mkdir 'docs'

9. enable third app

   personal github -> settings -> applications -> configure

   codefactor: https://github.com/marketplace/codefactor and https://www.codefactor.io/dashboard

   sonarCloud: https://sonarcloud.io/projects/create

   codecov: https://github.com/marketplace/codecov  and https://app.codecov.io/gh

10. github seetings:
      koffloader.io/REPO  -> settings -> secrets and variable -> actions -> add secret 'WELAN_PAT' , 'ACTIONS_RUNNER_DEBUG'=true , 'ACTIONS_STEP_DEBUG'=true, 'CODECOV_TOKEN'

      koffloader.io  -> settings -> secrets -> actions -> grant secret to repo

      koffloader.io/REPO  -> settings -> general -> feature -> issue

      koffloader.io/ORG  -> settings -> actions -> general -> allow github action to create pr
      koffloader.io/REPO  -> settings -> actions -> general -> allow github action to create pr

      koffloader.io  -> settings -> packages -> public 

      repo -> packages -> package settings -> Change package visibility

      create 'github_pages' branch, and repo -> settings -> pages -> add branch 'github_pages', directory 'docs'

      repo -> settings -> branch -> add protection rules for 'main' and 'github_pages' and 'release*'

      repo -> settings -> tag -> add protection rules for tags

11. add badge to readme:

    github/workflows/call-e2e.yaml

    github/workflows/badge.yaml

    auto nightly ci

    release version

    code coverage from https://app.codecov.io/gh

    go report from https://goreportcard.com

    codefactor: https://www.codefactor.io/dashboard

    sonarCloud: https://sonarcloud.io/projects

12. build base image , 
    update BASE_IMAGE in images/agent/Dockerfile and images/controller/Dockerfile
    run test

## local develop

1. `make build_local_image`

2. `make e2e_init`

3. `make e2e_run`

4. check proscope, browser vists http://NodeIP:4040

5. apply cr

        cat <<EOF > kcluster.yaml
        apiVersion: koffloader.koffloader.io/v1
        kind: Kcluster
        metadata:
          name: test
        spec:
          ipVersion: 4
          subnet: "1.0.0.0/8"
        EOF
        kubectl apply -f kcluster.yaml

## chart develop

helm repo add rock https://koffloader-io.github.io/koffloader/

## upgrade project 

1. golang version: edit golang version in Makefile.defs and `make update_go_version`

2. 更新所有包  go get -u ./...
