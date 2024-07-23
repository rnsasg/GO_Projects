#!/bin/bash

# Define the input string
input="ops(t2.select_entity_by_id=1) ops(t2.update_entity_name=1) ops(t2.delete_entity_by_id=1)"

# Split the input string based on whitespace and iterate over each part
index=1
for part in $input
do
    echo "Operation $index: $part"
    ((index++))
done





