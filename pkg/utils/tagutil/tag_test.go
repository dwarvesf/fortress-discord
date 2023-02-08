package tagutil

import "testing"

func Test_formatRole(t *testing.T) {
	type tc struct {
		name string
		role string
		want string
	}

	cases := []tc{{
		name: "valid format",
		role: "123",
		want: "<@&123>",
	}}

	for i := range cases {
		c := cases[i]

		t.Run(c.name, func(t *testing.T) {
			got := FormatRole(c.role)
			if got != c.want {
				t.Errorf("want %v, got %v", c.want, got)
			}
		})
	}
}

func Test_formatUser(t *testing.T) {
	type tc struct {
		name string
		user string
		want string
	}

	cases := []tc{{
		name: "valid format",
		user: "123",
		want: "<@123>",
	}}

	for i := range cases {
		c := cases[i]

		t.Run(c.name, func(t *testing.T) {
			got := FormatRole(c.user)
			if got != c.want {
				t.Errorf("want %v, got %v", c.want, got)
			}
		})
	}
}
