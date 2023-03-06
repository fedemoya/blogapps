#!/bin/bash

# Exit When Any Command Fails
set -e

# Starting rabbitmq
docker-compose up -d

rabbitmq_is_ready() {
  docker-compose exec rabbitmq rabbitmqadmin list queues > /dev/null
}

until rabbitmq_is_ready; do
  sleep 1
  echo "Waiting rabbitmq to start..."
done

# Cleanup old binaries and output files
rm -f bin/message-processor output/message-processor.txt

# Build & Running message-processor
mkdir -p bin output
go build -o bin/ cmd/message-processor.go
bin/message-processor &> output/message-processor.txt &
MESSAGE_PROCESSOR_PID=$!

# Send messages to rabbitmq
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

# Terminate message-processor
kill -n 15 $MESSAGE_PROCESSOR_PID

# Shut down rabbitmq"
docker-compose down

echo ""
echo "-- Message processor output --"
echo ""
cat message-processor.txt
