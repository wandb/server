GREEN='\033[0;32m'
BLUE='\033[0;94m'
LIGHT_BLUE='\033[0;34m'
YELLOW='\033[0;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

function log_success() {
    printf "${GREEN}✔ $1${NC}\n" 1>&2
}

function log() {
    printf "%s\n" "$1" 1>&2
}

function log_step() {
    printf "\n${BLUE}⚙  $1${NC}\n" 1>&2
}

function log_substep() {
    printf "\t${LIGHT_BLUE}- $1${NC}\n" 1>&2
}

function log_fail() {
    printf "${RED}$1${NC}\n" 1>&2
}

function log_warn() {
    printf "${YELLOW}$1${NC}\n" 1>&2
}

function bail() {
    log_fail "$@"
    exit 1
}