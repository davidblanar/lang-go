package main

import (
	"strings"
	"fmt"
)

type AstItem struct {
	astType string
	val interface{}
	leftOperand *AstItem
	rightOperand *AstItem
	name string
	args []string
	parsedArgs []AstItem
	body *AstItem
}

type Parser struct {
	tokens []Token
	pos int
	ast []AstItem
}

func (p *Parser) generate() AstItem {
	for {
		if p._hasNext() {
			p.ast = append(p.ast, p._parseExpression())
		} else {
			break
		}
	}
	return AstItem{
		astType: AST_ROOT,
		val: p.ast,
		leftOperand: nil,
		rightOperand: nil,
		name: "",
		args: nil,
		parsedArgs: nil,
		body: nil,
	}
}

func (p *Parser) _parseExpression() AstItem {
	p._requireVal("(")
	p._next()
	var expr AstItem
	for {
		if p._isEndOfExpression() {
			break
		}
		token := p._peek()
		tokenType := token.tokenType
		val := token.val
		switch tokenType {
			case TOKEN_NUMBER:
				expr = p._parseNumber()
				break
			case TOKEN_STRING:
				expr = p._parseString()
				break
			case TOKEN_SYMBOL:
				expr = p._parseSymbol()
				break
			case TOKEN_IDENTIFIER:
				if val == "var" {
					expr = p._parseVarDeclaration()
					break
				}
			  	if val == "fn" {
					expr = p._parseFnDeclaration()
					break
			  	}
				if val == "call" {
					expr = p._parseFnCall()
					break
				}
				if p._isBoolean(val.(string)) {
					expr = p._parseBoolean()
					break
				}
				if p._isNull(val.(string)) {
					expr = p._parseNull()
					break
				}
				if val == "if" {
					expr = p._parseIfCondition()
					break
				}
				expr = p._parseVarReference()
				break
			default:
				throwError(fmt.Sprintf("Unrecognized token type %s, value %s", tokenType, val), token.lineno, token.col)
				break
		}
	}
	p._requireVal(")")
	p._next()
	return expr
}

func (p *Parser) _requireType(expected string) {
	token := p._peek()
	if token.tokenType != expected {
		throwError(fmt.Sprintf("Unexpected expression of type %s, expected %s", token.tokenType, expected), token.lineno, token.col)
	}
}

func (p *Parser) _requireVal(expected string) {
	token := p._peek()
	if token.val != expected {
		throwError(fmt.Sprintf("Unexpected token %s, expected %s", token.val, expected), token.lineno, token.col)
	}
}

func (p *Parser) _isBoolean(tokenVal string) bool {
	return tokenVal == "true" || tokenVal == "false"
}

func (p *Parser) _isNull(tokenVal string) bool {
	return tokenVal == "null"
}

func (p *Parser) _parseNumber() AstItem {
	return AstItem{
		astType: AST_NUMBER,
		val: p._next().val,
		leftOperand: nil,
		rightOperand: nil,
		name: "",
		args: nil,
		parsedArgs: nil,
		body: nil,
	}
}

func (p *Parser) _parseString() AstItem {
	return AstItem{
		astType: AST_STRING,
		val: p._next().val,
		leftOperand: nil,
		rightOperand: nil,
		name: "",
		args: nil,
		parsedArgs: nil,
		body: nil,
	}
}

func (p *Parser) _parseBoolean() AstItem {
	return AstItem{
		astType: AST_BOOLEAN,
		val: p._next().val == "true",
		leftOperand: nil,
		rightOperand: nil,
		name: "",
		args: nil,
		parsedArgs: nil,
		body: nil,
	}
}

func (p *Parser) _parseNull() AstItem {
	p._next()
	return AstItem{
		astType: AST_NULL,
		val: nil,
		leftOperand: nil,
		rightOperand: nil,
		name: "",
		args: nil,
		parsedArgs: nil,
		body: nil,
	}
}

func (p *Parser) _parseIfCondition() AstItem {
	p._next()
	condition := p._parseExpression();
	leftOperand := p._parseExpression();
	rightOperand := p._parseExpression();
	return AstItem{
		astType: AST_IF_CONDITION,
		val: condition,
		leftOperand: &leftOperand,
		rightOperand: &rightOperand,
		name: "",
		args: nil,
		parsedArgs: nil,
		body: nil,
	}
}

func (p *Parser) _parseVarDeclaration() AstItem {
	p._next()
	p._requireType(TOKEN_IDENTIFIER)
	varName := p._next().val.(string)
	return AstItem{
		astType: AST_VAR_DECLARATION,
		val: p._parseVarValue(),
		leftOperand: nil,
		rightOperand: nil,
		name: varName,
		args: nil,
		parsedArgs: nil,
		body: nil,
	}
}

func (p *Parser) _parseVarValue() AstItem {
	if p._isBeginningOfExpression() {
		return p._parseExpression()
	}
	token := p._next()
	val := token.val
	tokenType := token.tokenType

	switch tokenType {
		case TOKEN_NUMBER:
			return AstItem{
				astType: AST_NUMBER,
				val: val,
				leftOperand: nil,
				rightOperand: nil,
				name: "",
				args: nil,
				parsedArgs: nil,
				body: nil,
			}
		case TOKEN_STRING:
			return AstItem{
				astType: AST_STRING,
				val: val,
				leftOperand: nil,
				rightOperand: nil,
				name: "",
				args: nil,
				parsedArgs: nil,
				body: nil,
			}
		case TOKEN_IDENTIFIER:
			if p._isBoolean(val.(string)) {
				return AstItem{
					astType: AST_BOOLEAN,
					val: val == "true",
					leftOperand: nil,
					rightOperand: nil,
					name: "",
					args: nil,
					parsedArgs: nil,
					body: nil,
				}
			}
			if (p._isNull(val.(string))) {
				return AstItem{
					astType: AST_NULL,
					val: nil,
					leftOperand: nil,
					rightOperand: nil,
					name: "",
					args: nil,
					parsedArgs: nil,
					body: nil,
				}
			}
			return AstItem{
				astType: AST_VAR_REFERENCE,
				val: val,
				leftOperand: nil,
				rightOperand: nil,
				name: "",
				args: nil,
				parsedArgs: nil,
				body: nil,
			}
		default:
			throwError(fmt.Sprintf("Unexpected token type %s when parsing variable value", tokenType), token.lineno, token.col)
			return AstItem{}
	}
}

func (p *Parser) _parseVarReference() AstItem {
	return AstItem{
		astType: AST_VAR_REFERENCE,
		val: p._next().val,
		leftOperand: nil,
		rightOperand: nil,
		name: "",
		args: nil,
		parsedArgs: nil,
		body: nil,
	}
}

func (p *Parser) _parseSymbol() AstItem {
	var sb strings.Builder
	operation := p._next().val
	sb.WriteString(operation.(string))
	if operation == ">" || operation == "<" {
		if p._peek().val == "=" {
			sb.WriteString(p._next().val.(string))
		}
	}
	leftOperand := p._parseVarValue()
	rightOperand := p._parseVarValue()
	return AstItem{
		astType: AST_OPERATION,
		val: sb.String(),
		leftOperand: &leftOperand,
		rightOperand: &rightOperand,
		name: "",
		args: nil,
		parsedArgs: nil,
		body: nil,
	}
}

func (p *Parser) _parseFnDeclaration() AstItem {
	p._next()
	p._requireType(TOKEN_IDENTIFIER)
	isFnName := true
	var fnName string
	var body AstItem
	var args []string
	for {
		if p._isEndOfExpression() {
			break
		}
		if isFnName {
			isFnName = false
			fnName = p._next().val.(string)
		}
		if p._isBeginningOfExpression() {
			body = p._parseExpression()
		} else {
			args = append(args, p._next().val.(string))
		}
	}
	return AstItem{
		astType: AST_FN_DECLARATION,
		val: fnName,
		leftOperand: nil,
		rightOperand: nil,
		name: "",
		args: args,
		parsedArgs: nil,
		body: &body,
	}
}

func (p *Parser) _parseFnCall() AstItem {
	p._next()
	p._requireType(TOKEN_IDENTIFIER)
	isFnReference := true
	var fnReference AstItem
	var args []AstItem
	for {
		if p._isEndOfExpression() {
			break
		}
		if isFnReference {
			isFnReference = false
			fnReference = p._parseVarReference()
			continue
		}
		if p._isBeginningOfExpression() {
			args = append(args, p._parseExpression())
		} else {
			args = append(args, p._parseVarValue())
		}
	}
	return AstItem{
		astType: AST_FN_CALL,
		val: fnReference,
		leftOperand: nil,
		rightOperand: nil,
		name: "",
		args: nil,
		parsedArgs: args,
		body: nil,
	}
}

func (p *Parser) _isBeginningOfExpression() bool {
	return p._peek().val == "("
}

func (p *Parser) _isEndOfExpression() bool {
	return p._peek().val == ")"
}

func (p *Parser) _hasNext() bool {
	return p.pos < len(p.tokens)
}

func (p *Parser) _next() Token {
	t := p.tokens[p.pos]
	p.pos++
	return t
}

func (p *Parser) _peek() Token {
	return p.tokens[p.pos]
}
