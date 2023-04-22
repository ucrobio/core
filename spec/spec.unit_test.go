package spec

import "testing"

func TestSpecUnitSuite(t *testing.T) {
	Handle(
		"spec unit suite",

		func(err error) { t.Fatal(err) },

		Test(
			"#Test",

			func() error { return Expect(Test("test name").IsContext(), To[bool](BeEqual(false))) },
			func() error { return Expect(Test("test name").IsHook(), To[bool](BeEqual(false))) },
			func() error { return Expect(Test("test name").IsTest(), To[bool](BeEqual(true))) },
			func() error { return Expect(Test("test name").Test().Title, To[string](BeEqual("test name"))) },
		),
	)
}
