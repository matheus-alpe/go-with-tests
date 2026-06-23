package generics

import "testing"

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := NewStack[int]()

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)

		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		first, _ := myStackOfInts.Pop()
		second, _ := myStackOfInts.Pop()
		AssertEqual(t, first+second, 3)
	})

	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := NewStack[string]()

		AssertTrue(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("123")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("456")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "456")

		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "123")

		AssertTrue(t, myStackOfStrings.IsEmpty())
	})
}
