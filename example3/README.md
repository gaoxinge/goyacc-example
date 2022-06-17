```
$ go mod init
$ go get -v github.com/pingcap/parser@3a18f1e
$ go get -v github.com/pingcap/tidb/types/parser_driver@328b6d0
$ go mod tidy
$ go mod vendor
$ go run .
*ast.SelectStmt
*ast.TableOptimizerHint
*ast.FieldList
*ast.SelectField
*ast.ColumnNameExpr
*ast.ColumnName
*ast.SelectField
*ast.ColumnNameExpr
*ast.ColumnName
*ast.SelectField
*ast.ColumnNameExpr
*ast.ColumnName
*ast.TableRefsClause
*ast.Join
*ast.TableSource
*ast.TableName
*ast.BinaryOperationExpr
*ast.BinaryOperationExpr
*ast.BinaryOperationExpr
*ast.ColumnNameExpr
*ast.ColumnName
*driver.ValueExpr
*ast.BinaryOperationExpr
*ast.ColumnNameExpr
*ast.ColumnName
*driver.ValueExpr
*ast.BinaryOperationExpr
*ast.ColumnNameExpr
*ast.ColumnName
*driver.ValueExpr
```