export THREAD=1

echo "###  Tag Table 10K Records ####"

echo "$$$ Tag Table 10K Records Insert (Load) "
/opt/cassandra/tools/bin/cassandra-stress user n=10K profile=t1_tags.yaml "ops(t1.insert=1)" truncate=once -graph file=tags_10K.html title="NVO_Cassandra_POC_TAG_Table_10K_Records" revision="Insert_10K" -rate  threads=$THREAD -log file=tags_10K_write.log

sleep 2
echo "$$$ Tag Table 10K Records Read"

/opt/cassandra/tools/bin/cassandra-stress user n=10K profile=t1_tags.yaml "ops(t1.select_tag_by_id=1)" -graph file=tags_10K.html title="NVO_Cassandra_POC_TAG_Table_10K_Records" revision="Read_10K" -rate  threads=$THREAD -log file=tags_10K_read.log

sleep 2
echo "$$$ Tag Table 10K Records Update"

/opt/cassandra/tools/bin/cassandra-stress user n=10K profile=t1_tags.yaml "ops(t1.update_tag_properties=1)"  -graph file=tags_10K.html title="NVO_Cassandra_POC_TAG_Table_10K_Records" revision="Update_10K" -rate  threads=$THREAD -log file=tags_10K_update.log

sleep 2
echo "$$$ Tag Table 10K Records Delete"

/opt/cassandra/tools/bin/cassandra-stress user n=10K profile=t1_tags.yaml "ops(t1.delete_tag_by_id=1)" -graph file=tags_10K.html title="NVO_Cassandra_POC_TAG_Table_10K_Records" revision="Delete_10K" -rate  threads=$THREAD -log file=tags_10K_delete.logs


sleep 5

echo "###  Tag Table 100K Records ####"

echo "$$$ Tag Table 100K Records Insert (Load) "
/opt/cassandra/tools/bin/cassandra-stress user n=100K profile=t1_tags.yaml "ops(t1.insert=1)" truncate=once -graph file=tags_100K.html title="NVO_Cassandra_POC_TAG_Table_100K_Records" revision="Insert_100K" -rate  threads=$THREAD -log file=tags_100K_write.log

sleep 2
echo "$$$ Tag Table 100K Records  Read"

/opt/cassandra/tools/bin/cassandra-stress user n=100K profile=t1_tags.yaml "ops(t1.select_tag_by_id=1)" -graph file=tags_100K.html title="NVO_Cassandra_POC_TAG_Table_100K_Records" revision="Read_100K" -rate  threads=$THREAD -log file=tags_100K_read.log

sleep 2
echo "$$$ Tag Table 100K Records Update"

/opt/cassandra/tools/bin/cassandra-stress user n=100K profile=t1_tags.yaml "ops(t1.update_tag_properties=1)"  -graph file=tags_100K.html title="NVO_Cassandra_POC_TAG_Table_100K_Records" revision="Update_100K" -rate  threads=$THREAD -log file=tags_100K_update.log

sleep 2
echo "$$$ Tag Table 100K Records Delete "
/opt/cassandra/tools/bin/cassandra-stress user n=100K profile=t1_tags.yaml "ops(t1.delete_tag_by_id=1)" -graph file=tags_100K.html title="NVO_Cassandra_POC_TAG_Table_100K_Records" revision="Delete_100K" -rate  threads=$THREAD -log file=tags_100K_delete.logs

sleep 5
echo "###  Tag Table 1M Records ####"


echo "$$$ Tag Table 1M Records Insert (Load)"
/opt/cassandra/tools/bin/cassandra-stress user n=1M profile=t1_tags.yaml "ops(t1.insert=1)" truncate=once -graph file=tags_1M.html title="NVO_Cassandra_POC_TAG_Table_1M_Records" revision="Insert_1M" -rate  threads=$THREAD -log file=tags_1M_write.log

sleep 2
echo "$$$ Tag Table 1M Records Read"
/opt/cassandra/tools/bin/cassandra-stress user n=1M profile=t1_tags.yaml "ops(t1.select_tag_by_id=1)" -graph file=tags_1M.html title="NVO_Cassandra_POC_TAG_Table_1M_Records" revision="Read_1M" -rate  threads=$THREAD -log file=tags_1M_read.log

sleep 2
echo "$$$ Tag Table 1M Records Update"
/opt/cassandra/tools/bin/cassandra-stress user n=1M profile=t1_tags.yaml "ops(t1.update_tag_properties=1)"  -graph file=tags_1M.html title="NVO_Cassandra_POC_TAG_Table_1M_Records" revision="Update_1M" -rate  threads=$THREAD -log file=tags_1M_update.log

sleep 2
echo "$$$ Tag Table 1m Records Delete"
/opt/cassandra/tools/bin/cassandra-stress user n=1M profile=t1_tags.yaml "ops(t1.delete_tag_by_id=1)" -graph file=tags_1M.html title="NVO_Cassandra_POC_TAG_Table_1M_Records" revision="Delete_1M" -rate  threads=$THREAD -log file=tags_1M_delete.logs