package spec

import "testing"

func TestSpecUnitSuite(t *testing.T) {
	Handle(
		"spec unit suite",

		func(err error) { t.Fatal(err) },

		Describe(
			"#Test",

			Inline(func() error { return Expect(Test("test name").IsContext(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(Test("test name").IsHook(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(Test("test name").IsTest(), To[bool](BeEqual(true))) }),
			Inline(func() error { return Expect(Test("test name").Test().Title, To[string](BeEqual("test name"))) }),
		),

		Describe(
			"#It",

			Inline(func() error { return Expect(It("test name").IsContext(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(It("test name").IsHook(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(It("test name").IsTest(), To[bool](BeEqual(true))) }),
			Inline(func() error { return Expect(It("test name").Test().Title, To[string](BeEqual("test name"))) }),
		),

		Describe(
			"#Inline",

			Inline(func() error { return Expect(Inline(func() error { return nil }).IsContext(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(Inline(func() error { return nil }).IsHook(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(Inline(func() error { return nil }).IsTest(), To[bool](BeEqual(true))) }),
			Inline(func() error {
				return Expect(Inline(func() error { return nil }).Test().Title, To[string](BeEqual("inline")))
			}),
		),
	)
}
