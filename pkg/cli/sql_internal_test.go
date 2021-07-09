// Copyright 2016 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package cli

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/sql/parser"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/log"
	"github.com/stretchr/testify/assert"
)

func TestIsEndOfStatement(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)

	tests := []struct {
		in         string
		isEnd      bool
		isNotEmpty bool
	}{
		{
			in:         ";",
			isEnd:      true,
			isNotEmpty: true,
		},
		{
			in:         "; /* comment */",
			isEnd:      true,
			isNotEmpty: true,
		},
		{
			in:         "; SELECT",
			isNotEmpty: true,
		},
		{
			in:         "SELECT",
			isNotEmpty: true,
		},
		{
			in:         "SET; SELECT 1;",
			isEnd:      true,
			isNotEmpty: true,
		},
		{
			in:         "SELECT ''''; SET;",
			isEnd:      true,
			isNotEmpty: true,
		},
		{
			in: "  -- hello",
		},
		{
			in:         "select 'abc", // invalid token
			isNotEmpty: true,
		},
		{
			in:         "'abc", // invalid token
			isNotEmpty: true,
		},
		{
			in:         `SELECT e'\xaa';`, // invalid token but last token is semicolon
			isEnd:      true,
			isNotEmpty: true,
		},
	}

	for _, test := range tests {
		lastTok, isNotEmpty := parser.LastLexicalToken(test.in)
		if isNotEmpty != test.isNotEmpty {
			t.Errorf("%q: isNotEmpty expected %v, got %v", test.in, test.isNotEmpty, isNotEmpty)
		}
		isEnd := isEndOfStatement(lastTok)
		if isEnd != test.isEnd {
			t.Errorf("%q: isEnd expected %v, got %v", test.in, test.isEnd, isEnd)
		}
	}
}

// Test handleCliCmd cases for client-side commands that are aliases for sql
// statements.
func TestHandleCliCmdSqlAlias(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	initCLIDefaults()

	clientSideCommandTestsTable := []struct {
		commandString string
		wantSQLStmt   string
	}{
		{`\l`, `SHOW DATABASES`},
		{`\dt`, `SHOW TABLES`},
		{`\dT`, `SHOW TYPES`},
		{`\du`, `SHOW USERS`},
		{`\d mytable`, `SHOW COLUMNS FROM mytable`},
		{`\d`, `SHOW TABLES`},
	}

	var c cliState
	for _, tt := range clientSideCommandTestsTable {
		c = setupTestCliState()
		c.lastInputLine = tt.commandString
		gotState := c.doHandleCliCmd(cliStateEnum(0), cliStateEnum(1))

		assert.Equal(t, cliRunStatement, gotState)
		assert.Equal(t, tt.wantSQLStmt, c.concatLines)
	}
}

func TestHandleCliCmdSlashDInvalidSyntax(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	initCLIDefaults()

	clientSideCommandTests := []string{`\d goodarg badarg`, `\dz`}

	var c cliState
	for _, tt := range clientSideCommandTests {
		c = setupTestCliState()
		c.lastInputLine = tt
		gotState := c.doHandleCliCmd(cliStateEnum(0), cliStateEnum(1))

		assert.Equal(t, cliStateEnum(0), gotState)
		assert.Equal(t, errInvalidSyntax, c.exitErr)
	}
}

func TestHandleDemoNodeCommandsInvalidNodeName(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	initCLIDefaults()

	demoNodeCommandTests := []string{"shutdown", "*"}

	c := setupTestCliState()
	c.handleDemoNodeCommands(demoNodeCommandTests, cliStateEnum(0), cliStateEnum(1))
	assert.Equal(t, errInvalidSyntax, c.exitErr)
}

func setupTestCliState() cliState {
	c := cliState{}
	c.ins = noLineEditor
	return c
}
