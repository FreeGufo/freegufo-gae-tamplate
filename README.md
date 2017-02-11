# 概要

フリーグーフォ合同会社で使っているGoogle App Engine for Go (GAE/Go)の新規プロジェクト作成用のテンプレートです。

ドキュメントは作成中です。

## 特徴

- vendoringが可能な構成
- vendorには、glideを利用
- cssは、Bootstrap4をscssで利用


この構成は、以下のBlog記事をベースとしています。  
https://www.freegufo.com/tech-blog/2016/gae-vendoring


構成は以下のようになっています。

```
freegufo-gae-template (GOPATH)
└── src
    ├── freegufo-gae-tamplate
    │   └── main.go
    │
    ├── static
    │   ├── css
    │   ├── font
    │   ├── html
    │   ├── image
    │   ├── js
    │   └── scss
    │
    ├── vendor (glideで取得）
    │   ├── google.golang.com
    │   …
    │
    ├── app.yaml
    ├── glide.lock
    └── glide.yaml
```

## Golangの設定


### glideの設定

公式ドキュメントをご覧下さい。  
https://glide.sh/

ぐぐると多くの情報が出てきます。

インストール後は、以下のコマンドを使えば大体行けます。

```
glide init
glide up
```

### GOPATHの設定

プロジェクトルートをGOPATHに設定

※ 以前は、vendoringが使えていなかったので、 `go get` するためのGOPATHをプロジェクトの外のディレクトリに設定していたため、プロジェクトルートと含めて複数のPathをGOPATHに設定していました。


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

#### FontAwesame

Webフォント用のライブラリ。  
bootstrap4 alpha.6では、Web Font が、まだTODOとなっており出来ていないようでしたので、こちらを利用しています。

- scssで利用


#### scssの構成

`bootstrap` と `font-awesome` 用のディレクトリを作成。  
変数の上書き等に使うため、 `_freegufo_custum_after.scss` `_freegufo_custum_before.scss` を作成して、`freegufo.scss` 内で以下の様に記載してコンパイルしています。

```
@import "freegufo_custum_before";
@import "bootstrap/bootstrap";
@import "font-awesome/font-awesome";
@import "freegufo_custum_after";
```

#### コンパイル用コマンド

以下は、コンパイルのコマンドのメモです。

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