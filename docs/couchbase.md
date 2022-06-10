## couchbase 101

<TODO>

```go
// clusters are like bunch of couchbase(running containers) servers/nodes.
	cluster, err := gocb.Connect(*scheme+*connStr, clusterOpts)
	if err != nil {
		panic(err)
	}

	// Create a bucket instance, which we'll need for access to scopes and collections.
	// buckets are like dbs.
	bucket := cluster.Bucket(travelSampleBucketName)

	app := &TravelSampleApp{
		cluster: cluster,
		bucket:  bucket,
		logger:  logrusLogger,
	}

```

![DynamoDb Partitions](images/couchbase-internal-architecture.png "DynamoDb Partitions") 


### resources:
- [Everything You Need To Know About Couchbase Architecture](https://dzone.com/articles/couchbase-architecture-deep)
- [Using Couchbase with microservices](https://blog.couchbase.com/microservices-architecture-in-couchbase/)