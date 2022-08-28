package sql2struct

import (
	"fmt"
	"os"
	"text/template"

	"github.com/zhaowalilangka/tour/internal/word"
)

const structTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }} {{.Name | ToCamelCase}} {{.Type}} {{.Tag}} {{else}} {{.Name}} {{end}} 	 	{{$length := len .Comment}} {{if gt $length 0 }}//{{.Comment}} {{else}}// {{.Name}} {{end}}
{{end}}}
`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColmns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColmns))
	for _, column := range tbColmns {
		tag := fmt.Sprintf("`json:\"%s\" db:\"%s\" form:\"%s\"`", column.ColumnName, column.ColumnName, column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2Struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))
	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}
