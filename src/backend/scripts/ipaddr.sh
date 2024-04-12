#!/bin/bash

IPADDR=$(ipconfig getifaddr en0)

yq eval ".network.en0.ip_address = \"$IPADDR\"" -i internal/config/values.yaml
