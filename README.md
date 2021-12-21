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
Benchmark_Context_debug_at_error_level-8                                169843735                7.041 ns/op           0 B/op          0 allocs/op
Benchmark_Context_debug_at_error_level_with_complex_parameters-8        12824102                92.68 ns/op           16 B/op          1 allocs/op
Benchmark_Context_error_at_error_level-8                                   88299             13264 ns/op             656 B/op          7 allocs/op
Benchmark_Context_debug_at_error_level_1000x-8                            174613              6776 ns/op               0 B/op          0 allocs/op
Benchmark_Context_error_at_error_level_1000x-8                                85          13269595 ns/op          656719 B/op       7002 allocs/op
Benchmark_Context_WithContext-8                                           100618             12081 ns/op            1824 B/op         15 allocs/op
Benchmark_debug_at_error_level-8                                        1000000000               0.8259 ns/op          0 B/op          0 allocs/op
Benchmark_debug_at_error_level_with_complex_parameters-8                13513452                83.05 ns/op           16 B/op          1 allocs/op
Benchmark_error_at_error_level-8                                          256467              4697 ns/op             584 B/op          7 allocs/op
Benchmark_debug_at_error_level_1000x-8                                   1468938               810.7 ns/op             0 B/op          0 allocs/op
Benchmark_error_at_error_level_1000x-8                                       255           4700730 ns/op          584391 B/op       7001 allocs/op
Benchmark_error_at_error_level_1000x_zerolog-8                              2242            536708 ns/op           16011 B/op       1000 allocs/op
Benchmark_NewLog-8                                                       1458918               802.5 ns/op          1320 B/op          9 allocs/op
Benchmark_ConfigCopy_NewLog-8                                            1466047               800.4 ns/op          1320 B/op          9 allocs/op
```

Release
```
go test -bench=. -benchmem -count=1 -parallel 8 -tags release

goos: linux
goarch: amd64
pkg: github.com/wspowell/log
cpu: AMD Ryzen 9 4900HS with Radeon Graphics         
Benchmark_Context_debug_at_error_level-8                                164714614                7.318 ns/op           0 B/op          0 allocs/op
Benchmark_Context_debug_at_error_level_with_complex_parameters-8        12601136                93.15 ns/op           16 B/op          1 allocs/op
Benchmark_Context_error_at_error_level-8                                 1803454               661.4 ns/op            32 B/op          2 allocs/op
Benchmark_Context_debug_at_error_level_1000x-8                            168908              7033 ns/op               0 B/op          0 allocs/op
Benchmark_Context_error_at_error_level_1000x-8                              1803            661097 ns/op           32024 B/op       2000 allocs/op
Benchmark_Context_WithContext-8                                           890742              1264 ns/op            1808 B/op         15 allocs/op
Benchmark_debug_at_error_level-8                                        1000000000               0.8156 ns/op          0 B/op          0 allocs/op
Benchmark_debug_at_error_level_with_complex_parameters-8                13834814                85.39 ns/op           16 B/op          1 allocs/op
Benchmark_error_at_error_level-8                                         2191300               546.1 ns/op            16 B/op          1 allocs/op
Benchmark_debug_at_error_level_1000x-8                                   1455393               812.9 ns/op             0 B/op          0 allocs/op
Benchmark_error_at_error_level_1000x-8                                      2176            548146 ns/op           16011 B/op       1000 allocs/op
Benchmark_error_at_error_level_1000x_zerolog-8                              2199            548922 ns/op           16010 B/op       1000 allocs/op
Benchmark_NewLog-8                                                       1479292               806.5 ns/op          1320 B/op          9 allocs/op
Benchmark_ConfigCopy_NewLog-8                                            1488816               807.2 ns/op          1320 B/op          9 allocs/op
```