# log

Replacement for golang `log`.

# Benchmarks

Take with a bucket of salt.

Debug
```
go test -bench=. -benchmem -count=1 -parallel 8

goos: linux
goarch: amd64
pkg: github.com/wspowell/log
cpu: AMD Ryzen 9 4900HS with Radeon Graphics         
Benchmark_Context_debug_at_error_level-8                                171908761                6.958 ns/op           0 B/op          0 allocs/op
Benchmark_Context_debug_at_error_level_with_complex_parameters-8        12861198                92.07 ns/op           16 B/op          1 allocs/op
Benchmark_Context_error_at_error_level-8                                   89876             12965 ns/op             656 B/op          7 allocs/op
Benchmark_Context_debug_at_error_level_1000x-8                            176293              6803 ns/op               0 B/op          0 allocs/op
Benchmark_Context_error_at_error_level_1000x-8                                86          13025992 ns/op          656722 B/op       7002 allocs/op
Benchmark_Context_WithContext-8                                           104426             11603 ns/op            1824 B/op         15 allocs/op
Benchmark_Context_Tag-8                                                   145516              8186 ns/op              88 B/op          1 allocs/op
Benchmark_debug_at_error_level-8                                        1000000000               0.8027 ns/op          0 B/op          0 allocs/op
Benchmark_debug_at_error_level_with_complex_parameters-8                14631933                82.25 ns/op           16 B/op          1 allocs/op
Benchmark_error_at_error_level-8                                          254371              4697 ns/op             584 B/op          7 allocs/op
Benchmark_debug_at_error_level_1000x-8                                   1625730               741.1 ns/op             0 B/op          0 allocs/op
Benchmark_error_at_error_level_1000x-8                                       254           4695899 ns/op          584399 B/op       7001 allocs/op
Benchmark_error_at_error_level_1000x_zerolog-8                              2221            531979 ns/op           16011 B/op       1000 allocs/op
Benchmark_NewLog-8                                                       1487022               797.6 ns/op          1320 B/op          9 allocs/op
Benchmark_ConfigCopy_NewLog-8                                            1480128               812.2 ns/op          1320 B/op          9 allocs/op
Benchmark_Tag-8                                                         11105218               101.8 ns/op           102 B/op          1 allocs/op
```

Release
```
go test -bench=. -benchmem -count=1 -parallel 8 -tags release

goos: linux
goarch: amd64
pkg: github.com/wspowell/log
cpu: AMD Ryzen 9 4900HS with Radeon Graphics         
Benchmark_Context_debug_at_error_level-8                                165479329                7.276 ns/op           0 B/op          0 allocs/op
Benchmark_Context_debug_at_error_level_with_complex_parameters-8        13082940                92.82 ns/op           16 B/op          1 allocs/op
Benchmark_Context_error_at_error_level-8                                 1800238               666.9 ns/op            32 B/op          2 allocs/op
Benchmark_Context_debug_at_error_level_1000x-8                            170800              6969 ns/op               0 B/op          0 allocs/op
Benchmark_Context_error_at_error_level_1000x-8                              1809            666134 ns/op           32023 B/op       2000 allocs/op
Benchmark_Context_WithContext-8                                           957746              1290 ns/op            1808 B/op         15 allocs/op
Benchmark_Context_Tag-8                                                  7357278               156.0 ns/op            99 B/op          1 allocs/op
Benchmark_debug_at_error_level-8                                        1000000000               0.8073 ns/op          0 B/op          0 allocs/op
Benchmark_debug_at_error_level_with_complex_parameters-8                14169472                84.50 ns/op           16 B/op          1 allocs/op
Benchmark_error_at_error_level-8                                         2184783               544.2 ns/op            16 B/op          1 allocs/op
Benchmark_debug_at_error_level_1000x-8                                   1605920               757.2 ns/op             0 B/op          0 allocs/op
Benchmark_error_at_error_level_1000x-8                                      2229            545037 ns/op           16012 B/op       1000 allocs/op
Benchmark_error_at_error_level_1000x_zerolog-8                              2226            532053 ns/op           16012 B/op       1000 allocs/op
Benchmark_NewLog-8                                                       1425229               821.0 ns/op          1320 B/op          9 allocs/op
Benchmark_ConfigCopy_NewLog-8                                            1458351               814.2 ns/op          1320 B/op          9 allocs/op
Benchmark_Tag-8                                                         11095915               104.3 ns/op           102 B/op          1 allocs/op
```