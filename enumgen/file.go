// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on http://github.com/dmarkham/enumer and
// golang.org/x/tools/cmd/stringer:

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package enumgen

import (
	"go/ast"
	exact "go/constant"
	"go/token"
	"go/types"
	"log"
	"strings"
)

// File holds a single parsed file and associated data.
type File struct {
	Pkg  *Package  // Package to which this file belongs.
	File *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	TypeName string  // The name of the constant type we are currently looking for.
	BitFlag  bool    // Whether the constant type we are currently looking for is a bitflag.
	Values   []Value // Accumulator for constant values of that type.
	Config   Config  // The configuration information
}

// GenDecl processes one declaration clause.
func (f *File) GenDecl(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if !ok || decl.Tok != token.CONST {
		// We only care about const declarations.
		return true
	}
	// The name of the type of the constants we are declaring.
	// Can change if this is a multi-element declaration.
	typ := ""
	// Loop over the elements of the declaration. Each element is a ValueSpec:
	// a list of names possibly followed by a type, possibly followed by values.
	// If the type and value are both missing, we carry down the type (and value,
	// but the "go/types" package takes care of that).
	for _, spec := range decl.Specs {
		vspec := spec.(*ast.ValueSpec) // Guaranteed to succeed as this is CONST.
		if vspec.Type == nil && len(vspec.Values) > 0 {
			// "X = 1". With no type but a value, the constant is untyped.
			// Skip this vspec and reset the remembered type.
			typ = ""
			continue
		}
		if vspec.Type != nil {
			// "X T". We have a type. Remember it.
			ident, ok := vspec.Type.(*ast.Ident)
			if !ok {
				continue
			}
			typ = ident.Name
		}
		if typ != f.TypeName {
			// This is not the type we're looking for.
			continue
		}
		// We now have a list of names (from one line of source code) all being
		// declared with the desired type.
		// Grab their names and actual values and store them in f.values.
		for _, n := range vspec.Names {
			if n.Name == "_" {
				continue
			}
			// This dance lets the type checker find the values for us. It's a
			// bit tricky: look up the object declared by the n, find its
			// types.Const, and extract its value.
			obj, ok := f.Pkg.Defs[n]
			if !ok {
				log.Fatalf("no value for constant %s", n)
			}
			info := obj.Type().Underlying().(*types.Basic).Info()
			if info&types.IsInteger == 0 {
				log.Fatalf("can't handle non-integer constant type %s", typ)
			}
			value := obj.(*types.Const).Val() // Guaranteed to succeed as this is CONST.
			if value.Kind() != exact.Int {
				log.Fatalf("can't happen: constant is not an integer %s", n)
			}
			i64, isInt := exact.Int64Val(value)
			u64, isUint := exact.Uint64Val(value)
			if !isInt && !isUint {
				log.Fatalf("internal error: value of %s is not an integer: %s", n, value.String())
			}
			if !isInt {
				u64 = uint64(i64)
			}
			v := Value{
				originalName: n.Name,
				name:         n.Name,
				value:        u64,
				signed:       info&types.IsUnsigned == 0,
				str:          value.String(),
			}
			if c := vspec.Comment; f.Config.LineComment && c != nil && len(c.List) == 1 {
				v.name = strings.TrimSpace(c.Text())
			}

			f.Values = append(f.Values, v)
		}
	}
	return false
}