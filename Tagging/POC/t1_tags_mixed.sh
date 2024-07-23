#!/bin/bash
export THREAD=1

####

echo "###  Tag Table 10K Records ####"
echo "$$$ Tag Table Load 10K Records mix"

/opt/cassandra/tools/bin/cassandra-stress user n=10K profile=t1_tags.yaml "ops(t1.insert=1)" truncate=once -graph file=tags_10K_mix.html title="NVO_Cassandra_POC_TAG_Table_10K_Records_MIX_Operation" revision="Load_10K" -rate  threads=$THREAD -log file=tags_10K_load_operation.log

sleep 2
/opt/cassandra/tools/bin/cassandra-stress user profile=t1_tags.yaml duration=2m "ops(t1.select_tag_by_id=1,t1.update_tag_properties=1,t1.delete_tag_by_id=1,t1.insert_tag=1)" -graph file=tags_10K_mix.html title="NVO_Cassandra_POC_TAG_Table_10K_Records_MIX_Operation" revision="MIXED_10K" -rate  threads=$THREAD -log file=tags_10K_mix_operation.log

####

sleep 5
echo "###  Tag Table 100K Records ####"
echo "$$$ Tag Table Load 100K Records"

echo "$$$ Tag Table 10K Records Insert (Load) "
/opt/cassandra/tools/bin/cassandra-stress user n=100K profile=t1_tags.yaml "ops(t1.insert=1)" truncate=once -graph file=tags_100K_mix.html title="NVO_Cassandra_POC_TAG_Table_100K_Records_MIX_Operation" revision="Load_10K" -rate  threads=$THREAD -log file=tags_100K_load_operation.log

sleep 2

/opt/cassandra/tools/bin/cassandra-stress user profile=t1_tags.yaml duration=2m "ops(t1.select_tag_by_id=1,t1.update_tag_properties=1,t1.delete_tag_by_id=1,t1.insert_tag=1)" -graph file=tags_100K_mix.html title="NVO_Cassandra_POC_TAG_Table_100K_Records_MIX_Operation" revision="MIXED_10K" -rate  threads=$THREAD -log file=tags_100K_mix_operation.log

####

sleep 5
echo "###  Tag Table 10K Records ####"
echo "$$$ Tag Table Load 10K Records"

echo "$$$ Tag Table 10K Records Insert (Load) "
/opt/cassandra/tools/bin/cassandra-stress user n=10M profile=t1_tags.yaml "ops(t1.insert=1)" truncate=once -graph file=tags_1M_mix.html title="NVO_Cassandra_POC_TAG_Table_10M_Records_MIX_Operation" revision="LOAD_1M" -rate  threads=$THREAD -log file=tags_10K_load_operation.log

sleep 2

/opt/cassandra/tools/bin/cassandra-stress user profile=t1_tags.yaml duration=2m "ops(t1.select_tag_by_id=1,t1.update_tag_properties=1,t1.delete_tag_by_id=1,t1.insert_tag=1)" -graph file=tags_1M_mix.html title="NVO_Cassandra_POC_TAG_Table_10M_Records_MIX_Operation" revision="Insert_10K" -rate  threads=$THREAD -log file=tags_1M_mix_operation.log




