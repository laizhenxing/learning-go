package main

import (
	"fmt"
	"math"
	"strings"
)

type Env map[Var]float64

type Expr interface {
	Eval(env Env) float64
	// Check reports errors in this Expr and adds its Vars to the set
	Check(vars map[Var]bool) error
}

//  A Var identifies a variable, e.g., x
type Var string

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}

// A literal is a numeric constant, e.g., 3.1314
type literal float64

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune // one of '+', '-'
	x Expr
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return nil
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.Eval(env)
	case '-':
		return -u.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

// A binary represents a binary operator expression, e.g., x+y
type binary struct {
	op rune // one of '+','-'.'*','/'
	x, y Expr
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

type call struct {
	fn string // one of 'pow'. 'sin', 'sqrt'
	args []Expr
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %q", c.fn))
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}


