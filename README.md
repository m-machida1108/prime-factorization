# prime-factorization
`1000000000`までの数値を素因数分解します。
以下はdockerコンテナ上で実行する手順です。

1. コンテナ立ち上げ
  - `$ make up`
1. 実行
  - `$ make exec target=<素因数分解したい数値>`

コンテナを落とす場合は`$ make down`を実行してください。
