# AUTH0のサンプル

ここを参考にした
https://auth0.com/docs/quickstart/backend/golang/01-authorization?download=true


## JWTの取得
Auth0のAPIを作成したあと、こんな感じのRequestをauth0に飛ばすとJWTがGetできる
```
curl --request POST \
  --url https://dev-xxx-xxxxx.us.auth0.com/oauth/token \
  --header 'content-type: application/json' \
  --data '{"client_id":"...","client_secret":"...","audience":"https://...","grant_type":"client_credentials"}'
```

## JWTを検証する
.envに auth0の設定を書くようにしている。
`go run main.go` 
でアプリを起動したら、/api/privateに対してこんな感じのRequestを飛ばす。JWTが検証できるか確認

```
curl --request GET \
  --url http://localhost:3010/api/private \
  --header 'authorization: Bearer eyJhbxxxxxx.eyJpcxxxxxxx.xxxxxxxxxxxx'
```