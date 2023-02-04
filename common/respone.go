package common

type successRespone struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessRespone(data, paging, filter interface{}) *successRespone {
	return &successRespone{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessRespone(data interface{}) *successRespone {
	return NewSuccessRespone(data, nil, nil)
}
