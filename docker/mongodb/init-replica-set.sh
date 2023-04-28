#!/bin/bash
set -e

# Start mongod in the background
mongod --replSet rs0 --keyFile /etc/mongo/keyfile --bind_ip_all &

echo "Waiting for MongoDB to start..."
until mongosh --host mongo1 --eval "db.adminCommand('ping')" >/dev/null 2>&1; do
  sleep 1
done

# Check if the replica set is already initialized
if ! mongosh --host mongo1 -u rust_drop -p "<H;wFO&:L:ym;9" --eval "rs.status()" | grep -q "NotYetInitialized"; then
  echo "Replica set is already initialized. Skipping initiation."
else
  echo "Initializing replica set..."
  mongosh --host mongo1 -u rust_drop -p "<H;wFO&:L:ym;9" --eval "rs.initiate();"
  sleep 5

  # Ensure mongo1 is the primary before adding other members
  until mongosh --host mongo1 -u rust_drop -p "<H;wFO&:L:ym;9" --eval "rs.status().myState" | grep -q 1; do
    echo "Waiting for mongo1 to become primary..."
    sleep 1
  done

  echo "Adding replica set members..."
  mongosh --host mongo1 -u rust_drop -p "<H;wFO&:L:ym;9" --eval "rs.add('mongo2');"
  mongosh --host mongo1 -u rust_drop -p "<H;wFO&:L:ym;9" --eval "rs.add('mongo3');"

  echo "Waiting for secondary members to catch up..."
  until mongosh --host mongo1 -u rust_drop -p "<H;wFO&:L:ym;9" --eval "rs.status().members.every(member => member.stateStr === 'PRIMARY' || member.stateStr === 'SECONDARY')" | grep -q true; do
    sleep 1
  done
fi

mongosh --host mongo1 -u rust_drop -p "<H;wFO&:L:ym;9" --eval "rs.status();"
echo "Replica set initialized successfully."

# Keep the script running as long as the mongod process is running
wait $!
