#!/bin/bash

# Generate the new community manifests, given the current wmco manifests.

# Must be ran within community-operators-prod root.

# Usage: community.sh <new_community_wmco_version>
# Example:
# Update community manifests v5.0.1
# Run: bash ../windows-machine-config-operator/hack/community/community.sh 5.0.1

# create and copy manifests to new folder and rename
mkdir operators/community-windows-machine-config-operator/$1
cp -R ../windows-machine-config-operator/bundle/manifests ../windows-machine-config-operator/bundle/metadata operators/community-windows-machine-config-operator/$1
mv operators/community-windows-machine-config-operator/$1/manifests/windows-machine-config-operator.clusterserviceversion.yaml operators/community-windows-machine-config-operator/$1/manifests/community-windows-machine-config-operator.$1.clusterserviceversion.yaml
