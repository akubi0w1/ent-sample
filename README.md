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

スキーマを定義して。。...

generate

```
$ go generate ./ent
```

`./ent`の部分は、entディレクトリ以下にある、generate.goのあるディレクトリを刺す

基本編集するのはschemaディレクトリのみ?

main.goにてマイグレーションしてます

memo: schemaのIDは勝手にbigintでつくらしい