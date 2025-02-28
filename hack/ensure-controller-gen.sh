#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

BIN_ROOT="./bin"

check_controller_gen_installed() {
	if ! [ -x "$(command -v "${BIN_ROOT}/controller-gen")" ]; then
		echo 'controller-gen is not found installing'
		echo 'Installing controller-gen' && install_controller_gen
	fi
}

verify_controller_gen_version() {

	local controller_gen_version
	controller_gen_version="$(${BIN_ROOT}/controller-gen --version | grep -Eo "([0-9]{1,}\.)+[0-9]{1,}" | head -1)"
	if [[ "${MINIMUM_CONTROLLER_GEN_VERSION}" != "${controller_gen_version}" ]]; then
		cat <<EOF
Detected controller-gen version: ${controller_gen_version}
Install ${MINIMUM_CONTROLLER_GEN_VERSION} of controller-gen
EOF

		echo 'Installing controller-gen' && install_controller_gen
	else
		cat <<EOF
Detected controller-gen version: ${controller_gen_version}.
controller-gen ${MINIMUM_CONTROLLER_GEN_VERSION} is already installed.
EOF
	fi
}

install_controller_gen() {
	if [[ "${OSTYPE}" == "linux"* ]]; then
		if ! [ -d "${BIN_ROOT}" ]; then
			mkdir -p "${BIN_ROOT}"
		fi
		go install -v sigs.k8s.io/controller-tools/cmd/controller-gen@master
		cp "$HOME/go/bin/controller-gen" "${BIN_ROOT}/controller-gen"
	else
		set +x
		echo "The installer does not work for your platform: $OSTYPE"
		exit 1
	fi
}

check_controller_gen_installed "$@"
verify_controller_gen_version "$@"
