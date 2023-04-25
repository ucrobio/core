package spec

import "fmt"

func Test(title string, lines ...Fallible) Engine {
	return (TestEngine{}).Initialize(title, lines)
}

func It(title string, lines ...Fallible) Engine {
	return Test(title, lines...)
}

func TestInline(line Fallible) Engine {
	return It("inline", line)
}

func Inline(line Fallible) Engine {
	return TestInline(line)
}

func BeforeAll(lines ...Callable) Engine {
	return (BeforeAllHookEngine{}).Initialize(lines)
}

func BeforeEach(lines ...Callable) Engine {
	return (BeforeEachHookEngine{}).Initialize(lines)
}

func DeferEach(lines ...Callable) Engine {
	return (DeferEachHookEngine{}).Initialize(lines)
}

func DeferAll(lines ...Callable) Engine {
	return (DeferAllHookEngine{}).Initialize(lines)
}

func Describe(title string, engines ...Engine) Engine {
	return (ContextEngine{}).Initialize(title, engines)
}

func Context(title string, engines ...Engine) Engine {
	return Describe(title, engines...)
}

func When(title string, engines ...Engine) Engine {
	return Describe(title, engines...)
}

func Handle(name string, handler Handler, engines ...Engine) {
	Describe(name, engines...).Context().Handle("handle", handler)
}

func Expect[T any](expected T, matcher NeutralMatcher[T]) error {
	if matcher.Match(expected) {
		return nil
	}

	return matcher.Error(expected)
}

func To[T any](matcher PositiveMatcher[T]) NeutralMatcher[T] {
	return &PositiveMatcherCast[T]{matcher}
}

func NotTo[T any](matcher NegativeMatcher[T]) NeutralMatcher[T] {
	return &NegativeMatcherCast[T]{matcher}
}

func BeEqual[T comparable](value T) *BeEqualMatcher[T] { return &BeEqualMatcher[T]{value} }

type (
	Engine interface {
		IsHook() bool
		Hook() HookEngine
		IsTest() bool
		Test() TestEngine
		IsContext() bool
		Context() ContextEngine
	}

	HookEngine interface {
		IsBeforeAllHookEngine() bool
		IsBeforeEachHookEngine() bool
		IsDeferEachHookEngine() bool
		IsDeferAllHookEngine() bool

		Handle(title string, handler Handler)
	}
)

type BaseEngine struct {
}

func (BaseEngine) IsTest() bool {
	return false
}

func (BaseEngine) Test() TestEngine {
	return TestEngine{}
}

func (BaseEngine) IsHook() bool {
	return false
}

func (BaseEngine) Hook() HookEngine {
	return nil
}

func (BaseEngine) IsContext() bool {
	return false
}

func (BaseEngine) Context() ContextEngine {
	return ContextEngine{}
}

type TestEngine struct {
	BaseEngine

	Title string
	Lines []Fallible
}

func (TestEngine) Initialize(title string, lines []Fallible) (ret TestEngine) {
	ret.Title = title
	ret.Lines = lines

	return ret
}

func (TestEngine) IsTest() bool {
	return true
}

func (it TestEngine) Test() TestEngine {
	return it
}

func (it TestEngine) Handle(title string, handler Handler, hooks ...HookEngine) {
	for _, hook := range hooks {
		switch {
		case hook.IsBeforeEachHookEngine():
			hook.Handle(fmt.Sprintf("%s: %s", title, it.Title), handler)
		case hook.IsDeferEachHookEngine():
			defer hook.Handle(fmt.Sprintf("%s: %s", title, it.Title), handler)
		}
	}

	for _, line := range it.Lines {
		if handler.Call(enriches(title, line.Call())) {
			return
		}
	}
}

type BaseHookEngine struct {
	BaseEngine

	Lines []Callable
}

func (BaseHookEngine) Initialize(lines []Callable) (ret BaseHookEngine) {
	ret.Lines = lines

	return ret
}

func (BaseHookEngine) IsBeforeAllHookEngine() bool {
	return false
}

func (BaseHookEngine) IsBeforeEachHookEngine() bool {
	return false
}

func (BaseHookEngine) IsDeferEachHookEngine() bool {
	return false
}

func (BaseHookEngine) IsDeferAllHookEngine() bool {
	return false
}

func (it BaseHookEngine) Handle(title string, handler Handler) {
	for _, line := range it.Lines {
		handler.Call(enriches(title, line.Call()))
	}
}

type BeforeAllHookEngine struct {
	BaseHookEngine
}

func (BeforeAllHookEngine) Initialize(lines []Callable) (ret BeforeAllHookEngine) {
	ret.BaseHookEngine = ret.BaseHookEngine.Initialize(lines)

	return ret
}

func (BeforeAllHookEngine) IsBeforeAllHookEngine() bool {
	return true
}

func (BeforeAllHookEngine) IsHook() bool {
	return true
}

func (it BeforeAllHookEngine) Hook() HookEngine {
	return it
}

func (it BeforeAllHookEngine) Handle(title string, handler Handler) {
	it.BaseHookEngine.Handle(fmt.Sprintf("%s: before all", title), handler)
}

type BeforeEachHookEngine struct {
	BaseHookEngine
}

func (BeforeEachHookEngine) Initialize(lines []Callable) (ret BeforeEachHookEngine) {
	ret.BaseHookEngine = ret.BaseHookEngine.Initialize(lines)

	return ret
}

func (BeforeEachHookEngine) IsBeforeEachHookEngine() bool {
	return true
}

func (BeforeEachHookEngine) IsHook() bool {
	return true
}

func (it BeforeEachHookEngine) Hook() HookEngine {
	return it
}

func (it BeforeEachHookEngine) Handle(title string, handler Handler) {
	it.BaseHookEngine.Handle(fmt.Sprintf("%s: before each", title), handler)
}

type DeferEachHookEngine struct {
	BaseHookEngine
}

func (DeferEachHookEngine) Initialize(lines []Callable) (ret DeferEachHookEngine) {
	ret.BaseHookEngine = ret.BaseHookEngine.Initialize(lines)

	return ret
}

func (DeferEachHookEngine) IsDeferEachHookEngine() bool {
	return true
}

func (DeferEachHookEngine) IsHook() bool {
	return true
}

func (it DeferEachHookEngine) Hook() HookEngine {
	return it
}

func (it DeferEachHookEngine) Handle(title string, handler Handler) {
	it.BaseHookEngine.Handle(fmt.Sprintf("%s: defer each", title), handler)
}

type DeferAllHookEngine struct {
	BaseHookEngine
}

func (DeferAllHookEngine) Initialize(lines []Callable) (ret DeferAllHookEngine) {
	ret.BaseHookEngine = ret.BaseHookEngine.Initialize(lines)

	return ret
}

func (DeferAllHookEngine) IsDeferAllHookEngine() bool {
	return true
}

func (DeferAllHookEngine) IsHook() bool {
	return true
}

func (it DeferAllHookEngine) Hook() HookEngine {
	return it
}

func (it DeferAllHookEngine) Handle(title string, handler Handler) {
	it.BaseHookEngine.Handle(fmt.Sprintf("%s: defer all", title), handler)
}

type ContextEngine struct {
	BaseEngine

	Title string

	Hooks    []HookEngine
	Tests    []TestEngine
	Contexts []ContextEngine
}

func (ContextEngine) Initialize(title string, engines []Engine) (ret ContextEngine) {
	ret.Title = title

	for _, engine := range engines {
		switch {
		case engine.IsContext():
			ret.Contexts = append(ret.Contexts, engine.Context())
		case engine.IsHook():
			ret.Hooks = append(ret.Hooks, engine.Hook())
		case engine.IsTest():
			ret.Tests = append(ret.Tests, engine.Test())
		}
	}

	return ret
}

func (ContextEngine) IsContext() bool {
	return true
}

func (it ContextEngine) Context() ContextEngine {
	return it
}

func (it ContextEngine) Handle(title string, handler Handler, hooks ...HookEngine) {
	for _, hook := range it.Hooks {
		switch {
		case hook.IsBeforeAllHookEngine():
			hook.Handle(it.parentTitle(title), handler)
		case hook.IsDeferAllHookEngine():
			defer hook.Handle(it.parentTitle(title), handler)
		}
	}

	for _, test := range it.Tests {
		test.Handle(it.parentTitle(title), handler, append(hooks, it.Hooks...)...)
	}

	for _, context := range it.Contexts {
		context.Handle(it.parentTitle(title), handler, append(hooks, it.Hooks...)...)
	}
}

func (it ContextEngine) parentTitle(parentTitle string) string {
	return fmt.Sprintf("%s: %s", parentTitle, it.Title)
}

type Handler func(err error)

func (it Handler) Call(err error) bool {
	if err != nil {
		it(err)
		return true
	}

	return false
}

type Callable func()

func (it Callable) Call() (err error) {
	defer func() {
		rec := recover()

		if rec != nil {
			err = fmt.Errorf("recovered from: %v", rec)
		}
	}()

	it()

	return
}

type Fallible func() error

func (it Fallible) Call() (err error) {
	defer func() {
		rec := recover()

		if rec != nil {
			err = fmt.Errorf("recovered from: %v", rec)
		}
	}()

	return it()
}

func enriches(title string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", title, err)
}

type PositiveMatcher[T any] interface {
	PositiveMatch(expected T) bool
	PositiveError(expected T) error
}

type NegativeMatcher[T any] interface {
	NegativeMatch(expected T) bool
	NegativeError(expected T) error
}

type Matcher[T any] interface {
	PositiveMatcher[T]
	NegativeMatcher[T]
}

type NeutralMatcher[T any] interface {
	Match(expected T) bool
	Error(expected T) error
}

type PositiveMatcherCast[T any] struct {
	PositiveMatcher[T]
}

func (it PositiveMatcherCast[T]) Match(expected T) bool {
	return it.PositiveMatch(expected)
}

func (it PositiveMatcherCast[T]) Error(expected T) error {
	return it.PositiveError(expected)
}

type NegativeMatcherCast[T any] struct {
	NegativeMatcher[T]
}

func (it NegativeMatcherCast[T]) Match(expected T) bool {
	return it.NegativeMatch(expected)
}

func (it NegativeMatcherCast[T]) Error(expected T) error {
	return it.NegativeError(expected)
}

type BeEqualMatcher[T comparable] struct {
	Value T
}

func (it BeEqualMatcher[T]) PositiveMatch(expected T) bool {
	return expected == it.Value
}

func (it BeEqualMatcher[T]) PositiveError(expected T) error {
	return fmt.Errorf("should be equal: %v != %v", expected, it.Value)
}

func (it BeEqualMatcher[T]) NegativeMatch(expected T) bool {
	return expected != it.Value
}

func (it BeEqualMatcher[T]) NegativeError(expected T) error {
	return fmt.Errorf("shouldn't be equal: %v == %v", expected, it.Value)
}
