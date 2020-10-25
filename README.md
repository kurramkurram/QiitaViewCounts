# QiitaViewCounts
QiitaのLGTMとView数を取得

# 利用方法
1. 以下を実行
```
https://github.com/kurramkurram/QiitaViewCounts.git
```
2. Qiitaの[ユーザー管理画面](https://qiita.com/settings/applications)からトークンを取得

3. 2で取得したトークンを記載した`token.txt`を作成し、main.goと同じディレクトリに配置

4. 実行
`main.go`のディレクトリに移動し、以下を実行
```
go run main.go
```
5. 結果を確認
`main.go`と同じディレクトリにresult.csvが生成できていることを確認