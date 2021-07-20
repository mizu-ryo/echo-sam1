# git cloneしたら
`docker-compose build`

# 立ち上げ
`docker-compose up`
`https://localhost:8080`へアクセスするとhello worldがでる．

# route
## index
get:`https://localhost:8080/users/index`でuser一覧が得られる
現在はupするたびにシードが作られる状態になっているので，upするたびにサンプルデータが作られるようになってしまっている．

## create
post:`http://localhost:8080/users/new`にリクエストを送るとuserが作成できる．
postmanでbodyのform-dataに以下のようなデータを入れて確認できる．
| key           | value          | description  |
| ------------- | -------------- | ------------ |
| name          | aaa            |              |
| email         | aaa@email.com  |              |


## delete
delete:`http://localhost:8080/users/delete/:id`をするとidのuserを削除する

# dbのデータ削除
`docker-compose down -v`をするとdbのデータを削除する


