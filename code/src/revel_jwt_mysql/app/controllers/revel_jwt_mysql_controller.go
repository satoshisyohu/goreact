
package controllers

import (
    "time"
    "strconv"
    "fmt"
    "github.com/revel/revel"
    "revel_jwt_mysql/app/models"
    "path/filepath"
    "io/ioutil"
    "crypto/rsa"
    "github.com/dgrijalva/jwt-go"
)

// モデルを設定
// revelでURLパラメータ処理が実行可能
// gorpでDBの処理

type Revel_JWT_MySQL struct {
    *revel.Controller
    GorpController
}

// json形式で結果を返すための設定

type Response struct {
    Results interface{} `json:"results"`
}

// modelにデータを保持するための設定
var revel_jwt_mysql = []*models.Revel_JWT_MySQL{}

// URLで受け取るパラメータ
var keys = []string{"Title", "Picture", "Picture_Link", "Caption", "All", "AllDiary", "Date"}

// URLから受け取ったパラメータを最初に処理する部分
func (c Revel_JWT_MySQL) List() revel.Result {
    var res models.Revel_JWT_MySQL
    // modelにurlのパラメータからデータを設定
    c.Params.Bind(&res.Revel_JWT_MySQL_Id, "Id")
    c.Params.Bind(&res.Date, "Date")
    c.Params.Bind(&res.Title, "Title")
    c.Params.Bind(&res.Picture, "Picture")
    c.Params.Bind(&res.Picture_Link, "Picture_Link")
    revel_jwt_mysql = c.getReveJWTMysql()

    // 取得したの情報から日付とIDが一致する情報に対して、JWTトークン認証を行なう
    for _, r := range revel_jwt_mysql {
        if (r.Date == res.Date && r.Revel_JWT_MySQL_Id == res.Revel_JWT_MySQL_Id) {
            return c.checkJWTToken(res, "GET")
        }
    }
    revel.ERROR.Printf("Could not find diary")
    return c.NotFound("Could not find diary")
}

// の全情報を取得。SQL部分を変更することで取得する情報のサイズや条件に制限を加えることが可能
// Args:無し
// return: の全情報をmodelで返却

func (c Revel_JWT_MySQL) getReveJWTMysql() []*models.Revel_JWT_MySQL {
    var revel_jwt_mysqls []*models.Revel_JWT_MySQL
    revel_jwt_mysqls = loadReveJWTMysqls(c.Txn.Select(models.Revel_JWT_MySQL{}, `select * from Revel_JWT_MySQL`))

    return revel_jwt_mysqls
}

// SQLで取得したの情報をモデルに保存して全体を取得。
// Args:
//     results; SQLで取得された結果
// return:
//     []*models.Revel_JWT_MySQL: モデルの配列を返却

func loadReveJWTMysqls(results []interface{}, err error) []*models.Revel_JWT_MySQL {
    if err != nil {
        panic(err)
    }
    var revel_jwt_mysqls []*models.Revel_JWT_MySQL
    for _, r := range results {
        revel_jwt_mysqls = append(revel_jwt_mysqls, r.(*models.Revel_JWT_MySQL))
    }
    return revel_jwt_mysqls
}

// JWT トークンの認証処理
// Args:
//     res: URLパラメータに保存されている各種パラメータ
//     choose_parameter: 取得、登録、更新、削除で処理が異なる
// 正常終了:
//     return: choose_parameterによって処理が異なる
// 異常終了(404):
//     return: JWTトークンが異なる
//     return: JWTトークンが設定されていない
//     return: が存在しない

func (c Revel_JWT_MySQL) checkJWTToken(res models.Revel_JWT_MySQL, choose_parameter string) revel.Result {
    var no_key  bool
    no_key = true
    var check_token bool

    // 情報を取得
    revel_jwt_mysql = c.getReveJWTMysql()
    // Debug
    revel.TRACE.Printf("Check initial regist DB Value")
    for _, tmp := range revel_jwt_mysql {
        revel.TRACE.Printf("%s", *tmp)
    }
    // Debug
    revel.TRACE.Printf("Check Paramater Query")
    revel.TRACE.Printf("%s", c.Params.Query)

    // JWTトークン認証及び各メソッド処理への振り分け
    if val, ok := c.Params.Query["JWTtoken"]; ok {
        check_token = CheckJWTHandler(c.Params.Query["JWTtoken"][0])
        if check_token {
            switch choose_parameter{
                case "GET":
                    for _, k := range keys {
                        if _, ok := c.Params.Query[k]; ok {
                            revel.TRACE.Printf("Param OK")
                            no_key = false
                            return c.returnParameter(k, res.Revel_JWT_MySQL_Id)
                        }
                    }
                    if no_key {
                        revel.ERROR.Printf("No settings %#v", val)
                        return c.Forbidden("No setting.")
                    }
                case "GET_ALL":
                    length, err := strconv.Atoi(c.Params.Query["Length"][0])
                    if err != nil {
                        revel.ERROR.Printf("No settings %#v", length)
                        return c.Forbidden("No setting length.")
                    }else{
                        return c.GetRobohonAllMethod(length)
                    }
                case "GET_JWT":
                    limit, err := strconv.Atoi(c.Params.Query["Limit"][0])
                    if err != nil {
                        revel.ERROR.Printf("No settings %#v", limit)
                        return c.Forbidden("No setting length.")
                    }else{
                        return c.GetJWTMethod(limit)
                    }
                case "POST":
                    return c.RegistRobohonMethod(res)
                case "DELETE":
                    return c.DeleteRobohonMethod(res.Revel_JWT_MySQL_Id)
                case "UPDATE" :
                    for _, k := range keys {
                        if _, ok := c.Params.Query[k]; ok {
                            revel.TRACE.Printf("Param OK")
                            no_key = false
                            return c.UpdateRobohonMethod(k, c.Params.Query[k][0], res.Revel_JWT_MySQL_Id, res.Date)
                        }
                    }
                    if no_key {
                        revel.ERROR.Printf("No settings %#v", val)
                        return c.Forbidden("No setting.")
                    }
            }
        }else {
            revel.ERROR.Printf("Failed JWTtoken not correct.")
            return c.Forbidden("Failed JWTtoken not correct.")
        }
    }else {
        revel.ERROR.Printf("JWTtoken no setting. The Title is %#v", val)
        return c.Forbidden("JWTtoken no setting.")
    }
    return c.NotFound("Could not find diary")
}

// JWT トークンの認証処理
// Args:
//     jwt_token: jwtトークン設定
// Return:
//     true: JWTトークンの鍵認証が通る
//     false: JWTトークンが異なる
//     false: JWTトークンが設定されていない

func CheckJWTHandler(jwt_token string) (bool){
    tokenString := jwt_token
    if tokenString == "" {
        return false
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return LookupPublicKey()
    })

    if err != nil || !token.Valid {
        revel.ERROR.Printf("JWT token invalid Failed")
        return false
    }
    return true
}

// 公開鍵認証処理
// Return:
//     *rsa.PublicKey: 公開鍵情報

func LookupPublicKey() (*rsa.PublicKey, error) {
    key := read_key(true)
    parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
    return parsedKey, err
}

// 鍵読み込み処理
// Args:
//     pub_flag: 公開鍵か秘密鍵か読み込むのに設定
// Return:
//     key: 読み込んだ鍵の結果をバイト形式で返す

func read_key(pub_flag bool) ([]byte) {
    prev, err := filepath.Abs(".")
    if err != nil {
        defer revel.ERROR.Printf("read public key Error")
    }
    // Debug
    revel.TRACE.Printf("check rsa directory")
    revel.TRACE.Printf("%s", prev)
    if pub_flag {
        key, _ := ioutil.ReadFile(prev + "/demo.rsa.pub")
        return key
    }else{
        key, _ := ioutil.ReadFile(prev + "/demo.rsa")
        return key
    }
}
// 返却情報処理(Getメソッド)
// Args:
//     key: 何を取得するかを選択
//     revel_jwt_mysql_id: 取得するid
// Return:
//     revel.Result: 結果を返す

func (c Revel_JWT_MySQL) returnParameter (key string, revel_jwt_mysql_id int) revel.Result {

    var result *models.Revel_JWT_MySQL
    revel_jwt_mysqls, err := c.Txn.Select(models.Revel_JWT_MySQL{},
        `select * from Revel_JWT_MySQL where Revel_JWT_MySQL_Id = ?`,
        revel_jwt_mysql_id)
    if err != nil {
        panic(err)
        return c.Forbidden("SQL Select Error")
    }
    if len(revel_jwt_mysql) == 0 {
        return c.NotFound("Could not find diary")
    }
    result = revel_jwt_mysqls[0].(*models.Revel_JWT_MySQL)

    switch key {
    case "Title":
        return c.RenderJson(result.Title)
    case "Picture":
        return c.RenderJson(result.Picture)
    case "Picture_Link":
        return c.RenderJson(result.Picture_Link)
    case "All":
        return c.RenderJson(result)
    default:
        return c.Forbidden("Failed no setting Parameter.")
    }
}

// 登録処理(Registメソッド)
// URLパラメータから必要な情報を取得し登録処理を行なう
// Return:
//     JWTのチェック処理にPOSTメソッドともに送る

func (c Revel_JWT_MySQL) RegistRobohon() revel.Result {
    var res models.Revel_JWT_MySQL
    c.Params.Bind(&res.Revel_JWT_MySQL_Id, "Id")
    c.Params.Bind(&res.Title, "Title")
    c.Params.Bind(&res.Date, "Date")
    c.Params.Bind(&res.Picture, "Picture")
    c.Params.Bind(&res.Picture_Link, "Picture_Link")
    return c.checkJWTToken(res, "POST")
}

// 情報登録処理(registメソッド)
// Args:
//     res: 情報を登録
// Return:
//     revel.Result: 登録した結果を返す

func (c Revel_JWT_MySQL) RegistRobohonMethod(res models.Revel_JWT_MySQL) revel.Result {
    res.Validate(c.Validation)

    // Handle errors
    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Forbidden("URL parameter Validate error")
    }

    t_value := time.Now().Format("2006-01-02")

    // Revel_JWT_MySQL_Idは主キーに指定しているため、自動で設定されるがモデルの構造上設定が必要
    robohon := models.Revel_JWT_MySQL{res.Revel_JWT_MySQL_Id, t_value, res.Title, res.Picture,
                                                        res.Picture_Link}
    // DB にデータを登録
    if err := c.Txn.Insert(&robohon); err != nil {
        panic(err)
        return c.Forbidden("SQL Regist Error")
    }

    r := Response{robohon}
    return c.RenderJson(r)
}

// 情報削除処理(Deleteメソッド)
// URLパラメータからidを取得し削除処理を行なう
// Return:
//     JWTのチェック処理にDELETEメソッドともに送る

func (c Revel_JWT_MySQL) DeleteRobohon() revel.Result {
    var res models.Revel_JWT_MySQL
    c.Params.Bind(&res.Revel_JWT_MySQL_Id, "Id")
    return c.checkJWTToken(res, "DELETE")
}

// 情報削除処理(deleteメソッド)
// Args:
//     robohonId: 削除したいid
// Return:
//     c.RenderJson: 削除したidを表示

func (c Revel_JWT_MySQL) DeleteRobohonMethod(robohonId int) revel.Result {

    _, err := c.Txn.Delete(&models.Revel_JWT_MySQL{Revel_JWT_MySQL_Id: robohonId})
    if err != nil {
        panic(err)
        return c.Forbidden("SQL Deelete Error")
    }

    show_char := "Delete " + string(robohonId)
    r := Response{show_char}
    return c.RenderJson(r)
}


// 情報更新処理(POSTメソッド)
// Return:
//     更新したidを表示

func (c Revel_JWT_MySQL) UpdateRobohon() revel.Result {
    var res models.Revel_JWT_MySQL
    c.Params.Bind(&res.Revel_JWT_MySQL_Id, "Id")
    c.Params.Bind(&res.Title, "Title")
    c.Params.Bind(&res.Date, "Date")
    c.Params.Bind(&res.Picture, "Picture")
    c.Params.Bind(&res.Picture_Link, "Picture_Link")
    return c.checkJWTToken(res, "UPDATE")
}

// 情報更新処理(POSTメソッド)
// Args:
//     key: 更新したいパラメータ
//     value: 更新する値
//     robohonId: 更新したいid
//     date: 更新したい日付
// Return:
//     c.RenderJson: 更新したの情報を表示

func (c Revel_JWT_MySQL) UpdateRobohonMethod(key string, value string, robohonId int, date string) revel.Result {

    switch key {
    case "Title":
        _, err := c.Txn.Exec("update Revel_JWT_MySQL set Title = ? where Revel_JWT_MySQL_Id = ? and Date = ?",
            value, robohonId, date)
        if err != nil {
            panic(err)
        }
    case "Picture":
        _, err := c.Txn.Exec("update Revel_JWT_MySQL set Picture = ? where Revel_JWT_MySQL_Id = ? and Date = ?",
            value, robohonId, date)
        if err != nil {
            panic(err)
        }
    case "Picture_Link":
        _, err := c.Txn.Exec("update Revel_JWT_MySQL set Picture_Link = ? where Revel_JWT_MySQL_Id = ? and Date = ?",
            value, robohonId, date)
        if err != nil {
            panic(err)
        }
    case "Caption":
        _, err := c.Txn.Exec("update Revel_JWT_MySQL set Caption = ? where Revel_JWT_MySQL_Id = ? and Date = ?",
            value, robohonId, date)
        if err != nil {
            panic(err)
        }
    default:
        return c.Forbidden("Failed no setting Parameter.")
    }

    show_char := "UpdateRobohon " + value
    r := Response{show_char}
    return c.RenderJson(r)
}

// id取得処理
// idの一覧を取得する。デフォルトの長さは10
// Args:
//     length: idの取得する数(今後増えるので)
// Return:
//     一覧を表示

func (c Revel_JWT_MySQL) ListAll() revel.Result {
    var res models.Revel_JWT_MySQL
    return c.checkJWTToken(res, "GET_ALL")
}

// 一覧情報取得処理(Getメソッド)
// Args:
//     length: の取得する数(今後増えるので)
// Return:
//     revel.Result: 一覧を返す

func (c Revel_JWT_MySQL) GetRobohonAllMethod (length int) revel.Result {

    result, err := c.Txn.Select(models.Revel_JWT_MySQL{},
        `select * from Revel_JWT_MySQL limit ?`,
        length)
    if err != nil {
        panic(err)
        return c.Forbidden("SQL Select Error")
    }
    if len(result) == 0 {
        return c.NotFound("Could not find diary")
    }
    revel.TRACE.Printf("%s", result)

    return c.RenderJson(result)
}

// JWTトークン発行処理
// JWTトークンの発行処理
// "iss": "https://jwt-idp.example.com", apiのドメインを設定予定
// "sub": "mailto:mike@example.com", apiのサブドメインを設定予定
// "nbf": 1454503213, JWTが有効になる日時を示す
// "exp": 1490667232, JWT発行者の識別子
// "iat": 1454503213, JWTを発行した時刻を示す
// "jti": "id123456", OPTIONALパラメータ JWT発行者の識別子
// "typ": "https://example.com/register" OPTIONALパラメータ JWTであることを明示するため
// Args:
//     nbf: JWTが有効になる日時を示す
//     exp: JWT発行者の識別子
//     iat: JWTを発行した時刻を示す
// Return:
//     JWTトークンを表示

func (c Revel_JWT_MySQL) GetJWT() revel.Result {
    var res models.Revel_JWT_MySQL
    return c.checkJWTToken(res, "GET_JWT")
}

// JWTトークン発行処理(Getメソッド)
// Args:
//     limit: 何日後に無効になるか設定
// Return:
//     revel.Result: JWTトークンを返す

func (c Revel_JWT_MySQL) GetJWTMethod (limit int) revel.Result {
    // 時間情報のフォーマットを設定している
    t_value := time.Now().Format("2006-01-02")
    // Create the token
    token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
        "iss": "https://jwt-idp.example.com",
        "sub": "mailto:mike@example.com",
        "nbf": t_value,
        "exp": time.Now().AddDate(0, 0, limit),
        "iat": t_value,
        "jti": "id123456",
        "type": "JWT",
    })
    // Set some claims
    // Sign and get the complete encoded token as a string
    key := read_key(false)
    tokenString, err := token.SignedString(key)
    // Debug
    revel.TRACE.Printf(tokenString)
    if err != nil{
        return c.Forbidden("JWT token not publish")
    }
    return c.RenderJson(tokenString)
}
