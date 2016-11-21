package tests

import (
       	"github.com/revel/revel/testing"
       	"strings"
)

type RobohonDiaryTest struct {
       	testing.TestSuite
}

func (t *RobohonDiaryTest) Before() {
       	println("Set up")
}

//ロボホンの日記タイトル取得処理Test

func (t *RobohonDiaryTest) TestRobohonDiaryIdDateTitle() {
       	t.Get("/revel_jwt_mysql_diary?Id=1&Date=2016-04-02&Title&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertContains("revel_jwt_mysql_title_1")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

//ロボホンの日記写真取得処理Test

func (t *RobohonDiaryTest) TestRobohonDiaryIdDatePicture() {
       	t.Get("/revel_jwt_mysql_diary?Id=1&Date=2016-04-02&Picture&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertContains("http://ec2-52-7-241-51.compute-1.amazonaws.com:9000/public/img/20091207-1260117798.jpg")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

//ロボホンの日記写真リンク取得処理Test

func (t *RobohonDiaryTest) TestRobohonDiaryIdDatePictureLink() {
       	t.Get("/revel_jwt_mysql_diary?Id=1&Date=2016-04-02&Picture_Link&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertContains("revel_jwt_mysql_picture_link_1")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

//ロボホンのキャプション取得処理Test

func (t *RobohonDiaryTest) TestRobohonDiaryIdDateCaption() {
       	t.Get("/revel_jwt_mysql_diary?Id=1&Date=2016-04-02&Caption&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertContains("revel_jwt_mysql_caption_1")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

//ロボホンの全情報取得処理Test

func (t *RobohonDiaryTest) TestRobohonDiaryIdAll() {
       	t.Get("/revel_jwt_mysql_diary?Id=1&Date=2016-04-02&All&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertContains("revel_jwt_mysql_title_1")
       	t.AssertContains("http://ec2-52-7-241-51.compute-1.amazonaws.com:9000/public/img/20091207-1260117798.jpg")
       	t.AssertContains("revel_jwt_mysql_picture_link_1")
       	t.AssertContains("revel_jwt_mysql_caption_1")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

//ロボホンの情報登録処理Test

func (t *RobohonDiaryTest) TestRobohonRegist() {
       	r:=strings.NewReader("4")
       	t.Post("/revel_jwt_mysql_diary_regist?Date=2016-03-29&Title=revel_jwt_mysql_title_regist&Picture=revel_jwt_mysql_picture_regist&Picture_Link=revel_jwt_mysql_picture_link_regist&Caption=revel_jwt_mysql_caption_regist&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g", "application/json; charset=utf-8", r)
       	t.AssertContains("revel_jwt_mysql_title_regist")
       	t.AssertContains("revel_jwt_mysql_picture_regist")
       	t.AssertContains("revel_jwt_mysql_picture_link_regist")
       	t.AssertContains("revel_jwt_mysql_caption_regist")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

//ロボホンの情報登録処理Error Test

func (t *RobohonDiaryTest) TestRobohonRegistError() {
       	r:=strings.NewReader("4")
       	// Date Nothing
       	t.Post("/revel_jwt_mysql_diary_regist?Title=revel_jwt_mysql_title_regist&Picture=revel_jwt_mysql_picture_regist&Picture_Link=revel_jwt_mysql_picture_link_regist&Caption=revel_jwt_mysql_caption_regist&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g", "application/json; charset=utf-8", r)
       	t.AssertStatus(403)
       	// Title Nothing
       	t.Post("/revel_jwt_mysql_diary_regist?Date=2016-03-29&Picture=revel_jwt_mysql_picture_regist&Picture_Link=revel_jwt_mysql_picture_link_regist&Caption=revel_jwt_mysql_caption_regist&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g", "application/json; charset=utf-8", r)
       	t.AssertStatus(403)
       	// Picture Nothing
       	t.Post("/revel_jwt_mysql_diary_regist?Date=2016-03-29&Title=revel_jwt_mysql_title_regist&Picture_Link=revel_jwt_mysql_picture_link_regist&Caption=revel_jwt_mysql_caption_regist&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g", "application/json; charset=utf-8", r)
       	t.AssertStatus(403)
       	// Picture Link Nothing
       	t.Post("/revel_jwt_mysql_diary_regist?Date=2016-03-29&Title=revel_jwt_mysql_title_regist&Caption=revel_jwt_mysql_caption_regist&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g", "application/json; charset=utf-8", r)
       	t.AssertStatus(403)
       	// Caption Nothing
       	t.Post("/revel_jwt_mysql_diary_regist?Date=2016-03-29&Title=revel_jwt_mysql_title_regist&Picture=revel_jwt_mysql_picture_regist&Picture_Link=revel_jwt_mysql_picture_link_regist&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g", "application/json; charset=utf-8", r)
       	t.AssertStatus(403)
}

//ロボホン日記情報削除処理Test

func (t *RobohonDiaryTest) TestRobohonDelete() {
       	t.Delete("/revel_jwt_mysql_delete?Id=5&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertContains("5")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

//ロボホン日記情報更新処理Test

func (t *RobohonDiaryTest) TestRobohonUpdate() {
       	r:=strings.NewReader("4")
       	// Update Title
       	t.Post("/revel_jwt_mysql_diary_update?Id=4&Date=2016-03-29&Title=hogeTitle&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g", "application/json; charset=utf-8", r)
       	t.AssertContains("hogeTitle")
       	t.AssertOk()
       	// Update Picture
       	t.Post("/revel_jwt_mysql_diary_update?Id=4&Date=2016-03-29&Picture=HogePicture&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g", "application/json; charset=utf-8", r)
       	t.AssertContains("HogePicture")
       	t.AssertOk()
       	// Update Picture Link
       	t.Post("/revel_jwt_mysql_diary_update?Id=4&Date=2016-03-29&Picture_Link=HogePicture_Link&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g", "application/json; charset=utf-8", r)
       	t.AssertContains("HogePicture_Link")
       	t.AssertOk()
       	// Update Caption
       	t.Post("/revel_jwt_mysql_diary_update?Id=4&Date=2016-03-29&Caption=HogeCaption&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g", "application/json; charset=utf-8", r)
       	t.AssertContains("HogeCaption")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

//ロボホン日記情報見つからないときの処理Test

func (t *RobohonDiaryTest) TestRobohonDiaryIdNotFound() {
       	t.Get("/revel_jwt_mysql_diary")
       	t.AssertNotFound()
}

//ロボホン日記情報の登録されている日付がないときの処理Test

func (t *RobohonDiaryTest) TestRobohonDiaryDateNotfound() {
       	t.Get("/revel_jwt_mysql_diary?Id=3&Date=2016-03-26&Title&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertNotFound()
}

//ロボホン日記情報のJWTトークンが間違っている時の処理Test

func (t *RobohonDiaryTest) TestRobohonDiaryApiNotCorrect() {
       	t.Get("/revel_jwt_mysql_diary?Id=3&Date=2016-04-02&Title&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8a")
       	t.AssertStatus(403)
}

//ロボホン日記情報の日本語が正しく表示されるかテスト

func (t *RobohonDiaryTest) TestRobohonDiaryJapanese() {
       	t.Get("/revel_jwt_mysql_diary?Id=3&Date=2016-04-02&All&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertOk()
       	t.Get("/revel_jwt_mysql_diary?Id=3&Date=2016-04-02&Title&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertOk()
       	t.AssertContains("ロボフォンタイトル題_3")
       	t.Get("/revel_jwt_mysql_diary?Id=3&Date=2016-04-02&Picture_Link&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertOk()
       	t.AssertContains("ロボフォン写真リンク_3")
       	t.Get("/revel_jwt_mysql_diary?Id=3&Date=2016-04-02&Caption&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertOk()
       	t.AssertContains("ロボフォン説明_3")
       	t.AssertContentType("application/json; charset=utf-8")
}

//ロボホン日記一覧情報取得テスト

func (t *RobohonDiaryTest) TestRobohonDiaryGetAll() {
       	t.Get("/revel_jwt_mysql_diary_all?Length=2&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertOk()
       	t.AssertContains("[{\"Revel_JWT_MySQL_Id\":1,\"Date\":\"2016-04-02\",\"Title\":\"revel_jwt_mysql_title_1\",\"Picture\":\"http://ec2-52-7-241-51.compute-1.amazonaws.com:9000/public/img/20091207-1260117798.jpg\",\"Picture_Link\":\"revel_jwt_mysql_picture_link_1\",\"Caption\":\"revel_jwt_mysql_caption_1\"},{\"Revel_JWT_MySQL_Id\":2,\"Date\":\"2016-04-02\",\"Title\":\"revel_jwt_mysql_title_2\",\"Picture\":\"http://ec2-52-7-241-51.compute-1.amazonaws.com:9000/public/img/IMG_2875ISUMI_TP_V.jpg\",\"Picture_Link\":\"revel_jwt_mysql_picture_link_2\",\"Caption\":\"revel_jwt_mysql_caption_2\"}]")
    t.Get("/revel_jwt_mysql_diary_all?Length=10&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
    t.AssertOk()
}

//ロボホン日記一覧情報取得テストパラメータなし

func (t *RobohonDiaryTest) TestRobohonDiaryGetAllNonParamater() {
       	t.Get("/revel_jwt_mysql_diary_all?JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertStatus(500)
}

//ロボホン日記一覧情報取得テストパラメータ文字

func (t *RobohonDiaryTest) TestRobohonDiaryGetAllNotCorrectParamater() {
       	t.Get("/revel_jwt_mysql_diary_all?Length=hoge&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertStatus(403)
}


//ロボホンJWTトークン取得テスト

func (t *RobohonDiaryTest) TestRobohonDiaryGetJWT() {
       	t.Get("/revel_jwt_mysql_diary_jwt?Limit=72&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertOk()
}

//JWTトークン取得テストパラメータなし

func (t *RobohonDiaryTest) TestRobohonDiaryGetJWTNonParamater() {
       	t.Get("/revel_jwt_mysql_diary_jwt?JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertStatus(500)
}

//JWTトークン取得テストパラメータが文字

func (t *RobohonDiaryTest) TestRobohonDiaryGetJWTNotCorrectParamater() {
       	t.Get("/revel_jwt_mysql_diary_jwt?Limit=hoge&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8g")
       	t.AssertStatus(403)
}

func (t *RobohonDiaryTest) After() {
       	println("Tear down")
}
