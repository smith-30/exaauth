package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	AuditTag = zap.Bool("audit", true)

	OperationFirstFields = []zapcore.Field{
		zap.Bool("first", true),
		zap.Bool("last", false),
		AuditTag,
	}

	OperationFirst = func() func() []interface{} {
		slice := MakeInterfaceSlice(OperationFirstFields)

		return func() []interface{} {
			return slice
		}
	}

	OperationContFields = []zapcore.Field{
		zap.Bool("first", false),
		zap.Bool("last", false),
		AuditTag,
	}

	OperationCont = func() func() []interface{} {
		slice := MakeInterfaceSlice(OperationContFields)

		return func() []interface{} {
			return slice
		}
	}

	OperationLastFields = []zapcore.Field{
		zap.Bool("first", false),
		zap.Bool("last", true),
		AuditTag,
	}

	OperationLast = func() func() []interface{} {
		slice := MakeInterfaceSlice(OperationLastFields)

		return func() []interface{} {
			return slice
		}
	}
)

func MakeInterfaceSlice(zf []zapcore.Field) []interface{} {
	interfaceSlice := make([]interface{}, len(zf))

	for i, d := range zf {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
