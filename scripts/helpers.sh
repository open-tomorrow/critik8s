#!/bin/bash
# Copyright © 2023 OpenTomorrow
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#     http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

function prepare_system_domain {
  if [[ -z "${CRITIK8S_SYSTEM_DOMAIN}" ]]; then
    echo -e "\e[32mCRITIK8S_SYSTEM_DOMAIN not set. Trying to use a magic domain...\e[0m"
    CRITIK8S_CLUSTER_IP=$(docker inspect k3d-critik8s-server-0 | jq -r '.[0]["NetworkSettings"]["Networks"]["critik8s"]["IPAddress"]')
    if [[ -z $CRITIK8S_CLUSTER_IP ]]; then
      echo "Couldn't find the cluster's IP address"
      exit 1
    fi
    export CRITIK8S_CLUSTER_IP="${CRITIK8S_CLUSTER_IP}"
    export CRITIK8S_SYSTEM_DOMAIN="${CRITIK8S_CLUSTER_IP}.omg.howdoi.website"
  fi
  echo -e "Using \e[32m${CRITIK8S_SYSTEM_DOMAIN}\e[0m for critik8s domain"
}
