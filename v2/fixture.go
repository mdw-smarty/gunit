package gunit

import (
	"log"
	"runtime/debug"
	"testing"
)

type Fixture struct {
	config *config
	TestingT
	Logger
}

type Logger interface {
	Printf(string, ...any)
}

// So is a convenience method for reporting assertion failure messages
// with the many assertion functions found in github.com/smarty/gunit/v2/should.
// Example: this.So(actual, should.Equal, expected)
func (this *Fixture) So(actual any, assert Assertion, expected ...any) {
	So(this, actual, assert, expected...)
}

// Run is analogous to *testing.T.Run and allows for running subtests from
// test fixture methods (such as for table-driven tests).
func (this *Fixture) Run(name string, test func(fixture *Fixture)) {
	this.TestingT.(*testing.T).Run(name, func(t *testing.T) {
		t.Helper()
		fixture := &Fixture{
			config:   this.config,
			TestingT: t,
			Logger:   log.New(t.Output(), this.config.logPrefix, this.config.logFlags),
		}
		defer func() {
			if r := recover(); r != nil {
				fixture.Fail()
				fixture.Log(panicReport(r, debug.Stack()))
			}
		}()
		test(fixture)
	})
}

type TestingT interface {
	Cleanup(func())
	Error(args ...any)
	Errorf(format string, args ...any)
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Helper()
	Log(args ...any)
	Logf(format string, args ...any)
	Name() string
	Setenv(key, value string)
	Skip(args ...any)
	SkipNow()
	Skipf(format string, args ...any)
	Skipped() bool
	TempDir() string
}
