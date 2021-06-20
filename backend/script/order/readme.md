# オーダー方法について

以下４つの方法でオーダー可能です。
- mobile
- pc
- delivery
- register


# mobile
mobile.restから任意のメニューをpostリクエストします。   
※rest Client拡張機能が必要です。

<br><br>

# pc
pc.restから任意のメニューをpostリクエストします。   
※rest Client拡張機能が必要です。

<br><br>

# delivery
delivery.shをcliから実行します。   
※処理実体はdelivery/delivery.goを参照

### grpcurl
コマンドラインでの確認は下記が必要です。
https://github.com/fullstorydev/grpcurl


### grpcインターフェイス確認コマンド   
```
grpcurl -plaintext localhost:3456 list
```


### grpc送信コマンド   
```
grpcurl -plaintext -d @ localhost:3456 delivery.DeliveryService/DeliveryRPC <<EOM
{
    "Order": {
        "Hamburgers": [
            {
                "Top": 1 ,
                "Middle": 0,
                "Bottom":1 ,
                "Beef":1 ,
                "Chicken": 0,
                "Fish": 0,
                "Lettuce": 1,
                "Tomato": 1,
                "Onion": 0,
                "Cheese": 1,
                "Pickles": 1
            }
        ]
    }
}
EOM
```

<br><br>
# register
registerディレクトリ内にjsonファイルを配置。   
order.exampleをコピーして拡張子をjsonに変更。   
<br><br>


