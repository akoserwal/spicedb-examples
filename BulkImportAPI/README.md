# How to Use SpiceDB Bulk Import for Relationships in Go

Using SpiceDB's BulkImport API to import relationship in Spicedb.

# Setup Spicedb
require docker-compose

```make spicedb-up```

## Import Schema using Zed
require spicedb [zed](https://github.com/authzed/zed) client

```
 zed import spicedb/schema-ex.yaml --endpoint localhost:50051 --token foobar --insecure
```

# Run Bulk Import Client

`make run`