**go_attendance_program**

## 起動
本プログラムの起動方法を記載します

### 1. アプリケーションを起動
```
go run main.go
```
### 2. ターミナル上に結果が返ってくること
```
PS C:\Users\aoshi\Documents\github\delight\go_attendance_program\v1> go run main.go
jsonファイルを保存しました。
保存先: .data/output/student_seat20250322182940.json
```

## パッケージ構成
本プログラムのパッケージ構成を記載します。

```
├── .data
|　　　├── output
|　　　└── student_info.json
├── .vscode
|　　　└── launch.json
├── components
|　　　├── file.go
|　　　└── util.go
├── config
|　　　└── const.go
├── domain
|　　　└── student.go
├── go.mod
└── main.go
```