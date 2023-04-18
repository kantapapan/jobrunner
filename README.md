# JobRunner

JobRunnerは、Go言語を使用してジョブスケジューラとウェブサーバを実装するサンプルプロジェクトです。
このプロジェクトでは、[github.com/bamzi/jobrunner](https://github.com/bamzi/jobrunner)ライブラリを使用してジョブをスケジュールし、[github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)ライブラリを使用してウェブサーバを実装しています。

## コードの説明

コードは以下の機能を持っています:

1. ジョブランナーを開始し、`MyJob`を5秒ごとに実行するようにスケジュールします。
2. Ginウェブフレームワークを使用してウェブサーバを実装し、ポート8080で起動します。
3. ジョブランナーのステータスをJSON形式で返すエンドポイント`/jobrunner/status`を提供します。

## 使い方

1. このリポジトリをクローンまたはダウンロードします。

```bash
git clone https://github.com/your_username/jobrunner.git
```

2. プロジェクトディレクトリに移動します。
```bash
cd jobrunner
```

3. 必要な依存関係をインストールします。
```bash
go get -u github.com/bamzi/jobrunner
go get -u github.com/gin-gonic/gin
```

4. プログラムを実行します。
```bash
go run main.go
```
ウェブサーバがポート8080で起動し、ジョブが5秒ごとに実行されます。ブラウザで http://localhost:8080/jobrunner/status にアクセスして、ジョブランナーのステータスを確認します。


