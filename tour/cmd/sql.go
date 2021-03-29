package cmd

import (
	"github.com/go-programming-tour-book/tour/internal/sql2struct"
	"github.com/spf13/cobra"
	"log"
)

var username string // 账号
var password string // 密码
var host string // HOST
var charset string // 编码
var dbType string // 数据库类型
var dbName string // 数据库名称
var tableName string // 表名称

var sqlCmd = &cobra.Command{
	Use: "sql",
	Short: "sql转换和处理",
	Long: "sql转换和处理",
	Run: func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
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
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "","","请输入数据库的账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "","","请输入数据库的密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "","127.0.0.1:3306","请输入数据库的HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "","utf8mb4","请输入数据库的编码")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "","mysql","请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "","","请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "","","请输入表名称")

}


