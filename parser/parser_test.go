package parser

import (
	"testing"

	"github.com/lavish-gambhir/qu-interpreter/ast"
	"github.com/lavish-gambhir/qu-interpreter/lexer"
)

func TestValStatement(t *testing.T) {
	input := `
	val x = 5;
	val y = 10;
	val foobar = 41231;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got = %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testValStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testValStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "val" {
		t.Errorf("s.TokenLiteral not 'val', got = %q", s.TokenLiteral())
		return false
	}
	valstmt, ok := s.(*ast.ValStatement)
	if !ok {
		t.Errorf("s not *ast.ValStatement, got = %T", s)
		return false
	}
	if valstmt.Name.Value != name {
		t.Errorf("valstmt.Name.Value not '%s', got = %s", name, valstmt.Name.Value)
		return false
	}
	if valstmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s', got = %s", name, valstmt.Name)
		return false
	}
	return true
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
