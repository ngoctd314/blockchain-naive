# Persistence and CLI

## BoltDB

BoltDB is a key/value storage, which means there're no tables like in SQL, no rows, no columns. Instead, data is stored as key-value pairs. Key-value paris are stored in buckets, which are intended to group similar pairs. Thus, in order to get a value, you need to know a bucket and a key.

## Database structure

Bitcoin Cores uses two "buckets" to store data:

1. blocks stores metadata describing all the blocks in a chain
2. chainstate stores the state of a chain, which is all currently unspent transaction outputs and some metadata.