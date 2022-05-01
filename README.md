The backend code that handles all the meta data for a user session.

This does not contain the code for rtmp handling because that is a completely different service

Updates and connects to two databases:
1. Local key-value store ([Badger](https://github.com/dgraph-io/badger))
2. DynamoDB 

Also handles case-based deletion on S3