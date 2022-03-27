package model

type (
	TagCreateInput struct {
		Id      int
		TagName string
	}

	TagUpdateInput struct {
		TagName string
	}
)
