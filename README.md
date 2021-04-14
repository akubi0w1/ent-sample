# やったこと

code generatorのDL。entも一緒に落ちてきます。

```
$ go get entgo.io/ent/cmd/ent
```

コードを生成してみる。

```
$ ent init User
```

targetを変えてみよう

```
$ ent init --target mysql/ent/schema User
```

generaterが作られなかったwので、これはw

スキーマを定義して。。...

generate

```
$ go generate ./ent
```

`./ent`の部分は、entディレクトリ以下にある、generate.goのあるディレクトリを刺す

基本編集するのはschemaディレクトリのみ?

main.goにてマイグレーションしてます

memo: schemaのIDは勝手にbigintでつくらしい


## reference

edge toだけでやったらどんなテーブルができる？
-> toは一応つくっぽい
Fromが必要なのは、postsからauthorを編集するなどの関数が作られない？

n:nをつくったらどうなる?
-> users_nicesってテーブルができるっぽい。この際、中間テーブルにidはなく、user_id, post_idのみが存在する。複合キーになっているのか...?
-> 複合キーになっているっぽおおい。

to, fromどっちに外部キーがつくのか？
-> fromにつくっぽい

一つのノードに同じ名前のedgeを持たせることはできないらしい

fromの決め方。
1:n -> 1=to, n=from
n:n -> どちらでも構わない
1:1 -> 外部キーを持たせたいnodeをfromにする