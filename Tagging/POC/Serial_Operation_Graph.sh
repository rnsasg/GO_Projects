#!/bin/bash

# Define arrays for records and thread counts
RECORDS=("1K")
THREADS=("1" "5")

SETS=(
    # Set 1: TAG Table
    "t1 tags ops(t1.select_tag_by_id=1) ops(t1.update_tag_properties=1) ops(t1.delete_tag_by_id=1)"

    # Set 2: Entities table
    "t2 entities ops(t2.select_entity_by_id=1) ops(t2.update_entity_name=1) ops(t2.delete_entity_by_id=1)"

    # Set 3: TAG Entities Table
     "t3 tag_entities ops(t3.select_entities_by_tag=1) ops(t3.delete_tag_entity=1)"
)


# Iterate through each set
for set in "${SETS[@]}"
do
    # Split the set into SPEC_NAME, TABLE_NAME, and OPERATIONS
    read -r SPEC_NAME TABLE_NAME OPERATIONS <<< "$set"
    echo "Processing Set: SPEC_NAME=$SPEC_NAME, TABLE_NAME=$TABLE_NAME OPERATIONS=$OPERATIONS"
    PROFILE="${SPEC_NAME}_${TABLE_NAME}.yaml"
    # Outer loop for records
    for TOTAL_RECORD in "${RECORDS[@]}"
    do
        echo "### Running $TABLE_NAME Table $TOTAL_RECORD Records ####"

        # Inner loop for thread counts
        for THREAD in "${THREADS[@]}"
        do
            echo "!!! $TABLE_NAME Table Load $TOTAL_RECORD Records with $THREAD Threads"
            /opt/cassandra/tools/bin/cassandra-stress user n=$TOTAL_RECORD profile=$PROFILE "ops(${SPEC_NAME}.insert=1)" truncate=once -graph file="${TABLE_NAME}_${TOTAL_RECORD}_threads_${THREAD}.html" title="NVO_Cassandra_POC_${TABLE_NAME}_Table_${TOTAL_RECORD}_Records_${THREAD}_Threads" revision="Insert" -rate threads=$THREAD -log file="${TABLE_NAME}_${TOTAL_RECORD}_write_${THREAD}_threads.log"

#            index=1
#            for part in $input
#            do
#                echo "Operation $index: $part"
#                ((index++))
#            done

            ## This loop is for Queries
            index=1
            for OPS in $OPERATIONS
            do
              sleep 2
              echo "!!! $TABLE_NAME Table $TOTAL_RECORD Records $OPS with $THREAD Threads"
              /opt/cassandra/tools/bin/cassandra-stress user n=$TOTAL_RECORD profile=$PROFILE "${OPS}" -graph file="${TABLE_NAME}_${TOTAL_RECORD}_threads_${THREAD}.html" title="NVO_Cassandra_POC_${TABLE_NAME}_Table_${TOTAL_RECORD}_Records_${THREAD}_Threads" revision=$index -rate threads=$THREAD -log file="${TABLE_NAME}_${TOTAL_RECORD}_${index}_${THREAD}_threads.log"
              ((index++))
            done

            sleep 5  # Adjust sleep time as necessary
        done
    done
     sleep 30
done