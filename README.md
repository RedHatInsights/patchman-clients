# patchman-clients
This repo stores automatically generated Go clients for
[Red Hat Insights](https://www.redhat.com/en/technologies/management/insights) projects:
- Inventory (https://github.com/RedHatInsights/insights-host-inventory)
- Rbac (https://github.com/RedHatInsights/insights-rbac)
- Vmaas (https://github.com/RedHatInsights/vmaas)

These clients are primarily used in
[patchman-engine](https://github.com/RedHatInsights/patchman-engine) project.

## Generate Go clients
This repo also contains script to generate current version of clients using
containerized
[openapi-generator-cli](https://github.com/openapitools/openapi-generator-cli) tool.

To generate current version of clients run:
~~~bash
./generate-clients.sh # Use --docker or --podman to select specific container engine.
~~~
Note: be careful and don't delete `{inventory,rbac,vmaas}/client_test.go` files
needed for build testing.

## Test compilation
You can test locally whether newly generated sources are valid and possible to build.
~~~bash
docker-compose -f docker-compose.test.yml up --build test
~~~
