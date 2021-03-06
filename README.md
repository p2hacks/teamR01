# P2HACKS令和元年 アピールシート

## プロダクト

### コンセプト
夢のない現実にサンタを実感されてくれるアプリ

### 説明
アプリを通してユーザーが他のユーザーにプレゼントを送り合うアプリ（サービス）.  
サービスの価値について考えてみました.  
クリスマスに夢がないユーザー（大学生）は、  
サンタからプレゼントが欲しいが、  
サンタはいないし, プレゼントももらえないので、  
アプリを通してサンタからプレゼントを受け取れることにサービスとしての価値があると考えた.  

### 推しポイント
ただプレゼントを送り合うだけでなく, プレゼントを送った相手から  
ありがとうのメッセージが送られることでほっこりできる仕様になっています.

## 開発技術

### 利用したプログラミング言語
- Swift
- Go
- Java script

### 利用したフレームワーク・ライブラリ
- ### iOS
- apollo-ios
- ### サーバ
- gin
- echo
- gorm
- JWT
- apollo
- ### インフラ
- Docker
- k8s

### その他開発に使用したツール
- MySQL

## 役割分担
- 山田咲太朗(サーバ & インフラ)
- 田澤卓也(サーバ)
- 野間直生(クライアント/iOS) 
- 伊藤晋梧(クライアント/iOS)

## チーム目標

### 目標
- チーム全員が描いているプロダクトを作りきる.  

### 達成度
- 75%

### 達成度の根拠
- サーバーサイドは技術的挑戦, 進捗ともに評価できるものができたのから.
- クライアントは技術選定の時点で失敗し, 後半に大きなミスが目立ったが動くものを作ることができた

## 個人目標と達成度

### 山田咲太朗
- 目標：kubernetesを理解して安定稼働するサービスを構築する. & クライアントの実装に優しいアーキテクチャを考える.
- 達成度： 80%
- 根拠：k8sの実装は終わらなかったが, やりたいことがしっかりとできたのでこの評価にした.

### 田澤卓也
- 目標：golangでechoを用いて動くサーバーアプリケーションを作る。データベースを操作すること。
- 達成度：80%
- 根拠:echoとgormを用いてサーバーアプリケーションを作ることができた。値を登録する取り出す以外のデータベース操作を行うAPIは担当できなかった。

### 野間直生
- 目標：apollo-iosを理解して、自由に使えることができるようになる。
- 達成度：30%
- 根拠：動くものを作ることを優先したので、apollo-iosに時間を使うことがあまりできませんでした。ただ、目標には書いていなかったモックアップの作成や、GitHubの利用などチーム開発をするにおいて必要なことを学ぶことができました。次もし機会があれば興味のある技術に触ることができるように余裕をもって大会に取り組みたいです。

### 伊藤晋梧
- 目標：iOSのアプリケーションの動くものを作る。チームのメンバーに迷惑をかけないよう頑張る.  
- 達成度：60%
- 根拠：自分はSwiftUIで画面作りをしていたが、直前でStoryboardとうまく噛み合わなくてSwiftUIで作っていた分を全部消し、Storyboardで作り直すことになった。なので、直接自分が関わって動くものは作れなかった。しかし、GitHubでしっかりコード管理ができ、メンバーと話し合いながら協力して開発できた。
