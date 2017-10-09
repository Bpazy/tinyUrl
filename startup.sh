#!/bin/bash
cd $(dirname "$0")
nohup ./tinyUrl -port :11000 > /dev/null 2>&1 &
