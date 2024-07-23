#!/bin/bash

#!/bin/bash

# Define arrays for records and thread counts
RECORDS=("1K" "2k")
THREADS=("1" "2")

# Default values for TABLE_NAME and PROFILE
TABLE_NAME="entities"
PROFILE="t2_entities.yaml"

# Outer loop for records
for TOTAL_RECORD in "${RECORDS[@]}"
do
    echo "### Running $TABLE_NAME Table $TOTAL_RECORD Records ####"

    # Inner loop for thread counts
    for THREAD in "${THREADS[@]}"
    do
        echo "!!! $TABLE_NAME Table Load $TOTAL_RECORD Records with $THREAD Threads"
        /opt/cassandra/tools/bin/cassandra-stress user n=$TOTAL_RECORD profile=$PROFILE "ops(t2.insert=1)" truncate=once -graph file="${TABLE_NAME}_${TOTAL_RECORD}_threads_${THREAD}.html" title="NVO_Cassandra_POC_${TABLE_NAME}_Table_${TOTAL_RECORD}_Records_${THREAD}_Threads" revision="Insert" -rate threads=$THREAD -log file="${TABLE_NAME}_${TOTAL_RECORD}_write_${THREAD}_threads.log"

        sleep 2
        echo "!!! $TABLE_NAME Table $TOTAL_RECORD Records Read with $THREAD Threads"
        /opt/cassandra/tools/bin/cassandra-stress user n=$TOTAL_RECORD profile=$PROFILE "ops(t2.select_entity_by_id=1)" -graph file="${TABLE_NAME}_${TOTAL_RECORD}_threads_${THREAD}.html" title="NVO_Cassandra_POC_${TABLE_NAME}_Table_${TOTAL_RECORD}_Records_${THREAD}_Threads" revision="Read" -rate threads=$THREAD -log file="${TABLE_NAME}_${TOTAL_RECORD}_read_${THREAD}_threads.log"

        sleep 2
        echo "!!! $TABLE_NAME Table $TOTAL_RECORD Records Update with $THREAD Threads"
        /opt/cassandra/tools/bin/cassandra-stress user n=$TOTAL_RECORD profile=$PROFILE "ops(t2.update_entity_name=1)" -graph file="${TABLE_NAME}_${TOTAL_RECORD}_threads_${THREAD}.html" title="NVO_Cassandra_POC_${TABLE_NAME}_Table_${TOTAL_RECORD}_Records_${THREAD}_Threads" revision="Update" -rate threads=$THREAD -log file="${TABLE_NAME}_${TOTAL_RECORD}_update_${THREAD}_threads.log"

        sleep 2
        echo "!!! $TABLE_NAME Table $TOTAL_RECORD Records Delete with $THREAD Threads"
        /opt/cassandra/tools/bin/cassandra-stress user n=$TOTAL_RECORD profile=$PROFILE "ops(t2.delete_entity_by_id=1)" -graph file="${TABLE_NAME}_${TOTAL_RECORD}_threads_${THREAD}.html" title="NVO_Cassandra_POC_${TABLE_NAME}_Table_${TOTAL_RECORD}_Records_${THREAD}_Threads" revision="Delete" -rate threads=$THREAD -log file="${TABLE_NAME}_${TOTAL_RECORD}_delete_${THREAD}_threads.log"

        sleep 5  # Adjust sleep time as necessary
    done
done


#export THREAD=1
#export TABLE_NAME="entities"
#export PROFILE="t2_entities.yaml"
#export TOTAL_RECORD="10K"  # Assuming you want to keep it as a string
#
#echo "### $TABLE_NAME Table $TOTAL_RECORD Records ####"
#
#echo "!!! $TABLE_NAME Table Load $TOTAL_RECORD Records"
#/opt/cassandra/tools/bin/cassandra-stress user n=$TOTAL_RECORD profile=$PROFILE "ops(t2.insert=1)" truncate=once -graph file=entities_${TOTAL_RECORD}.html title="NVO_Cassandra_POC_${TABLE_NAME}_Table_${TOTAL_RECORD}_Records" revision="Insert_${TOTAL_RECORD}" -rate threads=$THREAD -log file=entities_${TOTAL_RECORD}_write.log
#
#sleep 2
#echo "!!! $TABLE_NAME Table $TOTAL_RECORD Records Read"
#
#/opt/cassandra/tools/bin/cassandra-stress user n=${TOTAL_RECORD} profile=${PROFILE} "ops(t2.select_entity_by_id=1)" -graph file=entities_${TOTAL_RECORD}.html title="NVO_Cassandra_POC_${TABLE_NAME}_Table_${TOTAL_RECORD}_Records" revision="Read_${TOTAL_RECORD}" -rate threads=${THREAD} -log file=entities_${TOTAL_RECORD}_read.log
#
#sleep 2
#echo "!!! $TABLE_NAME Table $TOTAL_RECORD Records Update"
#
#/opt/cassandra/tools/bin/cassandra-stress user n=${TOTAL_RECORD} profile=${PROFILE} "ops(t2.update_entity_name=1)" -graph file=entities_${TOTAL_RECORD}.html title="NVO_Cassandra_POC_${TABLE_NAME}_Table_${TOTAL_RECORD}_Records" revision="Update_${TOTAL_RECORD}" -rate threads=${THREAD} -log file=entities_${TOTAL_RECORD}_update.log
#
#sleep 2
#echo "!!! $TABLE_NAME Table $TOTAL_RECORD Records Delete"
#/opt/cassandra/tools/bin/cassandra-stress user n=${TOTAL_RECORD} profile=${PROFILE} "ops(t2.delete_entity_by_id=1)" -graph file=entities_${TOTAL_RECORD}.html title="NVO_Cassandra_POC_${TABLE_NAME}_Table_${TOTAL_RECORD}_Records" revision="Delete_${TOTAL_RECORD}" -rate threads=${THREAD} -log file=entities_${TOTAL_RECORD}_delete.log
