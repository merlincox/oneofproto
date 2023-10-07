# oneofproto

Library experimenting with an alternative to long switch type clauses for protobuf `oneof` fields, so that new `oneof` options can be supported by simply bumping the protobuf dependency, thus not requiring new code.

Benchmarks seem to indicate little difference in speed with a large set of oneof options.

go test -bench=. -benchmem -benchtime 20s  
goos: linux  
goarch: amd64
pkg: github.com/merlincox/oneofproto  
cpu: Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz 

| Benchmark                   | Iterations | ns/op | B/op | allocs/op|
|-----------------------------|---|---|---|---|
| BenchmarkToStructpb-4       | 7832139  | 3092 ns/op | 1128 B/op | 29 allocs/op|
| BenchmarkToStructpbDefer-4  | 7626178  | 3183 ns/op |1128 B/op | 29 allocs/op|
| BenchmarkToStructpbUnsafe-4 | 7684249  | 3154 ns/op | 1128 B/op | 29 allocs/op|
| BenchmarkSwitched-4         | 7759136 | 3088 ns/op | 1128 B/op | 29 allocs/op|

