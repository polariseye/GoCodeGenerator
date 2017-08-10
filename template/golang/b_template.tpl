package config
{{$importList:=GetExtImport .TargetTable}}{{if gt (len $importList) 0}}
import ({{range $importList}}
	"{{.}}"{{end}}
)
{{end}}
// {{.TargetTable.Comment}} {{$nameMaxLen:=MaxLen .TargetTable.ALlColumnNames}}
type {{.TargetTable.TableName |EntityName}} struct { {{range $index,$colItem:=.TargetTable.Columns}}{{if eq $index 0}}
	// {{$colItem.Comment}}
	{{FieldName $colItem.ColumnName}} {{FieldType $colItem}}{{else}}

	// {{$colItem.Comment}}
	{{FieldName $colItem.ColumnName}} {{FieldType $colItem }}{{end}} `gorm:"column:{{$colItem.ColumnName}}{{if $colItem.IsPrimaryKey}};primary_key{{end}}"`{{end}}
}

// 获取对应的数据库表名
// 返回值:
// string:表名
func (this *{{.TargetTable.TableName |EntityName}}) TableName() string {
	return "{{.TargetTable.TableName}}"
}

// 新建{{.TargetTable.Comment}}对象{{range $index,$colItem:=.TargetTable.Columns}}
// {{FirstCharLower (FieldName $colItem.ColumnName)}}:{{$colItem.Comment}}{{end}}
// 返回值:
// *{{.TargetTable.TableName |EntityName}}:{{.TargetTable.Comment}}对象
func New{{.TargetTable.TableName |EntityName}}({{range $index,$colItem:=.TargetTable.Columns}}{{if gt $index 0}}, {{end}}{{$colItem.ColumnName | FieldName | FirstCharLower}} {{FieldType $colItem}}{{end}}) *{{.TargetTable.TableName |EntityName}} {
	return &{{.TargetTable.TableName |EntityName}}{ {{$maxColLen:=MaxLen .TargetTable.ALlColumnNames}}{{range $index,$colItem:=.TargetTable.Columns}}
		{{Assign (printf "%s:" (FieldName $colItem.ColumnName)) " " $maxColLen}} {{$colItem.ColumnName | FieldName | FirstCharLower}},{{end}}
	}
}
