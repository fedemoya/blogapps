#!/bin/bash

# Exit When Any Command Fails
set -e

echo "Initiating message processor test"
echo ""

echo "Starting rabbitmq"
echo ""
docker-compose up -d

rabbitmq_is_ready() {
  docker-compose exec rabbitmq rabbitmqadmin list queues > /dev/null
}

until rabbitmq_is_ready; do
  sleep 1
  echo "Waiting rabbitmq to start..."
done

echo "Cleanup old binaries and output files"
echo ""
rm -f message-processor message-processor-output

echo "Building & Running message-processor"
echo ""
go build
./message-processor &> message-processor-output.txt &
MESSAGE_PROCESSOR_PID=$!

echo "Sending messages to rabbitmq"
echo ""
for i in {1..10} ; do
    echo "Publishing withdrawal.created event with amount #$i"
    docker-compose exec rabbitmq rabbitmqadmin publish routing_key="message-consumer-queue" payload="
      {
    		\"type\": \"withdrawal.created\",
    		\"data\": {
    			\"withdrawal_id\": \"e728f3a7-b92f-46fe-b080-524442065cb3\",
    			\"amount\": $i,
    			\"source_account\": \"source account details\",
    			\"destination_account\": \"destination account details\"
    		}
    	}
    "
done

echo "Terminating message processor"
echo ""
kill -n 15 $MESSAGE_PROCESSOR_PID

echo "Shutting down rabbitmq"
echo ""
docker-compose down

echo "Printing message processor output"
echo ""
cat message-processor-output.txt
