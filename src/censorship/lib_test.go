package censorship

import (
  "fmt"
	"reflect"
	"testing"
)

func ObjectsAreEqual(expected, actual interface{}) bool {

	if expected == nil || actual == nil {
		return expected == actual
	}

	if reflect.DeepEqual(expected, actual) {
		return true
	}

	return false

}

// ObjectsAreEqualValues gets whether two objects are equal, or if their
// values are equal.
func ObjectsAreEqualValues(expected, actual interface{}) bool {
	if ObjectsAreEqual(expected, actual) {
		return true
	}

	actualType := reflect.TypeOf(actual)
	expectedValue := reflect.ValueOf(expected)
	if expectedValue.Type().ConvertibleTo(actualType) {
		// Attempt comparison after type conversion
		if reflect.DeepEqual(actual, expectedValue.Convert(actualType).Interface()) {
			return true
		}
	}

	return false
}

func TestCensorshipA(t *testing.T) {
	t.Parallel()

	actual := Censorship("You are a nice person")
	excepted := "You are a XXXX person"

	if !ObjectsAreEqual(excepted, actual) {
		t.Errorf("Censorship not working")
		t.Fail()
	}
}

func TestCensorshipB(t *testing.T) {
  t.Parallel()

  actual := Censorship("Such a nice day with a bright sun, makes me happy")
  excepted := "Such a XXXX day with a bright XXX, makes me XXXXX"

  if !ObjectsAreEqual(excepted, actual) {
    t.Errorf("Censorship not working on multiple words")
    t.Fail()
  }
}

func TestCensorshipC(t *testing.T) {
  t.Parallel()

  actual := Censorship("You are so friendly!")
  excepted := "You are so XXXXXXXXX"

  if !ObjectsAreEqual(excepted, actual) {
    t.Errorf(fmt.Sprintf("Censorship not working on compost words; actual: %s", actual))
    t.Fail()
  }
}

func TestCensorshipD(t *testing.T) {
  t.Parallel()

  actual := Censorship("Objection is bad, a better thing to do, is to agree.")
  var excepted string
  excepted = "Thoughtcrime is ungood, a gooder thing to do, is to crimestop."

  if !ObjectsAreEqual(excepted, actual) {
    t.Errorf(fmt.Sprintf("Censorship not working on replacement of words; actual: %s", actual))
    t.Fail()
  }

  actual = Replace("Objection is bad, a better thing to do, is to agree.")
  excepted = "Thoughtcrime is ungood, a gooder thing to do, is to crimestop."

  if !ObjectsAreEqual(excepted, actual) {
    t.Errorf(fmt.Sprintf("Replace not working on replacement of words; actual: %s", actual))
    t.Fail()
  }

}

