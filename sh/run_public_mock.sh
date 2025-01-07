#!/bin/bash

shell2http --port 4466 /tmp "curl -Ls http://localhost/public | dasel -r json -w toml"
