package parse

import (
	"fmt"
	"strings"
	"testing"
)

func TestOrderPrecedence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"1+2*3;",
			"(1 + (2 * 3))",
		},
		{
			"1+2+3;",
			`((1 + 2) + 3)`,
		},
		{
			"1 * 2 + 3;",
			"((1 * 2) + 3)",
		},
		{
			"a * b + c;",
			`((a * b) + c)`,
		},
		{
			"-a * b;",
			"((-a) * b)",
		},
		{
			"a - b;",
			"(a - b)",
		},
		{
			"!-a;",
			"(!(-a))",
		},
		{
			"(1+2)*3;",
			"((1 + 2) * 3)",
		},
		{
			"1==2 && 3!=4;",
			"((1 == 2) && (3 != 4))",
		},
		{
			"1 && 0 || 1+2 == 3;",
			"((1 && 0) || ((1 + 2) == 3))",
		},
	}

	for _, test := range tests {
		p := New(test.input)
		res, err := p.Parse()
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		// prettyPrint(res)
		actual := res.Program.String()
		// strip \n
		actual = strings.TrimSuffix(actual, "\n")
		if actual != test.expected {
			t.Errorf("expected=%q, got %q", test.expected, actual)
		}
	}
}

// func prettyPrint(body interface{}) {
// 	b, err := json.MarshalIndent(body, "", "\t")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(b))
// }

const cfg = `
	switch a {
	case 1:
		1+1;
		fallthrough;
	case 2:
		2+2;
		break;
	default:
		3+3;
	}
	fn y(a, b) {
		let c = a + b;
		return c;
	}
	let x = 0;
	while x < 10 {
		x = y(x, 1);
		//do loopy stuff
		if x > 0 {
			continue;
		} else {
			break;
		}
	}
	for x = 0; x < 5; x = x + 1 {
		let i = y(x);
	}
`

func TestCfg(t *testing.T) {
	p := New(cfg)
	fmt.Printf("parsing:\n %v\n", cfg)
	res, err := p.Parse()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println("output:")
	fmt.Println(res.Program.String())
}

const fntest = `
fn y(x) {
    print(x);
    return x +1;
}

let z = f(2);

print(z);

print("hi");
print([1, 2, hello="world", 3]);
`

func TestFnCall(t *testing.T) {
	p := New(fntest)
	res, err := p.Parse()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Println("output:")
	fmt.Println(res.Program.String())
}
