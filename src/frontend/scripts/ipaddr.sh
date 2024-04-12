#!/bin/bash

IPADDR=$(ipconfig getifaddr en0)

sed -i '' "s/VITE_IPADDR=.*/VITE_IPADDR=$IPADDR/" .env
