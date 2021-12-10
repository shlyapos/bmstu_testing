package builder

import "skema/app/model"

type CommentBuilder struct {
	Id     uint
	Owner  uint
	Schema uint
	Data   string
}

func (b *CommentBuilder) Build() *model.Comment {
	return &model.Comment{
		Id:     b.Id,
		Owner:  b.Owner,
		Schema: b.Schema,
		Data:   b.Data,
	}
}

type CommentMother struct{}

func (m *CommentMother) Obj0() *model.Comment {
	builder := CommentBuilder{
		Id:     1,
		Owner:  2,
		Schema: 2,
		Data:   "Good uml schema",
	}

	return builder.Build()
}

func (m *CommentMother) Obj1() *model.Comment {
	builder := CommentBuilder{
		Id:     2,
		Owner:  3,
		Schema: 4,
		Data:   "I like how this schema look",
	}

	return builder.Build()
}

func (m *CommentMother) Obj2() *model.Comment {
	builder := CommentBuilder{
		Id:     3,
		Owner:  3,
		Schema: 2,
		Data:   "Use case is wrong!",
	}

	return builder.Build()
}

func (m *CommentMother) Obj3() *model.Comment {
	builder := CommentBuilder{
		Id:     4,
		Owner:  1,
		Schema: 1,
		Data:   "I cant understand how make uml class diagram",
	}

	return builder.Build()
}

func (m *CommentMother) Obj4() *model.Comment {
	builder := CommentBuilder{
		Id:     5,
		Owner:  3,
		Schema: 3,
		Data:   "I like drawing schemes",
	}

	return builder.Build()
}

func (m *CommentMother) Obj5() *model.Comment {
	builder := CommentBuilder{
		Id:     6,
		Owner:  3,
		Schema: 2,
		Data:   "Skema is cool",
	}

	return builder.Build()
}
