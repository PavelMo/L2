package main

import "testing"

func TestUnpackWithoutSlash(t *testing.T) {
	cases := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		"qwe10qwe",
		"a1b2c1",
	}

	results := []string{
		"aaaabccddddde",
		"abcd",
		"",
		"",
		"qweeeeeeeeeeqwe",
		"abbc",
	}

	for i := range cases {
		unpacked, _ := Unpack(cases[i])
		if unpacked != results[i] {
			t.Errorf("%s != %s\n", results[i], unpacked)
		}
	}
}

func TestUnpackWithSlash(t *testing.T) {
	cases := []string{
		`qwe\4\5`,
		`qwe\45`,
		`qwe\\5`,
		`qwe\`,
		`qwe4\4`,
		`qwe\44`,
		`\`,
		`\\`,
		`\4`,
		`4\5`,
		`\45`,
	}

	results := []string{
		"qwe45",
		"qwe44444",
		`qwe\\\\\`,
		"qwe",
		"qweeee4",
		"qwe4444",
		"",
		`\`,
		"4",
		"",
		"44444",
	}

	for i := range cases {
		unpacked, _ := Unpack(cases[i])
		if unpacked != results[i] {
			t.Errorf("%s != %s\n", results[i], unpacked)
		}
	}
}
