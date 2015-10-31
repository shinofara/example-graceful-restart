Graceful RestartとDeployについて検証してみた
=====================

# Example

このサンプルでは、３つのターミナルを利用して検証しています。

| ID | 用途                 |
|----|----------------------|
| A  | Run a go application |
| B  | http clilent         |
| C  | process killer       | 

#### A term

Go Serverを起動

```
$ go run main.go
2015/10/31 13:46:26 Serving [::]:8080 with pid 65353
```

pid 99999が実行時のpidです

## Graceful ...

通常の停止・再起動だと、処理が走っていた場合でも、強制終了してしまうが、
Graceful なだと、処理を捌ききってから停止・再起動する。

### Graceful Restart

#### B term

立ち上げたhttpサーバに対してリクエストを発行
10秒後レスポンスが返ってきます。

```
$ curl http://localhost:8080/
```

#### C term

10秒間処理を実行してる間に、`USR2` を使ってプロセスを`kill` します。

```
$ sudo kill -USR2 65353
```

この時点ではまだプロセスは生きています。

#### A term

サーバ側では、pid　65353から65387に変更して再起動するよとメッセージが出力されます。

```
2015/10/31 13:47:29 Graceful handoff of [::]:8080 with new pid 65387 and old pid 65353
```

#### B term

curl結果が返ってきました。
最初に立ち上げた時のpidが返ってきます。

```
pid is 65353
```

#### A term

古いpidは終了しました

```
2015/10/31 13:47:29 Exiting pid 65353.
```

#### B term

```
$ curl http://localhost:8080/
pid is 65387
```

新しくなりました

### Graceful Stop

終了時も同じなので、詳細は省きます

```
$ sudo kill -TERM 65353
```
