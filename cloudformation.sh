#!/usr/bin/env bash

aws cloudformation describe-stacks --stack-name $STACK_NAME | jq '.[0]' | read stacks
echo $stacks
