#!/bin/bash

CONTAINER=openapitools/openapi-generator-cli:v5.0.1
CONTAINER_TOOL=$([ -x /usr/bin/podman ] && echo podman || echo docker)

if [ -n "$1" ] ; then
    case "$1" in
        --docker) CONTAINER_TOOL=docker
            ;;
        --podman) CONTAINER_TOOL=podman
            ;;
        --help|*) echo "Usage: $0 [ --docker | --podman ]" >&2
                  exit 1
            ;;
    esac
fi

# For now, we skip the dateTime parsing because of incompatible formats ( inventory does not produce fully compliant

function filter_oneof() {
    # filter out oneOf types which openapi generator can't handle
    sed -i 's|"type": {"oneOf": \[{"type": "string", "example": "security"}, {"type": "array", "items": {"type": "string", "example": "security"}}\]},|"type": {"type": "array", "items": {"type": "string", "example": "security"}},|;
            s|"severity": {"oneOf": \[{"type": "string", "enum": \["Low", "Moderate", "Important", "Critical", null\], "nullable": true}, {"type": "array", "items": {"type": "string", "enum": \["Low", "Moderate", "Important", "Critical", null\], "nullable": true}}\]}},| "severity": {"type": "array", "items": {"type": "string", "enum": \["Low", "Moderate", "Important", "Critical", null\], "nullable": true}}},|;
            s|"items":{"oneOf":\[{"properties":{"count":{"type":"integer"},"value":{"type":"string"}},"type":"object"},{"properties":{"count":{"type":"integer"},"value":{"type":"boolean"}},"type":"object"}\]},|"items":{"properties":{"count":{"type":"integer"},"value":{"type":"string"}},"type":"object"},|g;
            s|"additionalProperties":{"oneOf":\[{"$ref":"#/components/schemas/system_profile.spec.yaml_NestedObject","x-scope":\[""\]},{"not":{"type":"object"}}\]},|"additionalProperties":{"not":{"type":"object"}},|g;
            s|Fqdn|fqdn|g;
            s|"CrossAccountRequestDetail":{"oneOf":\[{"\$ref":"#/components/schemas/CrossAccountRequestDetailByAccount"},{"\$ref":"#/components/schemas/CrossAccountRequestDetailByUseId"}\]}|"CrossAccountRequestDetail":{"\$ref":"#/components/schemas/CrossAccountRequestDetailByAccount"}|g;
            s|"items":{"oneOf":\[{"\$ref":"#/components/schemas/CrossAccountRequestByAccount"},{"\$ref":"#/components/schemas/CrossAccountRequestByUserId"}\]}|"items":{"\$ref":"#/components/schemas/CrossAccountRequestByAccount"}|g;
            # workaround for converting nullable date-time to nullable strings
            s|"format":"date-time","nullable":true,"type":"string"|"nullable":true,"type":"string"|g;
            ' $1
}

function generate_client() {
  NAME=$1
  SPEC_URL=$2

  LOCAL_JSON="$NAME/api/openapi.json"
  curl $SPEC_URL -o "$LOCAL_JSON"

  filter_oneof "$LOCAL_JSON"

  HERE=$(pwd)
  # remove old code and doc
  find "$HERE/$NAME" -name client_test.go -prune -o \( -name '*.go' -o -name '*.md' \) -print0 \
        | xargs -0 rm -f

  $CONTAINER_TOOL run --rm -v ${HERE}:/local:z --security-opt=label=disable $CONTAINER generate \
      -i "/local/$LOCAL_JSON" \
      -g go \
      --api-package $NAME \
      -p packageName=$NAME,isGoSubmodule=true \
      --git-host "github.com" --git-user-id RedHatInsights --git-repo-id patchman-clients \
      --type-mappings DateTime=string \
      -o /local/$NAME

    if [ "$NAME" == "inventory" ] ; then
        sed -i 's/uNKNOWNBASETYPE/createCheckIn/g; s/UNKNOWN_BASE_TYPE/CreateCheckIn/g;' $NAME/api_hosts.go
    fi
}


$CONTAINER_TOOL image exists $CONTAINER || $CONTAINER_TOOL pull $CONTAINER
generate_client inventory "https://ci.cloud.redhat.com/api/inventory/v1/openapi.json"
generate_client vmaas "https://webapp-vmaas-ci.5a9f.insights-dev.openshiftapps.com/api/vmaas/v3/openapi.json"
generate_client rbac "https://ci.cloud.redhat.com/api/rbac/v1/openapi.json"
