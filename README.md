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
Benchmark_Context_debug_at_error_level-8                                33614010                32.99 ns/op           16 B/op          1 allocs/op
Benchmark_Context_debug_at_error_level_with_complex_parameters-8        12627285                94.81 ns/op           16 B/op          1 allocs/op
Benchmark_Context_error_at_error_level-8                                   79628             14702 ns/op            2822 B/op         36 allocs/op
Benchmark_Context_debug_at_error_level_1000x-8                             36200             33148 ns/op           16000 B/op       1000 allocs/op
Benchmark_Context_error_at_error_level_1000x-8                                80          14707341 ns/op         2822850 B/op      36014 allocs/op
Benchmark_Context_WithContext-8                                           134556              8960 ns/op             560 B/op          8 allocs/op
Benchmark_debug_at_error_level-8                                        1000000000               0.8156 ns/op          0 B/op          0 allocs/op
Benchmark_debug_at_error_level_with_complex_parameters-8                13955499                86.07 ns/op           16 B/op          1 allocs/op
Benchmark_error_at_error_level-8                                          151096              7913 ns/op            2845 B/op         37 allocs/op
Benchmark_debug_at_error_level_1000x-8                                   1458595               828.5 ns/op             0 B/op          0 allocs/op
Benchmark_error_at_error_level_1000x-8                                       150           7984927 ns/op         2845785 B/op      37011 allocs/op
Benchmark_NewLog-8                                                      13860189                88.34 ns/op           80 B/op          2 allocs/op
Benchmark_ConfigCopy_NewLog-8                                           13694815                89.02 ns/op           80 B/op          2 allocs/op
```

Release
```
go test -bench=. -benchmem -count=1 -parallel 8 -tags release

goos: linux
goarch: amd64
pkg: github.com/wspowell/log
cpu: AMD Ryzen 9 4900HS with Radeon Graphics         
Benchmark_Context_debug_at_error_level-8                                35668320                33.14 ns/op           16 B/op          1 allocs/op
Benchmark_Context_debug_at_error_level_with_complex_parameters-8        12321428                95.08 ns/op           16 B/op          1 allocs/op
Benchmark_Context_error_at_error_level-8                                  350462              3415 ns/op            1780 B/op         28 allocs/op
Benchmark_Context_debug_at_error_level_1000x-8                             35781             33826 ns/op           16000 B/op       1000 allocs/op
Benchmark_Context_error_at_error_level_1000x-8                               352           3356288 ns/op         1780588 B/op      28006 allocs/op
Benchmark_Context_WithContext-8                                          2415493               499.9 ns/op           552 B/op          8 allocs/op
Benchmark_debug_at_error_level-8                                        1000000000               0.8138 ns/op          0 B/op          0 allocs/op
Benchmark_debug_at_error_level_with_complex_parameters-8                14369535                85.78 ns/op           16 B/op          1 allocs/op
Benchmark_error_at_error_level-8                                          369042              3292 ns/op            1764 B/op         27 allocs/op
Benchmark_debug_at_error_level_1000x-8                                   1449620               817.1 ns/op             0 B/op          0 allocs/op
Benchmark_error_at_error_level_1000x-8                                       368           3253287 ns/op         1764582 B/op      27006 allocs/op
Benchmark_NewLog-8                                                      13962108                86.81 ns/op           80 B/op          2 allocs/op
Benchmark_ConfigCopy_NewLog-8                                           13366988                87.33 ns/op           80 B/op          2 allocs/op
```