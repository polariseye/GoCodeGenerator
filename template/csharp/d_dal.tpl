//***********************************************************************************
// 文件名称：{{.TargetTable.TableName |EntityName}}DAL.cs{{$primaryCols:=(GetColumns .TargetTable true)}}
// 功能描述：{{.TargetTable.Comment}}
// 数据表：{{.TargetTable.TableName}}
// 作者：{{.BaseConfig.HeaderInfo.AuthorName}}
// 日期：{{TimeFormat (Now) .BaseConfig.HeaderInfo.TimeFormat}}
// 修改记录：
//***********************************************************************************

using System;
using System.Data;

namespace Moqikaka.GameServer.Inuyasha.DAL
{
    using MySql.Data.MySqlClient;

    /// <summary>
    /// {{.TargetTable.Comment}}
    /// </summary>
    public class {{.TargetTable.TableName |EntityName}}DAL : ModelBaseDAL
    {
        #region 属性

        //定义表的所有字段（以,  进行分隔）
        const String columns = @"{{range $index,$colItem:=.TargetTable.Columns}}{{if ne $index 0}}, {{end}}{{$colItem.ColumnName}}{{end}}";

        //定义此类所能用到的所有参数名称（以pn = ParameterName开始）；双引号内是@+字段名称{{range .TargetTable.Columns}}
        const String pn{{FieldName .ColumnName}} = "@{{FieldName .ColumnName}}";{{end}}

        //定义此类所能用到的所有sql语句（以Command结尾，表示为sql命令）
        static String getAllCommand = String.Empty;
        static String insertCommand = String.Empty;
        static String updateCommand = String.Empty;

        #endregion

        #region 构造函数

        /// <summary>
        /// 初始化此类所能用到的所有sql语句
        /// </summary>
        static {{.TargetTable.TableName |EntityName}}DAL()
        {
            getAllCommand = String.Format("SELECT {0} FROM {{.TargetTable.TableName}};", columns);
            insertCommand = "INSERT INTO {{.TargetTable.TableName}} ({{range $index,$colItem:=.TargetTable.Columns}}{{if ne $index 0}}, {{end}}{{$colItem.ColumnName}}{{end}}) VALUES ({{range $index,$colItem:=.TargetTable.Columns}}{{if ne $index 0}}, {{end}}@{{FieldName $colItem.ColumnName}}{{end}});";
            updateCommand = "UPDATE {{.TargetTable.TableName}} SET {{range $index,$colItem:=(GetColumns .TargetTable false)}}{{if ne $index 0}}, {{end}}{{$colItem.ColumnName}} = @{{FieldName $colItem.ColumnName}}{{end}}{{range $index,$colItem:=(GetColumns .TargetTable true)}}{{if eq $index 0}} WHERE{{else}} AND{{end}} {{.ColumnName}} = @{{FieldName .ColumnName}}{{end}};";
        }

        #endregion

        #region 数据操作

        /// <summary>
        /// 获取所有数据
        /// </summary>
        /// <returns>数据列表</returns>
        public static DataTable GetList()
        {
            return ExecuteDataTable(getAllCommand);
        }

        /// <summary>
        /// 更新数据
        /// </summary>{{range .TargetTable.Columns}}
        /// <param name="{{FirstCharLower (FieldName .ColumnName)}}">{{.Comment}}</param>{{end}}
        /// <returns>受影响的行数</returns>
        public static Int32 ReplaceInfo({{range $index,$colItem:=.TargetTable.Columns}}{{if ne $index 0}}, {{end}}{{FieldType $colItem}} {{FirstCharLower (FieldName $colItem.ColumnName)}}{{end}})
        {
            //构造参数
            MySqlParameter[] paramList = new MySqlParameter[]
			{
{{range .TargetTable.Columns}}				new MySqlParameter(pn{{FieldName .ColumnName}}, {{FirstCharLower (FieldName .ColumnName)}}), 
{{end}}			};

            Int32 rows = ExecuteNonQuery(updateCommand, paramList);
            if (rows == 0)
            {
                rows = ExecuteNonQuery(insertCommand, paramList);
            }

            return rows;
        }

        /// <summary>
        /// 插入数据
        /// </summary>{{range .TargetTable.Columns}}
        /// <param name="{{FirstCharLower (FieldName .ColumnName)}}">{{.Comment}}</param>{{end}}
        /// <returns>受影响的行数</returns>
        public static Int32 InsertInfo({{range $index,$colItem:=.TargetTable.Columns}}{{if ne $index 0}}, {{end}}{{FieldType $colItem}} {{FirstCharLower (FieldName $colItem.ColumnName)}}{{end}})
        {
            //构造参数
            MySqlParameter[] paramList = new MySqlParameter[]
			{
{{range .TargetTable.Columns}}				new MySqlParameter(pn{{FieldName .ColumnName}}, {{FirstCharLower (FieldName .ColumnName)}}), 
{{end}}			};

            return ExecuteNonQuery(insertCommand, paramList);
        }

        /// <summary>
        /// 更新数据
        /// </summary>{{range .TargetTable.Columns}}
        /// <param name="{{FirstCharLower (FieldName .ColumnName)}}">{{.Comment}}</param>{{end}}
        /// <returns>受影响的行数</returns>
        public static Int32 UpdateInfo({{range $index,$colItem:=.TargetTable.Columns}}{{if ne $index 0}}, {{end}}{{FieldType $colItem}} {{FirstCharLower (FieldName $colItem.ColumnName)}}{{end}})
        {
            //构造参数
            MySqlParameter[] paramList = new MySqlParameter[]
			{
{{range .TargetTable.Columns}}				new MySqlParameter(pn{{FieldName .ColumnName}}, {{FirstCharLower (FieldName .ColumnName)}}), 
{{end}}			};

            return ExecuteNonQuery(updateCommand, paramList);
        }

        /// <summary>
        /// 删除数据
        /// </summary>{{range $primaryCols}}
        /// <param name="{{FirstCharLower (FieldName .ColumnName)}}">{{.Comment}}</param>{{end}}
        /// <returns>受影响的行数</returns>
        public static Int32 DeleteInfo({{range $index,$colItem:=$primaryCols}}{{if ne $index 0}}, {{end}}{{FieldType $colItem}} {{FirstCharLower (FieldName $colItem.ColumnName)}}{{end}})
        {
            //构造参数
            MySqlParameter[] paramList = new MySqlParameter[]
			{
{{range $primaryCols}}				new MySqlParameter(pn{{FieldName .ColumnName}}, {{FirstCharLower (FieldName .ColumnName)}}), 
{{end}}			};

            return ExecuteNonQuery(deleteCommand, paramList);
        }

        #endregion
    }
}