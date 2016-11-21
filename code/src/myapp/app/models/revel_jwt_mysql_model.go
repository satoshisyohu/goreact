package models

import ("net/url"
       "mime/multipart"
       	   "github.com/revel/revel"
)

// に必要な項目設定：DBと連動している
type Revel_JWT_MySQL struct {
       	Revel_JWT_MySQL_Id     int
       	Date          string
       	Title         string
       	Picture       string
       	Picture_Link  string
       	Caption       string
}

// Url パラメータを処理するのに必要
type Params struct {
       	url.Values
       	Files map[string][]*multipart.FileHeader
}

// の情報を登録する際にチェックする機能
func (revel_jwt_mysql Revel_JWT_MySQL) Validate(v *revel.Validation) {
       	v.Required(revel_jwt_mysql.Date)
       	v.Check(revel_jwt_mysql.Title,
       		revel.Required{},
       		revel.MinSize{0},
       		revel.MaxSize{120},
       	)
       	v.Required(revel_jwt_mysql.Picture)
       	v.Required(revel_jwt_mysql.Picture_Link)

       	v.Check(revel_jwt_mysql.Caption,
       		revel.Required{},
       		revel.MinSize{0},
       		revel.MaxSize{120},
       	)
}
