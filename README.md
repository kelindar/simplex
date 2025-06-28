<p align="center">
<img width="300" height="100" src=".github/logo.png" border="0" alt="kelindar/smutex">
<br>

<img src="https://img.shields.io/github/go-mod/go-version/kelindar/simplex" alt="Go Version">
<a href="https://pkg.go.dev/github.com/kelindar/simplex"><img src="https://pkg.go.dev/badge/github.com/kelindar/simplex" alt="PkgGoDev"></a>
<a href="https://goreportcard.com/report/github.com/kelindar/simplex"><img src="https://goreportcard.com/badge/github.com/kelindar/simplex" alt="Go Report Card"></a>
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License"></a>
<a href="https://coveralls.io/github/kelindar/simplex"><img src="https://coveralls.io/repos/github/kelindar/simplex/badge.svg" alt="Coverage"></a>
</p>

## Simplex Noise

This respository contains an experimental implementation of [simplex noise](https://weber.itn.liu.se/~stegu/simplexnoise/simplexnoise.pdf) based on the code from the public domain, found at [weber.itn.liu.se/~stegu/simplexnoise](https://weber.itn.liu.se/~stegu/simplexnoise/SimplexNoise.java). Note that this is not the genuine implementation of [Ken Perlin's simplex noise](https://mrl.cs.nyu.edu/~perlin/noise/) presented at SIGGRAPH 2002.

<p align="center">
<img width="800" height="800" src="examples/terrain.png" border="0" alt="kelindar/simplex">
</p>

## Benchmarks

```
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkNoise/10x10-8       763042   1568 ns/op      0 B/op   0 allocs/op
BenchmarkNoise/100x100-8     7402     159403 ns/op    0 B/op   0 allocs/op
BenchmarkNoise/1000x1000-8   74       15732020 ns/op  0 B/op   0 allocs/op
```

## Contributing

We are open to contributions, feel free to submit a pull request and we'll review it as quickly as we can. This library is maintained by [Roman Atachiants](https://www.linkedin.com/in/atachiants/)

## License

Tile is licensed under the [MIT License](LICENSE.md).
