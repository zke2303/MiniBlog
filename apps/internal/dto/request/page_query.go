package request

// PageQuery 分页查询请求体
type PageQuery struct {
	Page int `json:"page" binding:"page"`
	Size int `json:"size" binding:"size"`
}
