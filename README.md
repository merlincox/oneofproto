# oneofproto

Library experimenting with an alternative to long switch type clauses for protobuf `oneof` fields, so that new `oneof` options can be supported by simply bumping the protobuf dependency, thus not requiring new code.

Benchmarks seem to indicate little difference in speed with a large set of oneof options.

go test -bench=. -benchmem -benchtime 20s  
goos: linux  
goarch: amd64
pkg: github.com/merlincox/oneofproto  
cpu: Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz 

| Benchmark                      | Iterations | ns/op | B/op | allocs/op|
|--------------------------------|---|---|---|---|
| BenchmarkToStructpb100-4       | 7993056 | 2998 ns/op | 1128 B/op | 29 allocs/op|
| BenchmarkToStructpbDefer100-4  | 7796260 | 3050 ns/op |1128 B/op | 29 allocs/op|
| BenchmarkToStructpbUnsafe100-4 | 7876123 | 3124 ns/op | 1128 B/op | 29 allocs/op|
| BenchmarkSwitched100-4         | 8142153 | 2960 ns/op | 1128 B/op | 29 allocs/op|
| BenchmarkToStructpb1-4       | 7894995 | 3020 ns/op | 1128 B/op | 29 allocs/op|
| BenchmarkToStructpbDefer1-4  | 7945716 | 3011 ns/op |1128 B/op | 29 allocs/op|
| BenchmarkToStructpbUnsafe1-4 | 8090305| 3005 ns/op | 1128 B/op | 29 allocs/op|
| BenchmarkSwitched1-4         | 8136547 | 2982 ns/op | 1128 B/op | 29 allocs/op|

where the tests ending in 100 use a oneof with a hundred possible values, while those ending in 1 use a oneof with only one value.
