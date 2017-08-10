//***********************************************************************************
// 文件名称：{{.TargetTable.TableName |EntityName}}.cs
// 功能描述：{{.TargetTable.Comment}}
// 数据表：{{.TargetTable.TableName}}
// 作者：{{.BaseConfig.HeaderInfo.AuthorName}}
// 日期：{{TimeFormat (Now) .BaseConfig.HeaderInfo.TimeFormat}}
// 修改记录：
//***********************************************************************************

using System;
using System.Collections.Generic;

namespace Moqikaka.GameServer.Inuyasha.Model
{
	/// <summary>
	/// {{.TargetTable.Comment}}
	/// </summary>
	public sealed class {{.TargetTable.TableName |EntityName}} : BaseGlobal
	{
		#region 属性{{range .TargetTable.Columns}}

        /// <summary>
        /// {{.Comment}}
        /// </summary>
        public {{FieldType .}} {{FieldName .ColumnName}} { get; private set; }{{end}}

		#endregion

		#region 构造函数

        /// <summary>
        /// 构造函数
        /// </summary>
		/// <param name="global">全局模型对象</param>{{range $index,$colItem:=.TargetTable.Columns}}
        /// <param name="{{FirstCharLower (FieldName $colItem.ColumnName)}}">{{$colItem.Comment}}</param>{{end}}
        public {{.TargetTable.TableName |EntityName}}(Global global{{range .TargetTable.Columns}}, {{FieldType .}} {{FirstCharLower (FieldName .ColumnName)}}{{end}})
            : base(global)
        {
{{range .TargetTable.Columns}}            this.{{FieldName .ColumnName}} = {{FirstCharLower (FieldName .ColumnName)}};
{{end}}
        }

        /// <summary>
        /// 构造函数
        /// </summary>
		/// <param name="global">全局模型对象</param>
        /// <param name="dr">System.Data.DataRow</param>
        public {{.TargetTable.TableName |EntityName}}(Global global, System.Data.DataRow dr)
            : base(global)
        {
{{range .TargetTable.Columns}}            this.{{FieldName .ColumnName}} = {{if eq (FieldType .) "Guid"}}Guid.Parse(dr[GamePropertyName.{{FieldName .ColumnName}}].ToString()){{else}}Convert.To{{FieldType .}}(dr[GamePropertyName.{{FieldName .ColumnName}}]){{end}};
{{end}}        }

		#endregion

		#region 方法

        /// <summary>
        /// 构造对象
        /// </summary>{{range $index,$colItem:=.TargetTable.Columns}}
        /// <param name="{{FirstCharLower (FieldName $colItem.ColumnName)}}">{{$colItem.Comment}}</param>{{end}}
        public void ConstructObject({{range $index,$colItem:=.TargetTable.Columns}}{{if ne $index 0}}, {{end}}{{FieldType $colItem}} {{FirstCharLower (FieldName $colItem.ColumnName)}}{{end}})
        {
{{range .TargetTable.Columns}}            this.{{FieldName .ColumnName}} = {{FirstCharLower (FieldName .ColumnName)}};
{{end}}        }

		#endregion
	}
}
