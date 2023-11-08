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

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
source "${SCRIPT_DIR}/helpers.sh"

# Ensure we have a value for --system-domain
prepare_system_domain

echo "Preparing k3d environment ..."

#Install rabbitmq operator
set +e
helm install rabbitmq-operator charts/rabbitmq-operator -n rabbitmq --create-namespace --wait
set -e

#Install rabbitmq cluster
set +e
helm install rabbitmq-cluster charts/rabbitmq-cluster --values settings.yaml -n rabbitmq --create-namespace --wait
set -e

echo
echo "Done preparing k3d environment! ✔️"
