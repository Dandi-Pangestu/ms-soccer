package domains

type Paginator struct {
	TotalRecord int         `json:"total_record"`
	TotalPage   int         `json:"total_page"`
	Data        interface{} `json:"data"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"next_page"`
}
