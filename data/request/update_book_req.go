package request

type UpdateBookReq struct {
	//define the properties for updating this struct [class] and how to display in json
	Id   int
	Name string `validate:"required min=1,max=100" json:"name"`
}
