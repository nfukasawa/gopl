## 練習問題 1.11

`alexa.com`にある上位100万件のウェブサイトのように、より長い引数リストで`fetchall`を試しなさい。
あるウェブサイトが応答しない場合には、プログラムはどのように振る舞うでしょうか。(そのような場合に対処するための機構は8.9節で説明されています。)

## Result

`alexa.com`のTop20を`fetchall`で取得

```
ex11 $ sh fetchall.sh
+ go run main.go http://google.com http://youtube.com http://facebook.com http://baidu.com http://wikipedia.org http://yahoo.com http://reddit.com http://google.co.in http://qq.com http://taobao.com http://twitter.com http://amazon.com http://sohu.com http://google.co.jp http://live.com http://tmall.com http://vk.com http://instagram.com http://sina.com.cn http://360.cn
0.26s    10563  http://google.com
0.29s    10600  http://google.co.jp
0.33s    14455  http://google.co.in
0.36s       81  http://baidu.com
0.43s   249305  http://qq.com
0.45s   559714  http://youtube.com
0.68s   327715  http://twitter.com
0.96s   590985  http://sina.com.cn
1.18s   147548  http://facebook.com
1.17s   428329  http://sohu.com
1.21s   144675  http://reddit.com
1.23s   470793  http://yahoo.com
1.57s     9549  http://instagram.com
1.82s    86294  http://wikipedia.org
1.98s   224039  http://tmall.com
2.23s   233700  http://amazon.com
2.89s     6563  http://vk.com
3.32s    15606  http://live.com
4.14s    61551  http://taobao.com
5.34s   291217  http://360.cn
6.79s elapsed
```

応答しないウェブサイト`noressrv`を作成して試行

```
ex11 $ sh fetchall_nores.sh &
[1] 52365
+ go run noressrv/main.go
+ go run main.go http://localhost:8989
```

`fetchall`が永遠に`http://localhost:8989`からのレスポンスを待つ状態になったためサーバを停止

```
ex11 $ sh stop_noressrv.sh
+ curl http://localhost:8989/stop
9.22s        0  http://localhost:8989
9.22s elapsed
2017/05/07 20:04:50 http: Server closed
exit status 1
+ rm results_20170507200440.txt
[1]+  Done                    sh fetchall_nores.sh
```
