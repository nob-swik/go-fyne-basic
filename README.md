# Go と Fyne のインストール備忘録(Ubuntu)

## Go

### インストール

```
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```

### 確認

```
go version
```

### in VSCode

- Go の拡張機能をインストール

- インストールが完了したら、Ctrl + Shift + P でコマンドパレットを開いて、「GO: Install/Update tools」を入力してEnter

- 出てきたメニューの全てを選択しOK

## Fyne

### 準備(必要パッケージインストール)

```
sudo apt-get install gcc libgl1-mesa-dev xorg-dev
```

- 適当なフォルダを作成してその中に移動

### module 初期化

```
go mod init MODULE_NAME
```

### インストール

```
go get fyne.io/fyne/v2@latest
go install fyne.io/fyne/v2/cmd/fyne@latest
```

ここで場合によっては再起動や以下のコマンドが必要(？)

```
go mod tidy
```

### 確認

```
go run fyne.io/fyne/v2/cmd/fyne_demo@latest
```

※初回は時間がかかる
