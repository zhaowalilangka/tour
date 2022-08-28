package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zhaowalilangka/tour/internal/sql2struct"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换与处理",
	Long:  "sql转换与处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2StructCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转换",
	Long:  "sql转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbmodel.Connect error:%v", err)
		}

		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbmodel.GetColumns error:%v", err)
		}
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("dbmodel.Generate error:%v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2StructCmd)
	sql2StructCmd.Flags().StringVarP(&username, "username", "", "zlb", "请输入数据库的账号")
	sql2StructCmd.Flags().StringVarP(&password, "password", "", "zlb199153", "请输入数据库的密码")
	sql2StructCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的HOST")
	sql2StructCmd.Flags().StringVarP(&charset, "charset", "", "utf8", "请输入数据库的编码")
	sql2StructCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库的示例类型")
	sql2StructCmd.Flags().StringVarP(&dbName, "db", "", "wms", "请输入数据库名称")
	sql2StructCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}
