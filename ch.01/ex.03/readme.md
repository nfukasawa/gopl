## 練習問題1.3
非効率的な可能性のあるバージョンと `strings.Join` を使ったバージョンとで、実行時間の差を計測しなさい (1.6 節は `time` パッケージの一部を説明していますし、11.4 節では体系的に性能評価を行うためのベンチマークテストの書き方を説明しています)。

## Benchmark Test Result

```
$ go test -bench .
BenchmarkEfficientEcho100-4       	 1000000	      1528 ns/op
BenchmarkEfficientEcho10000-4     	   10000	    129523 ns/op
BenchmarkInefficientEcho100-4     	  100000	     13707 ns/op
BenchmarkInefficientEcho10000-4   	      20	  54828179 ns/op
PASS
ok  	github.com/nfukasawa/gopl/ch.01/ex.03	5.547s
