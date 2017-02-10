# 概要

フリーグーフォ合同会社で使っているGoogle App Engine for Go (GAE/Go)の新規プロジェクト作成用のテンプレートです。

ドキュメントは作成中です。

## 特徴

- vendoringが可能な構成
- vendorには、glideを利用


## Golangの設定

## glideの設定

TODO


### GOPATHの設定

プロジェクトルートをGOPATHに設定

Mac
```
cd [project_root]
export GOPATH=`pwd`
```

PowerShell(Windows)
```
cd [project_root]
$ENV:GOPATH = $(pwd).Path
```


## Client Library

### css


#### BootStrap4
 
- alpha.6 を利用
- scssを利用してカスタマイズ

#### Fontawesam

- scssで利用


#### コンパイルコマンド


node-sassのインストール
```
npm install -g node-sass
```

コンパイル
```
cd [project root]/src/static
node-sass -w --source-map true --output-style compressed scss/freegufo.scss  css/freegufo.min.css
```

minifyなし
```
node-sass -w --source-map true scss/freegufo.scss  css/freegufo.css

```