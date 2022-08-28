package gorme

type T struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type T2 struct {
	Uid      int      `json:"uid"`
	UserName string   `json:"userName"`
	Age      float64  `json:"age"`
	Likes    []string `json:"likes"`
}
