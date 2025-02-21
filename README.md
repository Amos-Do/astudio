# A Studio
個人工作室預約系統 API

## 計畫開發功能
1. 後台商家登入API (/auth)
2. 商家管理 （/vendor）
3. 商家會員管理
4. 預約管理
5. LINEbot Message 設定管理

## 目前開發功能
1. 後台商家登入API (/auth)
  * 完成商家登入取得授權 token（/auth/login）
  * 完成商家註冊（/auth/signup）
  * 完成商家更新授權 token（/auth/refresh）

## Run the Applications
```bash
# move to workspace
$ cd workspace

# clone into your workspace
$ git clone git@github.com:Amos-Do/astudio.git

# move to project
$ cd astudio

# run 
# http://localhost:5000/swagger/index.html
$ docker-compose up

```

## Run local dev
```bash
# move to project
$ cd astudio

# build postgresql
docker-compose up -d postgres_db

# run local dev
cd server
go run ./cmd/astudio/main.go -env dev
```