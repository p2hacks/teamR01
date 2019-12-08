# API

## 概要
Golangを書き慣れていないメンバーがいるため,  
**基本は**1つのAPIに1つの処理しか実装されていないAPIを量産する方針で開発を進めることにした.  
また, 作成したAPIはコンテナ化する

## 作成するAPI
### bff API
下記のAPIのエンドポイントをまとめるためのAPI
言語: Java Script

### User API
ユーザー情報を登録し, ユーザーIDを返すAPI
リクエスト時に必要な情報
言語: Golang

### Card-Info API
ユーザーのクレジットカードを登録するAPI
言語: Golang

### レート API
ユーザーが参加している部屋を管理するAPI
言語: Golang

### プレゼントAPI
ユーザーがほしいプレゼントを管理するAPI
言語: Golang

### to be サンタ API
ユーザーに購入させるプレゼントを振り分けるAPI
言語: Golang

### 決済API
決済API
言語: Golang

### Send メッセージAPI
ユーザーがプレゼント送り主に対してメッセージを送るAPI
言語: Golang

### Get メッセージAPI
ユーザーがプレゼント送りてからのメッセージを受け取るAPI
言語: Golang