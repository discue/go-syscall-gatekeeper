#!/bin/bash

set -uo pipefail

declare -r main_path="$1"

go run $main_path run --allow-file-system --no-implicit-allow sed -n '10,20p' run.sh

if [[ $? -ne 0 ]]; then
    exit 0
fi

exit 1
