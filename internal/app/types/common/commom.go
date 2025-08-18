package common

type ListReq struct {
	BaseListParam
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

func (l *ListReq) Adjust() {
	if l.Page == 0 {
		l.Page = 1
	}
	if l.PageSize == 0 {
		l.PageSize = 10
	}
}

func (l *ListReq) GetOffset() int {
	return (l.Page - 1) * l.PageSize
}

//------------- resp -------------//

type ListResp struct {
	BaseListResp
	Total     int64 `json:"total"`
	Current   int   `json:"current"`
	Size      int   `json:"size"`
	TotalPage int   `json:"totalPage"`
}

func (l *ListResp) Adjust() {
	if l.Size == 0 {
		l.TotalPage = 0
		return
	}
	l.TotalPage = (int(l.Total) + l.Size - 1) / l.Size
}

func (l *ListResp) GetTotalPage() int {
	if l.Size == 0 {
		return 0
	}
	return (int(l.Total) + l.Size - 1) / l.Size
}

type IdReq struct {
	Id int64 `json:"id" binding:"required"`
}
