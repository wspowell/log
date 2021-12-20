# log

Drop-in replacement for golang `log`.

# Benchmarks

Take with a bucket of salt.

Debug
```
go test -bench=. -benchmem -count=1 -parallel 8 

goos: linux
goarch: amd64
pkg: github.com/wspowell/log
cpu: AMD Ryzen 9 4900HS with Radeon Graphics         
Benchmark_Context_debug_at_error_level-8                                34026336                35.76 ns/op           16 B/op          1 allocs/op
Benchmark_Context_debug_at_error_level_with_complex_parameters-8        11997372                99.84 ns/op           16 B/op          1 allocs/op
Benchmark_Context_error_at_error_level-8                                   70053             17001 ns/op            2822 B/op         36 allocs/op
Benchmark_Context_debug_at_error_level_1000x-8                             34009             35453 ns/op           16000 B/op       1000 allocs/op
Benchmark_Context_error_at_error_level_1000x-8                                73          16915953 ns/op         2822854 B/op      36014 allocs/op
Benchmark_Context_WithContext-8                                           112440             10461 ns/op             600 B/op          9 allocs/op
Benchmark_debug_at_error_level-8                                        1000000000               0.8100 ns/op          0 B/op          0 allocs/op
Benchmark_debug_at_error_level_with_complex_parameters-8                13722502                87.56 ns/op           16 B/op          1 allocs/op
Benchmark_error_at_error_level-8                                          139208              8458 ns/op            2845 B/op         37 allocs/op
Benchmark_debug_at_error_level_1000x-8                                   1459600               822.1 ns/op             0 B/op          0 allocs/op
Benchmark_error_at_error_level_1000x-8                                       139           8612214 ns/op         2845789 B/op      37011 allocs/op
Benchmark_NewLog-8                                                      12922414                93.63 ns/op           80 B/op          2 allocs/op
Benchmark_ConfigCopy_NewLog-8                                           13571282                96.74 ns/op           80 B/op          2 allocs/op
```

Release
```
go test -bench=. -benchmem -count=1 -parallel 8 -tags release

goos: linux
goarch: amd64
pkg: github.com/wspowell/log
cpu: AMD Ryzen 9 4900HS with Radeon Graphics         
Benchmark_Context_debug_at_error_level-8                                35607039                34.30 ns/op           16 B/op          1 allocs/op
Benchmark_Context_debug_at_error_level_with_complex_parameters-8        12886431                92.57 ns/op           16 B/op          1 allocs/op
Benchmark_Context_error_at_error_level-8                                  329538              3485 ns/op            1780 B/op         28 allocs/op
Benchmark_Context_debug_at_error_level_1000x-8                             35020             34362 ns/op           16000 B/op       1000 allocs/op
Benchmark_Context_error_at_error_level_1000x-8                               343           3486383 ns/op         1780613 B/op      28007 allocs/op
Benchmark_Context_WithContext-8                                          2232276               548.3 ns/op           584 B/op          9 allocs/op
Benchmark_debug_at_error_level-8                                        1000000000               0.8128 ns/op          0 B/op          0 allocs/op
Benchmark_debug_at_error_level_with_complex_parameters-8                14779507                81.69 ns/op           16 B/op          1 allocs/op
Benchmark_error_at_error_level-8                                          364293              3354 ns/op            1764 B/op         27 allocs/op
Benchmark_debug_at_error_level_1000x-8                                   1485715               812.1 ns/op             0 B/op          0 allocs/op
Benchmark_error_at_error_level_1000x-8                                       356           3355820 ns/op         1764559 B/op      27006 allocs/op
Benchmark_NewLog-8                                                      12354562                94.14 ns/op           80 B/op          2 allocs/op
Benchmark_ConfigCopy_NewLog-8                                           12822279                93.07 ns/op           80 B/op          2 allocs/op
```