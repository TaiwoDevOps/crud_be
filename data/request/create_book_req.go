package request

type CreateBookReq struct {
	//define the property of this struct [class] and how to display in json
	Name string `validate:"required min=1,max=100" json:"name"`
}
