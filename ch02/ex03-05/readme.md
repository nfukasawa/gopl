## 練習問題 2.3

単一の式の代わりにループを使うように`PopCount`を書き直しなさい。
2つのバージョンの性能を比較しなさい。(11.4節で異なる実装の性能を体系的に比較する方法を説明しています。)

## 練習問題 2.4

引数をビットシフトしながら最下位ビットの検査を64回繰り返すことでビット数を数える`PopCount`のバージョンを作成しなさい。
テーブル参照を行うバージョンと性能を比較しなさい。

## 練習問題 2.5

式`x&(x-1)`は`x`で`1`が設定差入れている最下位ビットをクリアします。
この事実を使ってビット数を数える`PopCount`のバージョンを作成し、その性能を評価しなさい。

## Result

```
$  go test  -bench=. -v github.com/nfukasawa/gopl/ch02/ex03-05/popcount
=== RUN   TestPopCount
--- PASS: TestPopCount (0.00s)
BenchmarkPopCount-4             	2000000000	         0.36 ns/op
BenchmarkPopCountByLoop-4       	50000000	        27.1 ns/op
BenchmarkPopCountByBitShift-4   	20000000	        98.6 ns/op
BenchmarkPopCountByBitClear-4   	50000000	        25.2 ns/op
PASS
ok  	github.com/nfukasawa/gopl/ch02/ex03-05/popcount	5.502s
```