#!/bin/bash
scriptdir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
basedir="${scriptdir%/*}"

sudo mokapi --providers-file-directory "${basedir}/testing"
