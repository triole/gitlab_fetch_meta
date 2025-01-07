#!/bin/bash

r "http://localhost/repos" \
  -a "punch_public.toml,public.json" \
  --log-level debug
