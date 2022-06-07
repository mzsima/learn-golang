# Google の ID Tokens を表示する

## how to run

envに id/secret を設定してサーバーを起動

```sh
GOOGLE_OAUTH2_CLIENT_ID=XXXXXXXXX.apps.googleusercontent.com GOOGLE_OAUTH2_CLIENT_SECRET=XXXXX go run main.go 
```

`http://127.0.0.1:5556/` で起動するので、ブラウザで開く。

## GOOGLE のセットアップ

Google認証した結果のtokenを見たいのでここは頑張ってセットアップ。

- https://developers.google.com/identity/protocols/oauth2/openid-connect#appsetup

## 参考 ネットで見たサイト
- https://firebase.google.com/docs/auth/admin/verify-id-tokens
- https://github.com/coreos/go-oidc
- https://zenn.dev/satoken/articles/oauth-funiki
- https://zenn.dev/uma002/articles/51b80fb2b7b108