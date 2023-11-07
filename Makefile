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

########################################################################
# k3d cluster

k3d-start:
	@./scripts/k3d-create.sh
	k3d kubeconfig merge -ad
	kubectl config use-context k3d-critik8s
	@./scripts/k3d-prepare.sh

k3d-delete:
	k3d cluster delete critik8s
	@if test -f /usr/local/bin/rke2-uninstall.sh; then sudo sh /usr/local/bin/rke2-uninstall.sh; fi

########################################################################
# API Generator

gen-model-apis:
	@./scripts/gen-model-api.sh

########################################################################
# Monitor UI

monitor-ui-install:
	cd src/monitor && npm install

monitor-ui-start:
	cd src/monitor && npm run dev
