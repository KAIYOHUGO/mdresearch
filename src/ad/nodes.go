package ad

// is (left side value,right side value)(left side partial,right side partial)
type calcPartial func(*Node, *Node) (float64, float64)

// is (left side value,right side value)(this value)
type calcValue func(float64, float64) float64

type Node struct {
	// is (left side value,right side value)(left side partial,right side partial)
	calcPartial calcPartial
	// is (left side value,right side value)(this value)
	calcValue calcValue
	// diff
	partial float64
	// calc number & const number value
	value        float64
	rSide, lSide *Side
}

type Side struct {
	// child node
	node *Node
	// diff
	partial float64
}

func Varible(value float64) *Node {
	return &Node{
		value: value,
	}
}

func (s *Node) Calc() {
	s.forward()
	s.partial = 1
	s.backward()
}

func (s *Node) Set(value float64) *Node {
	s.value = value
	return s
}

func (s *Node) Partial() float64 {
	return s.partial
}

func (s *Node) Value() float64 {
	return s.value
}

func (s *Node) forward() float64 {
	if s.calcValue != nil {
		if s.lSide != nil {
			s.value = s.calcValue(s.lSide.node.forward(), s.rSide.node.forward())
		} else if s.rSide != nil {
			s.value = s.calcValue(0, s.rSide.node.forward())
		}
	}
	if s.calcPartial != nil {
		if s.lSide != nil {
			s.lSide.partial, s.rSide.partial = s.calcPartial(s.lSide.node, s.rSide.node)
		} else if s.rSide != nil {
			_, s.rSide.partial = s.calcPartial(nil, s.rSide.node)
		}
	}
	return s.value
}

func (s *Node) backward() {
	if s.rSide != nil {
		s.rSide.node.partial += s.partial * s.rSide.partial
		s.rSide.node.backward()
	}
	if s.lSide != nil {
		s.lSide.node.partial += s.partial * s.lSide.partial
		s.lSide.node.backward()
	}
}

func newNode(calcValue calcValue, calcPartial calcPartial) *Node {
	return &Node{
		calcPartial: calcPartial,
		calcValue:   calcValue,
	}
}

func newSide(value *Node) *Side {
	return &Side{
		node: value,
	}
}

// a \{1}    {b}
// 	 |--- a+b -- (a+b)*b
// b /----------/
//   {1}      {a+b}

// x^-1*y^-1
// if x
// -x^-2*y^-1
