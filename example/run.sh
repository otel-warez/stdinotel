#!/bin/bash
set -ex pipefail

docker compose up -d

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
GOOS=`go env GOOS`
GOARCH=`go env GOARCH`

export STDINOTEL_PROTOCOL=otlp
export STDINOTEL_TLS_INSECURE_SKIP_VERIFY=true

echo "Send data by piping from a file"
echo "foo" | $SCRIPT_DIR/../bin/stdinotel_${GOOS}_${GOARCH}
cat $SCRIPT_DIR/lorem.txt | $SCRIPT_DIR/../bin/stdinotel_${GOOS}_${GOARCH}

echo "Send data interactively. Finish by pressing enter."
$SCRIPT_DIR/../bin/stdinotel_${GOOS}_${GOARCH}
