# implement etcd server

## design?

### apis/guarantees?
- put key
- get key
  - get by revision
  - get latest
  - get range (low priority)
```json
{
  "header": {
    "cluster_id": 14841639068965180000,
    "member_id": 10276657743932975000,
    "revision": 10,
    "raft_term": 2
  },
  "kvs": [
    {
      "key": "Zm9v",
      "create_revision": 2,
      "mod_revision": 2,
      "version": 1,
      "value": "YmFy"
    }
  ],
  "count": 1
}
```
- watch (low priority)
- multi-nodes
  - get list node
  - get status node

### multi-nodes
- join/leave cluster
- route requests
- replicate from primary to secondary
- elect primary node
- join/rejoin: get data from other nodes
- raft stuff

## implement

### phase 1

- database: try sqlite first, if not ok then just use pg container
- data models:
  1. one table: revision, key, value, create_revision, mod_revision, version
    - revision, key, value: ok
    - delete count as one write? -> yes
    - version: get latest, if (latest==null || latest==tomb) then version = 1 else version++?
      -> slow insert? -> one-sql-query-able -> is get latest fast? -> nope
    - create_revision, mod_revision: can get from previous too
  2. two tables:
    - need transaction
  3. homemade db
-> try 1 first

- how people imple database in Go?: https://dgraph.io/blog/post/badger/

- gprc server
- client

### phase 2

raft

### phase 3

multi-nodes

### phase 4
