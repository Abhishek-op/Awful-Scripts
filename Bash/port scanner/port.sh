#!/bin/bash
host=$1
port=$2

function scan
{
  (echo >/dev/tcp/$host/$port) >/dev/null 2>&1 && echo "$port open"
}
scan