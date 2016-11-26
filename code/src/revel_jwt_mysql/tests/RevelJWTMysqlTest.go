package tests

import (
       	"github.com/revel/revel/testing"
       	"strings"
)

type RevelJWTMysqlTest struct {
       	testing.TestSuite
}

func (t *RevelJWTMysqlTest) Before() {
       	println("Set up")
}

//のタイトル取得処理Test

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlIdDateTitle() {
       	t.Get("/revel_jwt_mysql?Id=1&Date=2016-04-02&Title&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y")
       	t.AssertContains("test1_title")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

//の写真取得処理Test

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlIdDatePicture() {
       	t.Get("/revel_jwt_mysql?Id=1&Date=2016-04-02&Picture&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y")
       	t.AssertContains(":9000/public/img/test1.jpg")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

//の写真リンク取得処理Test

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlIdDatePictureLink() {
       	t.Get("/revel_jwt_mysql?Id=1&Date=2016-04-02&Picture_Link&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y")
       	t.AssertContains("test1_picture")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

//の全情報取得処理Test

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlIdAll() {
       	t.Get("/revel_jwt_mysql?Id=1&Date=2016-04-02&All&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y")
	t.AssertContains("test1_title")
	t.AssertContains(":9000/public/img/test1.jpg")
	t.AssertContains("test1_picture")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

//の情報登録処理Test

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlRegist() {
       	r:=strings.NewReader("4")
       	t.Post("/revel_jwt_mysql_regist?Date=2016-03-29&Title=revel_jwt_mysql_title_regist&Picture=revel_jwt_mysql_picture_regist&Picture_Link=revel_jwt_mysql_picture_link_regist&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y", "application/json; charset=utf-8", r)
       	t.AssertContains("revel_jwt_mysql_title_regist")
       	t.AssertContains("revel_jwt_mysql_picture_regist")
       	t.AssertContains("revel_jwt_mysql_picture_link_regist")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

//の情報登録処理Error Test

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlRegistError() {
       	r:=strings.NewReader("4")
       	// Date Nothing
       	t.Post("/revel_jwt_mysql_regist?Title=revel_jwt_mysql_title_regist&Picture=revel_jwt_mysql_picture_regist&Picture_Link=revel_jwt_mysql_picture_link_regist&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y", "application/json; charset=utf-8", r)
       	t.AssertStatus(403)
       	// Title Nothing
       	t.Post("/revel_jwt_mysql_regist?Date=2016-03-29&Picture=revel_jwt_mysql_picture_regist&Picture_Link=revel_jwt_mysql_picture_link_regist&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y", "application/json; charset=utf-8", r)
       	t.AssertStatus(403)
       	// Picture Nothing
       	t.Post("/revel_jwt_mysql_regist?Date=2016-03-29&Title=revel_jwt_mysql_title_regist&Picture_Link=revel_jwt_mysql_picture_link_regist&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y", "application/json; charset=utf-8", r)
       	t.AssertStatus(403)
       	// Picture Link Nothing
       	t.Post("/revel_jwt_mysql_regist?Date=2016-03-29&Title=revel_jwt_mysql_title_regist&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y", "application/json; charset=utf-8", r)
       	t.AssertStatus(403)
}

//情報削除処理Test

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlDelete() {
       	t.Delete("/revel_jwt_mysql_delete?Id=5&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y")
       	t.AssertContains("5")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

//情報更新処理Test

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlUpdate() {
       	r:=strings.NewReader("4")
       	// Update Title
       	t.Post("/revel_jwt_mysql_update?Id=4&Date=2016-03-29&Title=hogeTitle&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y", "application/json; charset=utf-8", r)
       	t.AssertContains("hogeTitle")
       	t.AssertOk()
       	// Update Picture
       	t.Post("/revel_jwt_mysql_update?Id=4&Date=2016-03-29&Picture=HogePicture&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y", "application/json; charset=utf-8", r)
       	t.AssertContains("HogePicture")
       	t.AssertOk()
       	// Update Picture Link
       	t.Post("/revel_jwt_mysql_update?Id=4&Date=2016-03-29&Picture_Link=HogePicture_Link&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y", "application/json; charset=utf-8", r)
       	t.AssertContains("HogePicture_Link")
       	t.AssertOk()
       	t.AssertContentType("application/json; charset=utf-8")
}

///情報見つからないときの処理Test

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlIdNotFound() {
       	t.Get("/revel_jwt_mysql_diary")
       	t.AssertNotFound()
}

//情報のJWTトークンが間違っている時の処理Test

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlApiNotCorrect() {
       	t.Get("/revel_jwt_mysql?Id=3&Date=2016-04-02&Title&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ1NDUwMzIxMywiZXhwIjoxNDkwNjY3MjMyLCJpYXQiOjE0NTQ1MDMyMTMsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Y6Wz8yods7NFxwYaTnL94CnZfJJ84aLkm592wwNWItY-Z07c2y4ZaoG29izwxZcyPwAKXPwkzmkwooDcwGaXClftoRKkQKd4uu3Ax4d2mJhrQNsmVzhYzsXZEWDH9ORqpGcISssJLA8rSO_dZWrMd8GIzm1jDWKv-zC1O380d8a")
       	t.AssertStatus(403)
}

//情報の日本語が正しく表示されるかテスト

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlJapanese() {
       	t.Get("/revel_jwt_mysql?Id=3&Date=2016-04-02&All&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y")
       	t.AssertOk()
       	t.Get("/revel_jwt_mysql?Id=3&Date=2016-04-02&Title&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y")
       	t.AssertOk()
       	t.AssertContains("タイトル題_3")
       	t.Get("/revel_jwt_mysql?Id=3&Date=2016-04-02&Picture_Link&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y")
       	t.AssertOk()
       	t.AssertContains("写真_3")
       	t.AssertContentType("application/json; charset=utf-8")
}

//一覧情報取得テストパラメータなし

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlGetAllNonParamater() {
       	t.Get("/revel_jwt_mysql_all?JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y")
       	t.AssertStatus(500)
}

//一覧情報取得テストパラメータ文字

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlGetAllNotCorrectParamater() {
       	t.Get("/revel_jwt_mysql_all?Length=hoge&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y")
       	t.AssertStatus(403)
}


//JWTトークン取得テスト

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlGetJWT() {
       	t.Get("/revel_jwt_mysql_jwt?Limit=72&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y")
       	t.AssertOk()
}

//JWTトークン取得テストパラメータなし

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlGetJWTNonParamater() {
       	t.Get("/revel_jwt_mysql_jwt?JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y")
       	t.AssertStatus(500)
}

//JWTトークン取得テストパラメータが文字

func (t *RevelJWTMysqlTest) TestRevelJWTMysqlGetJWTNotCorrectParamater() {
       	t.Get("/revel_jwt_mysql_jwt?Limit=hoge&JWTtoken=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.nK6gM51kP4klRVQoJ0w6dQLQcbQ2Jrf0TPYp9hZS_8mNZWQaHD9zq5-mIFam1xwJt_dorgBhZrfQzu7tKKbeyHmfj6TcWsSJ7T2W6mG0uFSHyAMHJRpjhFrwGB6K5dzUOvcYgw1B1L-AD6-37zWt6tXP_9Y8HVy4xL-oCR_979Y")
       	t.AssertStatus(403)
}

func (t *RevelJWTMysqlTest) After() {
       	println("Tear down")
}
