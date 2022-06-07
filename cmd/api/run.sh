#!/usr/bin/env bash

export JAEGER_DISABLED=false
export JAEGER_SAMPLER_TYPE="const"
export JAEGER_SAMPLER_PARAM=1
export JAEGER_REPORTER_LOG_SPANS=true
export JAEGER_AGENT_HOST="120.46.179.205"
export JAEGER_AGENT_PORT=6831
go run ./main.go
