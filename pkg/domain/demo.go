package domain

// 例子
type Demo struct {
	// 例子id
	DemoId int64 `json:"demoId"`
	// 例子名称
	DemoName string `json:"demoName"`
}

type DemoRepository interface {
	Save(demo *Demo) (*Demo, error)
	Remove(demo *Demo) (*Demo, error)
	FindOne(queryOptions map[string]interface{}) (*Demo, error)
	Find(queryOptions map[string]interface{}) (int64, []*Demo, error)
}

func (demo *Demo) Identify() interface{} {
	if demo.DemoId == 0 {
		return nil
	}
	return demo.DemoId
}

func (demo *Demo) Update(data map[string]interface{}) error {
	if demoName, ok := data["demoName"]; ok {
		demo.DemoName = demoName.(string)
	}
	return nil
}
