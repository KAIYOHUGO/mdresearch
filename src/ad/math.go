package ad

import "math"

func Add(value ...*Node) *Node {
	return buildOP(valueAdd, partialAdd, value...)
}

func (s *Node) Add(value *Node) *Node {
	return buildOP(valueAdd, partialAdd, s, value)
}

func Sub(value ...*Node) *Node {
	return buildOP(valueSub, partialSub, value...)
}

func (s *Node) Sub(value *Node) *Node {
	return buildOP(valueSub, partialSub, s, value)
}

func Mult(value ...*Node) *Node {
	return buildOP(valueMult, partialMult, value...)
}

func (s *Node) Mult(value *Node) *Node {
	return buildOP(valueMult, partialMult, s, value)
}

func Div(value ...*Node) *Node {
	return buildOP(valueDiv, partialDiv, value...)
}

func (s *Node) Div(value *Node) *Node {
	return buildOP(valueDiv, partialDiv, s, value)
}

func Neg(value *Node) *Node {
	return buildOP(valueMult, partialMult, Varible(-1), value)
}

func (s *Node) Neg() *Node {
	return buildOP(valueMult, partialMult, Varible(-1), s)
}

func buildOP(calcValue calcValue, calcPartial calcPartial, value ...*Node) *Node {
	node := newNode(calcValue, calcPartial)
	for _, v := range value {
		if node.rSide == nil {
			node.rSide = newSide(v)
			continue
		}
		if node.lSide == nil {
			node.lSide = newSide(v)
			continue
		}
		newnode := newNode(calcValue, calcPartial)
		newnode.rSide, node = newSide(node), newnode
	}
	return node
}

func partialAdd(_, _ *Node) (float64, float64) {
	return 1, 1
}

func partialSub(_, _ *Node) (float64, float64) {
	return -1, -1
}

func partialMult(lvalue, rvalue *Node) (float64, float64) {
	return rvalue.value, lvalue.value
}

func partialDiv(lvalue, rvalue *Node) (float64, float64) {
	return math.Pow(rvalue.value, -1), -math.Pow(lvalue.value, -2) * rvalue.value
}

func valueAdd(lvalue, rvalue float64) float64 {
	return lvalue + rvalue
}

func valueSub(lvalue, rvalue float64) float64 {
	return -lvalue - rvalue
}

func valueMult(lvalue, rvalue float64) float64 {
	return lvalue * rvalue
}

func valueDiv(lvalue, rvalue float64) float64 {
	return 1 / (lvalue * rvalue)
}

// func valueAdd(lvalue float64, rvalue float64) float64 {
// 	return
// }

// func opNeg(lvalue float64, rvalue float64) (float64, float64)  {

// }
