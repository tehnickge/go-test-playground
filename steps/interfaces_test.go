package steps

import "testing"

func testInterfaceImplementation(t *testing.T) {
	table := []struct {
		valid    bool
		instance Item
		expected string
	}{
		{true, Gun(10), "gun: 10"},
		{true, AnotherType{}, "AnotherType"},
		{false, Gun(20), "gun: 200"},
	}
	for _, testCase := range table {

		if testCase.valid {
			if testCase.instance.WhoIAm() != testCase.expected {
				t.Errorf("Expected %s, got %s", testCase.expected, testCase.instance.WhoIAm())
			}

		} else {
			if testCase.instance.WhoIAm() == testCase.expected {
				t.Errorf("Expected %s to not match, but got %s", testCase.expected, testCase.instance.WhoIAm())
			}
		}

	}

}

func TestChaneNameForDog(t *testing.T) {
	table := []struct {
		valid    bool
		player   Player
		expected string
	}{
		{true, &Dog{name: "dog"}, "testName"},
		{false, &Cat{name: "cat"}, "cat"},
	}

	for _, testCase := range table {
		chaneNameForDog(testCase.player)
		if testCase.valid {
			if testCase.player.GetName() != testCase.expected {
				t.Errorf("Expected %s to not match, but got %s", testCase.player.GetName(), testCase.expected)
			}

		} else {
			if testCase.player.GetName() != testCase.expected {
				t.Errorf("Expected %s to match, but got %s", testCase.player.GetName(), testCase.expected)
			}
		}

	}

}
