# Go Plurk Robot
[![Build Status](https://secure.travis-ci.org/elct9620/go-plurk-robot.png?branch=master)](https://travis-ci.org/elct9620/go-plurk-robot)
[![Coverage Status](https://coveralls.io/repos/elct9620/go-plurk-robot/badge.svg?branch=master&service=github)](https://coveralls.io/github/elct9620/go-plurk-robot?branch=master)

[![wercker status](https://app.wercker.com/status/f6b3c29f0fdfb254d486cd8667eefc6f/m "wercker status")](https://app.wercker.com/project/bykey/f6b3c29f0fdfb254d486cd8667eefc6f)

不到半個月就要進去國軍 Online 為了安撫寂寞的網友只好上了！

## 使用

這個機器人已經設計為可以自由 Deploy 的狀態，請依照下列指示進行設定。

### Plurk API

請到 [Plurk APP](http://www.plurk.com/PlurkApp) 頁面註冊新的應用服務，並且利用測試工具取得 Client 的 Token 與 Secret 在下一步驟使用。

### 環境變數

* `PLURK_APP_KEY` - Plurk APP 的 Key
* `PLURK_APP_SECRET` - Plurk APP 的 Secret
* `PLURK_OAUTH_TOKEN` - 機器人要登入的使用者的 OAuth Token
* `PLURK_OAUTH_SECRET` - 機器人要登入的使用者的 OAuth Secret
* `PLURK_ROBOT_NAMR` - 機器人的顯示名稱（目前無實質用途）

完成第一步後可以取得上述四項數值，請在運行的主機上設定。

若在 Heroku 上運行可用 `heroku config:set PLURK_APP_KEY=xxx` 的方式設定。

* `SECRET_KEY` - 這是用於加密 Session 的 Key 請生成一組亂入字串設定進去
* `MONGODB_URL` - MongoDB 的伺服器位置，在 Heroku 上可直接設定為 MongoLab 等服務的環境變數名稱

### 產生使用者

透過指令 `go-plurk-robot useradd [帳號] [密碼]` 可以在資料庫產生一個可管理系統的使用者。

### 啟動

* 指令 `go-plurk-robot server` 可以開啟網頁管理介面，編輯機器人的任務。
* 指令 `go-plurk-robot robot` 可以啟動機器人，運行自動任務或者對話。

## 限制

因為開發時間有限，所以目前僅值做完畢 CronJob 功能，可讓機器人定時運行指定任務。
原本應要使用 Redis 對機器人的更動做自動重整，但因時間的關係在編輯後需要重啟機器人。

## 未來計劃

兵役結束後（2016 年中）會重新設計這套系統，分為 Robot, Adapter, Admin UI 三個專案，並且增加相容性和可擴充性。
若有興趣或者任何建議，都可以在 Issue 頁面開設，我會在有空的時候跟各位進行討論。


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/elct9620/go-plurk-robot/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

