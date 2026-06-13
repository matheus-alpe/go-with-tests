package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"Matt"},
			ExpectedCalls: []string{"Matt"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Matt", "Palhoça"},
			ExpectedCalls: []string{"Matt", "Palhoça"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Matt", 30},
			ExpectedCalls: []string{"Matt"},
		},

		{
			Name:          "nested fields",
			Input:         Person{"Matt", Profile{30, "Palhoça"}},
			ExpectedCalls: []string{"Matt", "Palhoça"},
		},
		{
			Name:          "pointers to things",
			Input:         &Person{"Matt", Profile{30, "Palhoça"}},
			ExpectedCalls: []string{"Matt", "Palhoça"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{30, "Palhoça"},
				{31, "Floripa"},
			},
			ExpectedCalls: []string{"Palhoça", "Floripa"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{30, "Palhoça"},
				{31, "Floripa"},
			},
			ExpectedCalls: []string{"Palhoça", "Floripa"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{30, "Palhoça"}
			aChannel <- Profile{31, "Floripa"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Palhoça", "Floripa"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{30, "Palhoça"}, Profile{31, "Floripa"}
		}

		var got []string
		want := []string{"Palhoça", "Floripa"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}

}
