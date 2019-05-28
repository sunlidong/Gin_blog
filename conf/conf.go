package conf

// parmas
const (
	HtmlFile   = Filepath + File
	StaticFile = Filepath + Static
	// 定义查询文章，每页显示的文章量
	NUM = 5
)

//base
const (
	//
	Filepath = "C:/Users/sld/go/src/Gin_book/"
	//
	File   = "view/*"
	Static = "static/"
)

// init sql

const (
	DriverName   = "sql"
	SqlLoginName = "root"
	SqlLoginPass = "1234567"
	SqlUrl       = "127.0.0.1"
	SqlPort      = "3306"
	SqlDB        = "kalablock"
	SqlUnit      = "utf8"
)
