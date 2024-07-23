#!/bin/bash

### TAG Table
#SPEC_NAME="t1"
#TABLE_NAME="tags"
#OPERATIONS="t1.select_tag_by_id=1,t1.update_tag_properties=1,t1.delete_tag_by_id=1"
#
### Entities table
#SPEC_NAME="t2"
#TABLE_NAME="entities"
#OPERATIONS="t2.select_entity_by_id=1,t2.update_entity_name=1,t2.delete_entity_by_id=1"
#
### TAG Entities Table
#SPEC_NAME="t3"
#TABLE_NAME="tag_entities"
#OPERATIONS="t3.select_entities_by_tag=1,t3.delete_tag_entity=1"

# Define arrays for records and thread counts
RECORDS=("10")
THREADS=("1")


SETS=(
    # Set 1: TAG Table
    "t1 tags t1.select_tag_by_id=1,t1.update_tag_properties=1,t1.delete_tag_by_id=1"

    # Set 2: Entities table
    "t2 entities t2.select_entity_by_id=1,t2.update_entity_name=1,t2.delete_entity_by_id=1"

    # Set 3: TAG Entities Table
    "t3 tag_entities t3.select_entities_by_tag=1,t3.delete_tag_entity=1"
)

# Iterate through each set
for set in "${SETS[@]}"
do
    # Split the set into SPEC_NAME, TABLE_NAME, and OPERATIONS
    read -r SPEC_NAME TABLE_NAME OPS <<< "$set"
    echo "Processing Set: SPEC_NAME=$SPEC_NAME, TABLE_NAME=$TABLE_NAME"
    PROFILE="${SPEC_NAME}_${TABLE_NAME}.yaml"
    export PROFILE
    # Outer loop for records
    for TOTAL_RECORD in "${RECORDS[@]}"
    do
        echo "### Running $TABLE_NAME Table $TOTAL_RECORD Records ####"

        # Inner loop for thread counts
        for THREAD in "${THREADS[@]}"
        do
            echo "!!! $TABLE_NAME Table Load $TOTAL_RECORD Records with $THREAD Threads"
            /opt/cassandra/tools/bin/cassandra-stress user n=$TOTAL_RECORD profile=$PROFILE "ops(${SPEC_NAME}.insert=1)" truncate=once -graph file="${TABLE_NAME}_${TOTAL_RECORD}_threads_${THREAD}_MIXED.html" title="NVO_Cassandra_POC_${TABLE_NAME}_Table_${TOTAL_RECORD}_Records_${THREAD}_Threads" revision="Load" -rate threads=$THREAD -log file="${TABLE_NAME}_${TOTAL_RECORD}_write_${THREAD}_threads.log"

            echo "!!! $TABLE_NAME Table $TOTAL_RECORD Records $OPS with $THREAD Threads"
            /opt/cassandra/tools/bin/cassandra-stress user n=$TOTAL_RECORD profile=$PROFILE "ops(${OPS})" -graph file="${TABLE_NAME}_${TOTAL_RECORD}_threads_${THREAD}_MIXED.html" title="NVO_Cassandra_POC_${TABLE_NAME}_Table_${TOTAL_RECORD}_Records_${THREAD}_Threads" revision=${TABLE_NAME}_MIXED -rate threads=$THREAD -log file="${TABLE_NAME}_${TOTAL_RECORD}_MIXED_${THREAD}_threads.log"

            sleep 5  # Adjust sleep time as necessary
        done
    done
done