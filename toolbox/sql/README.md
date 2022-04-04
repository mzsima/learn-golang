# メモ

docker内のDBに登録したローソク足をSelectする

## 起動
```
go run .
```

## API
BTCのローソク足使いたいから、GMOコインの public api を使って取得　
https://api.coin.z.com/docs/#klines

```
[{"openTime":"1618606800000","open":"6758448","high":"6764919","low":"6758448","close":"6764919","volume":"0.224"},{"openTime":"1618606860000","open":"6764919","high":"6764940","low":"6758680","close":"6764900","volume":"0.2674"},{"openTime":"1618606920000","open":"6764940","high":"6764940","low":"6764940","close":"6764940","volume":"0.0002"},{"openTime":"1618606980000","open":"6764920","high":"6764920","low":"6764920","close":"6764920","volume":"0.04"},{"openTime":"1618607040000","open":"6764920","high":"6764920","low":"6764920","close":"6764920","volume":"0.002"},{"openTime":"1618607100000","open":"6764920","high":"6769879","low":"6764920","close":"6769879","volume":"0.12"},{"openTime":"1618607160000","open":"6760961","high":"6777656","low":"6760961","close":"6777656","volume":"1.342"},{"openTime":"1618607220000","open":"6761441","high":"6772754","low":"6761441","close":"6772299","volume":"0.0626"},{"openTime":"1618607280000","open":"6772257","high":"6772258","low":"6772257","close":"6772258","volume":"0.008"},{"openTime":"1618607340000","open":"6765244","high":"6765244","low":"6765244","close":"6765244","volume":"0.018"},{"openTime":"1618607400000","open":"6765640","high":"6772259","low":"6765640","close":"6765661","volume":"0.0296"},{"openTime":"1618607640000","open":"6772754","high":"6772754","low":"6772754","close":"6772754","volume":"0.0008"},{"openTime":"1618607700000","open":"6772680","high":"6772680","low":"6772680","close":"6772680","volume":"0.0002"},{"openTime":"1618607820000","open":"6766576","high":"6771560","low":"6765009","close":"6765009","volume":"0.1646"},{"openTime":"1618607880000","open":"6765009","high":"6769140","low":"6761460","close":"6761460","volume":"0.044"},{"openTime":"1618607940000","open":"6767850","high":"6767850","low":"6767518","close":"6767518","volume":"0.022"},{"openTime":"1618608000000","open":"6766320","high":"6766320","low":"6766300","close":"6766300","volume":"0.004"}]
```


## SQL作るスクリプト
```
records = [{"openTime":"1618606800000","open":"6758448","high":"6764919","low":"6758448","close":"6764919","volume":"0.224"}, ...]
query = ""
for (let i=0; i<records.length; i++) {
    let {open, high, low, close, volume} = records[i] 
    query += `INSERT INTO candlestick (id, open, high, low, close, volume) VALUES (${i}, ${open}, ${high}, ${low}, ${close}, ${volume}); \n`
}
console.log(query)
```
