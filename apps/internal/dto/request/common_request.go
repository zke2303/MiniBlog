package request

type IDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}
