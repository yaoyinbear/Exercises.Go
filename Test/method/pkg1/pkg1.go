package pkg1

type MyStruct struct {
    name string // 未导出字段
}

func (m *MyStruct) SetName(name string) {
    m.name = name
}
