set -uo pipefail

DEBUG=true

ROOT="$(cd "$(dirname "$0")/.." &>/dev/null; pwd -P)"
[ "$DEBUG" = true ] && echo root: $ROOT

PID=`pgrep main`
[ ! -z "$PID" ] && \
	{ echo PID: $PID. Killing and starting new instance... ; kill $PID ; } \
	|| echo No instace was up. Starting...

go run main.go &

