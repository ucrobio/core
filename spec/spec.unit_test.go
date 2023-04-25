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

		Describe(
			"#BeforeAll",

			Inline(func() error { return Expect(BeforeAll().IsContext(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(BeforeAll().IsHook(), To[bool](BeEqual(true))) }),
			Inline(func() error { return Expect(BeforeAll().IsTest(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(BeforeAll().Hook().IsBeforeAllHookEngine(), To[bool](BeEqual(true))) }),
		),

		Describe(
			"#BeforeEach",

			Inline(func() error { return Expect(BeforeEach().IsContext(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(BeforeEach().IsHook(), To[bool](BeEqual(true))) }),
			Inline(func() error { return Expect(BeforeEach().IsTest(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(BeforeEach().Hook().IsBeforeEachHookEngine(), To[bool](BeEqual(true))) }),
		),

		Describe(
			"#DeferEach",

			Inline(func() error { return Expect(DeferEach().IsContext(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(DeferEach().IsHook(), To[bool](BeEqual(true))) }),
			Inline(func() error { return Expect(DeferEach().IsTest(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(DeferEach().Hook().IsDeferEachHookEngine(), To[bool](BeEqual(true))) }),
		),

		Describe(
			"#DeferAll",

			Inline(func() error { return Expect(DeferAll().IsContext(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(DeferAll().IsHook(), To[bool](BeEqual(true))) }),
			Inline(func() error { return Expect(DeferAll().IsTest(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(DeferAll().Hook().IsDeferAllHookEngine(), To[bool](BeEqual(true))) }),
		),

		Describe(
			"#Describe",

			Inline(func() error { return Expect(Describe("title").IsContext(), To[bool](BeEqual(true))) }),
			Inline(func() error { return Expect(Describe("title").IsHook(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(Describe("title").IsTest(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(Describe("title").Context().Title, To[string](BeEqual("title"))) }),
		),

		Describe(
			"#Context",

			Inline(func() error { return Expect(Context("title").IsContext(), To[bool](BeEqual(true))) }),
			Inline(func() error { return Expect(Context("title").IsHook(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(Context("title").IsTest(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(Context("title").Context().Title, To[string](BeEqual("title"))) }),
		),

		Describe(
			"#When",

			Inline(func() error { return Expect(When("title").IsContext(), To[bool](BeEqual(true))) }),
			Inline(func() error { return Expect(When("title").IsHook(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(When("title").IsTest(), To[bool](BeEqual(false))) }),
			Inline(func() error { return Expect(When("title").Context().Title, To[string](BeEqual("title"))) }),
		),
	)
}
