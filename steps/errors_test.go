package steps

import "testing"

func TestErrors(t *testing.T) {
	table := []struct {
		valid    bool
		lhs      int
		rhs      int
		expected float64
	}{
		{true, 10, 10, 1},
		{false, 10, 0, 0},
	}

	for _, testCase := range table {
		res, err := divide(testCase.lhs, testCase.rhs)

		if testCase.valid {
			if err != nil {
				t.Errorf("have error, but testCase is valid. lhs %v rhs %v\n error: %v", testCase.lhs, testCase.rhs, err)
			}
			if testCase.expected != res {
				t.Errorf("bad result lhs %v rhs %v\n error: %v", testCase.lhs, testCase.rhs, err)
			}
		} else {
			if err == nil {
				t.Errorf("error is nil, but test case data is invalid")
			}

		}
	}
}
