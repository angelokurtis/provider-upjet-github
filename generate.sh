#!/usr/bin/env bash

set -e

export PROVIDER_NAME_LOWER=github
export PROVIDER_NAME_NORMAL=GitHub
export ORGANIZATION_NAME=angelokurtis
export CRD_ROOT_GROUP=kurtis.dev.br
hack/prepare.sh

export TERRAFORM_PROVIDER_SOURCE=integrations/github
export TERRAFORM_PROVIDER_REPO=https://github.com/integrations/terraform-provider-github
export TERRAFORM_PROVIDER_VERSION=6.5.0
export TERRAFORM_PROVIDER_DOWNLOAD_NAME=terraform-provider-github
export TERRAFORM_NATIVE_PROVIDER_BINARY=terraform-provider-github_v5.32.0
export TERRAFORM_DOCS_PATH=website/docs/r
go install golang.org/x/tools/cmd/goimports@latest
make generate
