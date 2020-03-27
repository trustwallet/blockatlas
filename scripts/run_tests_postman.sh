#!/usr/bin/env bash
# blockatlas host
HOST=$1

cd ..
make newman test=transaction host=$HOST
make newman test=token host=$1
make newman test=staking host=$1
make newman test=collection host=$1
make newman test=domain host=$1
make newman test=observer_test host=$1
