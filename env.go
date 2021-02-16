package main

import (
  "log"
  "math"
  "fmt"
)

type fnType = func(...interface {}) interface{}

type Env struct {
  parent *Env
  vars map[string]interface{}
}

func (env *Env) extend() Env {
  return Env{env, make(map[string]interface{})}
}

func (env *Env) add(name string, val interface{}) {
  env.vars[name] = val
}

func (env *Env) get(name string) interface{} {
  current := env
  for {
    if current == nil {
      log.Fatal(fmt.Sprintf("Undefined variable %s", name))
    } else {
      if _, ok := current.vars[name]; ok {
        return current.vars[name]
      }
      current = current.parent
    }
  }
}

func evalMathOperation(expr AstItem, env Env) interface{} {
  leftOperand := evalUnderEnv(*expr.leftOperand, env).(float64)
  rightOperand := evalUnderEnv(*expr.rightOperand, env).(float64)
  switch expr.val {
    case "+":
      return leftOperand + rightOperand
    case "-":
      return leftOperand - rightOperand
    case "*":
      return leftOperand * rightOperand
    case "/":
      if rightOperand == 0 {
        log.Fatal("Cannot divide by zero")
        return 0
      }
      return leftOperand / rightOperand
    case "%":
      return math.Mod(leftOperand, rightOperand)
    case ">":
      return leftOperand > rightOperand
    case "<":
      return leftOperand < rightOperand
    case ">=":
      return leftOperand >= rightOperand
    case "<=":
      return leftOperand <= rightOperand
    default:
      log.Fatal(fmt.Sprintf("Unrecognized operation %s", expr.val))
      return nil
  }
}

func evalOperation(expr AstItem, env Env) interface{} {
  if expr.leftOperand == nil || expr.rightOperand == nil {
    log.Fatal(fmt.Sprintf("Cannot apply operation %s to nil operands", expr.val))
  }
  if expr.val == "=" {
    leftOperand := evalUnderEnv(*expr.leftOperand, env)
    rightOperand := evalUnderEnv(*expr.rightOperand, env)
    return leftOperand == rightOperand
  } else {
    return evalMathOperation(expr, env)
  }
}

func evalUnderEnv(expr AstItem, env Env) interface{} {
  switch expr.astType {
    case AST_ROOT:
      for _, val := range expr.val.([]AstItem) {
        evalUnderEnv(val, env)
      }
      return nil
    case AST_NUMBER, AST_STRING, AST_BOOLEAN:
      return expr.val
    case AST_NULL:
      return nil
    case AST_IF_CONDITION:
      condition := evalUnderEnv(expr.val.(AstItem), env).(bool)
      if condition {
        return evalUnderEnv(*expr.leftOperand, env)
      } else {
        return evalUnderEnv(*expr.rightOperand, env)
      }
    case AST_VAR_DECLARATION:
      env.add(expr.name, evalUnderEnv(expr.val.(AstItem), env))
      return nil
    case AST_VAR_REFERENCE:
      return env.get(expr.val.(string))
    case AST_OPERATION:
      return evalOperation(expr, env)
    case AST_FN_DECLARATION:
      fn := func(fnArgs ...interface{}) interface{} {
        nestedEnv := env.extend()
        if len(expr.args) != len(fnArgs) {
          log.Fatal(fmt.Sprintf("Incorrect number of arguments supplied to function %s, expected %d, got %d", expr.val, len(expr.args), len(fnArgs)))
        }
        for idx, argName := range expr.args {
          nestedEnv.add(argName, fnArgs[idx])
        }
        return evalUnderEnv(*expr.body, nestedEnv)
      }
      env.add(expr.val.(string), fn)
      return fn
    case AST_FN_CALL:
      var args []interface{}
      for _, val := range expr.parsedArgs {
        args = append(args, evalUnderEnv(val, env))
      }
      fn := evalUnderEnv(expr.val.(AstItem), env).(fnType)
      return fn(args...)
    default:
      log.Fatal(fmt.Sprintf("Unrecognized expression of type %s", expr.astType))
      return nil
  }
}
