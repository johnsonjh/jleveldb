# JLevelDB

 [![GoReportCard](https://goreportcard.com/badge/github.com/johnsonjh/jleveldb)](https://goreportcard.com/badge/github.com/johnsonjh/jleveldb)
 [![CodacyBadge](https://api.codacy.com/project/badge/Grade/94d1d98a3b4944908b4212edb1adc878)](https://app.codacy.com/gh/johnsonjh/jleveldb?utm_source=github.com&utm_medium=referral&utm_content=johnsonjh/jleveldb&utm_campaign=Badge_Grade_Settings)
 [![Maintainability](https://api.codeclimate.com/v1/badges/7fc331ccc9fb4e4544ae/maintainability)](https://codeclimate.com/github/johnsonjh/jleveldb/maintainability)

---

JLevelDB is an implementation of the
[LevelDB key/value database](https://code.google.com/p/leveldb) in the
[Go programming language](https://golang.org), based on
[GoLevelDB](https://github.com/syndtr/goleveldb)

---

## Requirements

- Go 1.16 or later

---

## Availability

- [GitHub](https://github.com/johnsonjh/jleveldb)

---

## License

- [BSD-2-Clause License](<https://tldrlegal.com/license/bsd-2-clause-license-(freebsd)>)

---

## Examples

- Create or open a database:

```go
// The returned DB instance is safe for concurrent use. Which mean that all
// DB's methods may be called concurrently from multiple goroutine.
db, err := leveldb.OpenFile("path/to/db", nil)
...
defer db.Close()
...
```

---

- Read or modify the database content:

```go
// Remember that the contents of the returned slice should not be modified.
data, err := db.Get([]byte("key"), nil)
...
err = db.Put([]byte("key"), []byte("value"), nil)
...
err = db.Delete([]byte("key"), nil)
...
```

---

- Iterate over database content:

```go
iter := db.NewIterator(nil, nil)
for iter.Next() {
 // Remember that the contents of the returned slice should not be modified, and
 // only valid until the next call to Next.
 key := iter.Key()
 value := iter.Value()
 ...
}
iter.Release()
err = iter.Error()
...
```

---

- Seek-then-Iterate:

```go
iter := db.NewIterator(nil, nil)
for ok := iter.Seek(key); ok; ok = iter.Next() {
 // Use key/value.
 ...
}
iter.Release()
err = iter.Error()
...
```

---

- Iterate over subset of database content:

```go
iter := db.NewIterator(&util.Range{Start: []byte("foo"), Limit: []byte("xoo")}, nil)
for iter.Next() {
 // Use key/value.
 ...
}
iter.Release()
err = iter.Error()
...
```

---

- Iterate over subset of database content with a particular prefix:

```go
iter := db.NewIterator(util.BytesPrefix([]byte("foo-")), nil)
for iter.Next() {
 // Use key/value.
 ...
}
iter.Release()
err = iter.Error()
...
```

---

- Batch writes:

```go
batch := new(leveldb.Batch)
batch.Put([]byte("foo"), []byte("value"))
batch.Put([]byte("bar"), []byte("another value"))
batch.Delete([]byte("baz"))
err = db.Write(batch, nil)
...
```

---

- Use bloom filter:

```go
o := &opt.Options{
 Filter: filter.NewBloomFilter(10),
}
db, err := leveldb.OpenFile("path/to/db", o)
...
defer db.Close()
...
```

---
