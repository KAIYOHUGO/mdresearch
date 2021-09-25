package ad

import "testing"

func TestAD(t *testing.T) {
	a := Varible(2)
	b := Varible(3)
	c := Add(a, b)
	c.Calc()
	var dfs func(*Node, int)
	dfs = func(n *Node, d int) {
		t.Log(d, n.value, n.partial)
		if n.rSide != nil {
			dfs(n.rSide.node, d+1)
		}
		if n.lSide != nil {
			dfs(n.lSide.node, d+1)
		}
	}
	dfs(c, 0)
	t.Log(a.Partial(), b.Partial())
	if a.Partial() != 1 || b.Partial() != 1 {
		t.Fail()
	}
}

func TestAD1(t *testing.T) {
	a := Varible(2)
	b := Varible(3)
	c := Mult(a, b)
	c.Calc()
	var dfs func(*Node, int)
	dfs = func(n *Node, d int) {
		t.Log(d, n.value, n.partial)
		if n.rSide != nil {
			dfs(n.rSide.node, d+1)
		}
		if n.lSide != nil {
			dfs(n.lSide.node, d+1)
		}
	}
	dfs(c, 0)
	t.Log(a.Partial(), b.Partial())
	if a.Partial() != 3 || b.Partial() != 2 {
		t.Fail()
	}
}

func TestAD0and1(t *testing.T) {
	a := Varible(2)
	b := Varible(3)
	c := Add(Mult(a, b), a)
	c.Calc()
	var dfs func(*Node, int)
	dfs = func(n *Node, d int) {
		t.Log(d, n.value, n.partial)
		if n.rSide != nil {
			dfs(n.rSide.node, d+1)
		}
		if n.lSide != nil {
			dfs(n.lSide.node, d+1)
		}
	}
	dfs(c, 0)
	t.Log(a.Partial(), b.Partial())
	if a.Partial() != 4 || b.Partial() != 2 {
		t.Fail()
	}
}

// b \{a}(2)  {1}
// 	 | a*b ----- a*b+a {1}
// a /----------/
//   {b}(3)   {1}
// 2,3

func TestAD2(t *testing.T) {
	a := Varible(2)
	b := Varible(3)
	c := Mult(a, b.Neg())
	c.Calc()
	var dfs func(*Node, int)
	dfs = func(n *Node, d int) {
		t.Log(d, n.value, n.partial)
		if n.rSide != nil {
			dfs(n.rSide.node, d+1)
		}
		if n.lSide != nil {
			dfs(n.lSide.node, d+1)
		}
	}
	dfs(c, 0)
	t.Log(a.Partial(), b.Partial())
	if a.Partial() != -3 || b.Partial() != -2 {
		t.Fail()
	}
}

// -x
// if 2
// -1

func TestGmatSelfAdd(t *testing.T) {
	a := Varible(2)
	b := Varible(3)
	c := a.Add(b)
	c.Calc()
	var dfs func(*Node, int)
	dfs = func(n *Node, d int) {
		t.Log(d, n.value, n.partial)
		if n.rSide != nil {
			dfs(n.rSide.node, d+1)
		}
		if n.lSide != nil {
			dfs(n.lSide.node, d+1)
		}
	}
	dfs(c, 0)
	t.Log(a.Partial(), b.Partial())
	if a.Partial() != 1 || b.Partial() != 1 {
		t.Fail()
	}
}

func TestADSelfMult(t *testing.T) {
	a := Varible(2)
	b := Varible(3)
	c := a.Mult(b)
	c.Calc()
	var dfs func(*Node, int)
	dfs = func(n *Node, d int) {
		t.Log(d, n.value, n.partial)
		if n.rSide != nil {
			dfs(n.rSide.node, d+1)
		}
		if n.lSide != nil {
			dfs(n.lSide.node, d+1)
		}
	}
	dfs(c, 0)
	t.Log(a.Partial(), b.Partial())
	if a.Partial() != 3 || b.Partial() != 2 {
		t.Fail()
	}
}

func TestADSelfDiv(t *testing.T) {
	a := Varible(2)
	b := Varible(3)
	c := a.Div(b)
	c.Calc()
	var dfs func(*Node, int)
	dfs = func(n *Node, d int) {
		t.Log(d, n.value, n.partial)
		if n.rSide != nil {
			dfs(n.rSide.node, d+1)
		}
		if n.lSide != nil {
			dfs(n.lSide.node, d+1)
		}
	}
	dfs(c, 0)
	t.Log(a.Partial(), b.Partial())
	if a.Partial() != 3 || b.Partial() != 2 {
		t.Fail()
	}
}
