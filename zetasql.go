package main

/*
#include <stdlib.h>
#include "parse_query.h"
*/
import "C"

import (
	"context"
	"unsafe"
)

func IsValidStatement(_ context.Context, sql string) string {
	csql := C.CString(sql)
	defer C.free(unsafe.Pointer(csql))

	cIsValid := C.isValidStatement(csql)
	defer C.free(unsafe.Pointer(cIsValid))

	isValid := C.GoString(cIsValid)
	return isValid
}

func ParseStatement(_ context.Context, sql string) string {
	csql := C.CString(sql)
	defer C.free(unsafe.Pointer(csql))

	cAstOrErr := C.parseStatement(csql)
	defer C.free(unsafe.Pointer(cAstOrErr))

	astOrErr := C.GoString(cAstOrErr)
	return astOrErr
}

func FormatSQL(_ context.Context, sql string) string {
	csql := C.CString(sql)
	defer C.free(unsafe.Pointer(csql))

	cFormattedSQLOrErr := C.formatSQL(csql)
	defer C.free(unsafe.Pointer(cFormattedSQLOrErr))

	formattedSQLOrErr := C.GoString(cFormattedSQLOrErr)
	return formattedSQLOrErr
}
