package template

var ModelTpl = `package models

{{Import}}

type {{DbNameStruct}} struct {
{{SturctInfo}}
}

func (m *{{DbNameStruct}}) TableName() string {
	return "{{DbName}}"
}`
