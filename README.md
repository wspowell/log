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
Benchmark_Context_debug_at_error_level-8                                32912419                36.78 ns/op           16 B/op          1 allocs/op
Benchmark_Context_debug_at_error_level_with_complex_parameters-8        12438609                97.51 ns/op           16 B/op          1 allocs/op
Benchmark_Context_error_at_error_level-8                                   91748             13107 ns/op             576 B/op          6 allocs/op
Benchmark_Context_debug_at_error_level_1000x-8                             32630             36519 ns/op           16000 B/op       1000 allocs/op
Benchmark_Context_error_at_error_level_1000x-8                                92          13093504 ns/op          576674 B/op       6002 allocs/op
Benchmark_Context_WithContext-8                                           115984             10385 ns/op             696 B/op          9 allocs/op
Benchmark_debug_at_error_level-8                                        1000000000               0.7210 ns/op          0 B/op          0 allocs/op
Benchmark_debug_at_error_level_with_complex_parameters-8                14285186                82.57 ns/op           16 B/op          1 allocs/op
Benchmark_error_at_error_level-8                                          242006              4829 ns/op             560 B/op          5 allocs/op
Benchmark_debug_at_error_level_1000x-8                                   1473930               817.5 ns/op             0 B/op          0 allocs/op
Benchmark_error_at_error_level_1000x-8                                       241           4937112 ns/op          560387 B/op       5001 allocs/op
Benchmark_NewLog-8                                                       9259172               127.6 ns/op           176 B/op          2 allocs/op
Benchmark_ConfigCopy_NewLog-8                                            9068628               133.7 ns/op           176 B/op          2 allocs/op
```

Release
```
go test -bench=. -benchmem -count=1 -parallel 8 -tags release

goos: linux
goarch: amd64
pkg: github.com/wspowell/log
cpu: AMD Ryzen 9 4900HS with Radeon Graphics         
Benchmark_Context_debug_at_error_level-8                                32116561                36.51 ns/op           16 B/op          1 allocs/op
Benchmark_Context_debug_at_error_level_with_complex_parameters-8        12438739                92.60 ns/op           16 B/op          1 allocs/op
Benchmark_Context_error_at_error_level-8                                 1837284               647.9 ns/op            32 B/op          2 allocs/op
Benchmark_Context_debug_at_error_level_1000x-8                             32395             36626 ns/op           16000 B/op       1000 allocs/op
Benchmark_Context_error_at_error_level_1000x-8                              1862            649242 ns/op           32023 B/op       2000 allocs/op
Benchmark_Context_WithContext-8                                          1941897               612.2 ns/op           680 B/op          9 allocs/op
Benchmark_debug_at_error_level-8                                        1000000000               0.7162 ns/op          0 B/op          0 allocs/op
Benchmark_debug_at_error_level_with_complex_parameters-8                13970349                83.15 ns/op           16 B/op          1 allocs/op
Benchmark_error_at_error_level-8                                         2212515               545.1 ns/op            16 B/op          1 allocs/op
Benchmark_debug_at_error_level_1000x-8                                   1472370               823.5 ns/op             0 B/op          0 allocs/op
Benchmark_error_at_error_level_1000x-8                                      2221            542200 ns/op           16012 B/op       1000 allocs/op
Benchmark_NewLog-8                                                       8960593               131.1 ns/op           176 B/op          2 allocs/op
Benchmark_ConfigCopy_NewLog-8                                            9373520               134.2 ns/op           176 B/op          2 allocs/op
```