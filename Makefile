SHELL=/bin/bash -o pipefail

export PATH := .bin:${PATH}

.bin/cowk8s: Makefile