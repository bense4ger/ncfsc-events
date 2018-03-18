#!/usr/bin/env bash
DATE=`date '+%Y%m%d%H%M%S'`

stackId=$(aws cloudformation describe-stacks --region $EC2_REGION --stack-name $STACK_NAME | jq '.Stacks[0].StackId')
stackId=$(sed -e 's/^"//' -e 's/"$//' <<<"$stackId")

if [ "$stackId" = "null" ]
then
    stackId=$(aws cloudformation create-stack --stack-name $STACK_NAME --template-body file:///template.json | jq '.StackId')
else
    echo "Stack already exists $stackId"
fi

csId=$(aws cloudformation create-change-set --stack-name $stackId --change-set-name ncfsc-events-$DATE --use-previous-template | jq '.Id')
csId=$(sed -e 's/^"//' -e 's/"$//' <<<"$csId")

csStatus=$(aws cloudformation describe-change-set --change-set-name $csId | jq '.Status')
csStatus=$(sed -e 's/^"//' -e 's/"$//' <<<"$csStatus")

if [ "$csStatus" != "FAILED" ]
then
    aws cloudformation execute-change-set --change-set-name $csId
else
    echo "Change Set Not Executed"
fi

echo "Updating function code"
update=$(aws lambda update-function-code --function-name $FUNCTION_NAME --s3-bucket code-drops --s3-key $STACK_NAME.zip | jq '.FunctionName')


