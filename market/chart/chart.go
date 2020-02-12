package chart

type Chart struct {
	Id string
}

func (c *Chart) GetId() string {
	return c.Id
}
