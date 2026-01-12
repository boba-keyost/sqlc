package ast

import "github.com/boba-keyost/sqlc/internal/sql/format"

type CallStmt struct {
	FuncCall *FuncCall
}

func (n *CallStmt) Pos() int {
	if n.FuncCall == nil {
		return 0
	}
	return n.FuncCall.Pos()
}

func (n *CallStmt) Format(buf *TrackedBuffer, d format.Dialect) {
	buf.WriteString("CALL ")
	buf.astFormat(n.FuncCall, d)
}
