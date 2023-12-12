#!/usr/bin/env bash
#
# Copyright The OpenTelemetry Authors
# SPDX-License-Identifier: Apache-2.0
#
# This script checks the GitHub CODEOWNERS file for any code owners
# of contrib components and returns a string of the code owners if it
# finds them.

set -euo pipefail

get_component_type() {
  echo "${COMPONENT}" | cut -f 1 -d '/'
}

get_codeowners() {
  echo "$((grep -m 1 "^${1}" .github/CODEOWNERS || true) | sed 's/   */ /g' | cut -f3- -d ' ')"
}

if [[ -z "${COMPONENT:-}" ]]; then
    echo "COMPONENT has not been set, please ensure it is set."
    exit 1
fi

COMPONENT_TYPE=$(get_component_type "${COMPONENT}")
CUR_DIRECTORY=$(dirname "$0")
VALID_COMPONENT=$(bash "${CUR_DIRECTORY}/get-components.sh" | grep -x "${COMPONENT}" || true)
VALID_COMPONENT_WITH_TYPE=$(bash "${CUR_DIRECTORY}/get-components.sh" | grep -x "$COMPONENT$COMPONENT_TYPE" || true)

if [[ -z "${VALID_COMPONENT:-}" ]] && [[ -z "${VALID_COMPONENT_WITH_TYPE:-}" ]]; then
    echo ""
    exit 0
fi

# grep exits with status code 1 if there are no matches,
# so we manually set RESULT to 0 if nothing is found.
RESULT=$(grep -c "${COMPONENT}" .github/CODEOWNERS || true)

# there may be more than 1 component matching a label
# if so, try to narrow things down by appending the component
# or a forward slash to the label.
if [[ ${RESULT} != 1 ]]; then
    OWNERS="$(get_codeowners "${COMPONENT}${COMPONENT_TYPE}")"

    if [[ -z "${OWNERS:-}" ]]; then
        OWNERS="$(get_codeowners "${COMPONENT}/")"
    fi
else
    OWNERS="$(get_codeowners $COMPONENT)"
fi

echo "${OWNERS}"
