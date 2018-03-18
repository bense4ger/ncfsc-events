#!/usr/bin/env bash

aws cloudformation describe-stacks --region $EC2_REGION --stack-name $STACK_NAME | jq '.[0]' | read stacks
jq '.[0] | .StackId' | read stackId

echo $stackId
