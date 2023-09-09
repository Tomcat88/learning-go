package reflection

import (
	"reflect"
	"testing"
)

type Profile struct {
	City string
	Age  int
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name:          "struct with one field",
			Input:         struct{ Name string }{"Thomas"},
			ExpectedCalls: []string{"Thomas"},
		},
		{
			Name: "struct with two field",
			Input: struct {
				Name        string
				Description string
			}{"Thomas", "Person"},
			ExpectedCalls: []string{"Thomas", "Person"},
		},
		{
			Name: "struct with different types",
			Input: struct {
				Name        string
				Description int
			}{"Thomas", 35},
			ExpectedCalls: []string{"Thomas"},
		},
		{
			Name: "nested struct",
			Input: struct {
				Name    string
				Profile Profile
			}{"Thomas", Profile{"Milano", 35}},
			ExpectedCalls: []string{"Thomas", "Milano"},
		},
		{
			Name: "pointer",
			Input: &struct {
				Name    string
				Profile Profile
			}{"Thomas", Profile{"Milano", 35}},
			ExpectedCalls: []string{"Thomas", "Milano"},
		},
		{
			Name: "slices",
			Input: []struct {
				Name    string
				Profile Profile
			}{{"Thomas", Profile{"Milano", 35}}},
			ExpectedCalls: []string{"Thomas", "Milano"},
		},
		{
			Name: "slices",
			Input: [2]struct {
				Name    string
				Profile Profile
			}{{"Thomas", Profile{"Milano", 35}}, {"Jimmy", Profile{"New York", 35}}},
			ExpectedCalls: []string{"Thomas", "Milano", "Jimmy", "New York"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
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
		aMap := map[string]string{"Cow": "Moo", "Cat": "Meow"}
		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Meow")
	})

	t.Run("with chans", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{"T", 1}
			aChannel <- Profile{"J", 2}
			close(aChannel)
		}()

		var got []string
		want := []string{"T", "J"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("with fns", func(t *testing.T) {
		aFn := func () (Profile, Profile) {
			return Profile{"A", 1}, Profile{"B", 2}
		}

		var got []string
		want := []string{"A", "B"}

		walk(aFn, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, check string) {
	t.Helper()
	for _, v := range haystack {
		if v == check {
			return
		}
	}

	t.Errorf("haystack %v does not contain %s", haystack, check)
}
