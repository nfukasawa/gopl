## 練習問題 1.10

大量のデータを生成するウェブサイトを見つけなさい。
報告される時間が大きく変化するかを調べるために`fetchall`を2回続けて実行し、キャッシュされているかどうかを調査しなさい。
毎回同じ内容でしょうか。
`fetchall`を修正して、その出力をファイルへ保存するようにして調べられるようにしなさい

## Result

```
ex10 $ go run main.go https://en.wikipedia.org/wiki/Graham%27s_number
0.77s   125599  https://en.wikipedia.org/wiki/Graham%27s_number
0.90s elapsed
ex10 $ go run main.go https://en.wikipedia.org/wiki/Graham%27s_number
0.57s   125599  https://en.wikipedia.org/wiki/Graham%27s_number
0.69s elapsed
ex10 $ go run main.go https://en.wikipedia.org/wiki/Graham%27s_number
0.57s   125599  https://en.wikipedia.org/wiki/Graham%27s_number
0.69s elapsed
```

微妙であるが、2回目以降早くなっているのでキャッシュされているように見えなくもない