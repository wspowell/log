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
Benchmark_Context_debug_at_error_level-8                                32588680                37.03 ns/op           16 B/op          1 allocs/op
Benchmark_Context_debug_at_error_level_with_complex_parameters-8        13317145                90.60 ns/op           16 B/op          1 allocs/op
Benchmark_Context_error_at_error_level-8                                   86198             14013 ns/op             993 B/op         12 allocs/op
Benchmark_Context_debug_at_error_level_1000x-8                             32532             36888 ns/op           16000 B/op       1000 allocs/op
Benchmark_Context_error_at_error_level_1000x-8                                87          14073357 ns/op          993100 B/op      12003 allocs/op
Benchmark_Context_WithContext-8                                           115440             10533 ns/op             696 B/op          9 allocs/op
Benchmark_debug_at_error_level-8                                        1000000000               0.7265 ns/op          0 B/op          0 allocs/op
Benchmark_debug_at_error_level_with_complex_parameters-8                14402580                82.92 ns/op           16 B/op          1 allocs/op
Benchmark_error_at_error_level-8                                          201639              5762 ns/op             976 B/op         11 allocs/op
Benchmark_debug_at_error_level_1000x-8                                   1464444               822.4 ns/op             0 B/op          0 allocs/op
Benchmark_error_at_error_level_1000x-8                                       204           5860441 ns/op          976656 B/op      11001 allocs/op
Benchmark_NewLog-8                                                       9203222               127.2 ns/op           176 B/op          2 allocs/op
Benchmark_ConfigCopy_NewLog-8                                            8700562               133.9 ns/op           176 B/op          2 allocs/op
```

Release
```
go test -bench=. -benchmem -count=1 -parallel 8 -tags release

goos: linux
goarch: amd64
pkg: github.com/wspowell/log
cpu: AMD Ryzen 9 4900HS with Radeon Graphics         
Benchmark_Context_debug_at_error_level-8                                32978175                37.25 ns/op           16 B/op          1 allocs/op
Benchmark_Context_debug_at_error_level_with_complex_parameters-8        12733068                92.77 ns/op           16 B/op          1 allocs/op
Benchmark_Context_error_at_error_level-8                                  906562              1221 ns/op             448 B/op          8 allocs/op
Benchmark_Context_debug_at_error_level_1000x-8                             30084             37780 ns/op           16000 B/op       1000 allocs/op
Benchmark_Context_error_at_error_level_1000x-8                               984           1220583 ns/op          448302 B/op       8000 allocs/op
Benchmark_Context_WithContext-8                                          2001898               613.4 ns/op           680 B/op          9 allocs/op
Benchmark_debug_at_error_level-8                                        1000000000               0.7295 ns/op          0 B/op          0 allocs/op
Benchmark_debug_at_error_level_with_complex_parameters-8                14075584                84.39 ns/op           16 B/op          1 allocs/op
Benchmark_error_at_error_level-8                                          989649              1110 ns/op             432 B/op          7 allocs/op
Benchmark_debug_at_error_level_1000x-8                                   1445348               822.4 ns/op             0 B/op          0 allocs/op
Benchmark_error_at_error_level_1000x-8                                      1101           1101916 ns/op          432291 B/op       7000 allocs/op
Benchmark_NewLog-8                                                       9491328               129.4 ns/op           176 B/op          2 allocs/op
Benchmark_ConfigCopy_NewLog-8                                            9184879               128.8 ns/op           176 B/op          2 allocs/op
```