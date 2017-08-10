//***********************************************************************************
// 文件名称：{{.TargetTable.TableName |EntityName}}.cs
// 功能描述：{{.TargetTable.Comment}}
// 数据表：{{.TargetTable.TableName}}
// 作者：{{.BaseConfig.HeaderInfo.AuthorName}}
// 日期：{{TimeFormat (Now) .BaseConfig.HeaderInfo.TimeFormat}}
// 修改记录：
//***********************************************************************************

using System;
using System.Data;

namespace Moqikaka.GameServer.Inuyasha.Model
{
    /// <summary>
    /// {{.TargetTable.Comment}}
    /// </summary>
    [DBTable("{{.TargetTable.TableName}}")]
    public sealed class {{.TargetTable.TableName |EntityName}}
    {
        #region 属性
{{range .TargetTable.Columns}}
        /// <summary>
        /// {{.Comment}}
        /// </summary>
        public {{FieldType .}} {{FirstCharLower (FieldName .ColumnName)}} { get; private set; }
{{end}}
        #endregion
    }
}