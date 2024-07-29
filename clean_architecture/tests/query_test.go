package tests

import (
	"fmt"
	"github.com/gocql/gocql"
	"testing"
	"time"
)

func TestQueryPerformance(t *testing.T) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "tagging"
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = 10 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	dataSizes := []int{1000, 10000, 100000, 1000000}
	for _, size := range dataSizes {
		t.Run(fmt.Sprintf("%d Records", size), func(t *testing.T) {
			//GenerateTestData(session, size)

			start := time.Now()
			// Run your queries here, for example:
			QueryTags(session)
			QueryEntities(session)
			// Add other queries as needed
			elapsed := time.Since(start)

			t.Logf("Data size: %d, Total time: %s", size, elapsed)
		})
	}
}
