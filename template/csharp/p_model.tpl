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
	public sealed class {{.TargetTable.TableName |EntityName}} : BasePlayer
	{
		#region 属性{{range (GetColumnsExcloud .TargetTable "PlayerId")}}

        /// <summary>
        /// {{.Comment}}
        /// </summary>
        public {{FieldType .}} {{FieldName .ColumnName}} { get; private set; }{{end}}

		#endregion

		#region 构造函数

        /// <summary>
        /// 构造函数
        /// </summary>
        /// <param name="player">玩家相关数据对象</param>{{range $index,$colItem:=(GetColumnsExcloud .TargetTable "PlayerId")}}
        /// <param name="{{FirstCharLower (FieldName $colItem.ColumnName)}}">{{$colItem.Comment}}</param>{{end}}
        public {{.TargetTable.TableName |EntityName}}(Player player{{range (GetColumnsExcloud .TargetTable "PlayerId")}}, {{FieldType .}} {{FirstCharLower (FieldName .ColumnName)}}{{end}})
            : base(player)
        {
{{range (GetColumnsExcloud .TargetTable "PlayerId")}}            this.{{FieldName .ColumnName}} = {{FirstCharLower (FieldName .ColumnName)}};
{{end}}
        }

        /// <summary>
        /// 构造函数
        /// </summary>
        /// <param name="player">玩家相关数据对象</param>
        /// <param name="dr">System.Data.DataRow</param>
        public {{.TargetTable.TableName |EntityName}}(Player player, System.Data.DataRow dr)
            : base(player)
        {
{{range (GetColumnsExcloud .TargetTable "PlayerId")}}            this.{{FieldName .ColumnName}} = {{if eq (FieldType .) "Guid"}}Guid.Parse(dr[GamePropertyName.{{FieldName .ColumnName}}].ToString()){{else}}Convert.To{{FieldType .}}(dr[GamePropertyName.{{FieldName .ColumnName}}]){{end}};
{{end}}        }

		#endregion

		#region 方法

        /// <summary>
        /// 构造对象
        /// </summary>{{range $index,$colItem:=(GetColumnsExcloud .TargetTable "PlayerId")}}
        /// <param name="{{FirstCharLower (FieldName $colItem.ColumnName)}}">{{$colItem.Comment}}</param>{{end}}
        public void ConstructObject({{range $index,$colItem:=(GetColumnsExcloud .TargetTable "PlayerId")}}{{if ne $index 0}}, {{end}}{{FieldType $colItem}} {{FirstCharLower (FieldName $colItem.ColumnName)}}{{end}})
        {
{{range (GetColumnsExcloud .TargetTable "PlayerId")}}            this.{{FieldName .ColumnName}} = {{FirstCharLower (FieldName .ColumnName)}};
{{end}}        }

		#endregion
	}
}
