package eventchecker

import "github.com/simimpact/srsim/pkg/engine/event/handler"

// EventChecker returns True if the checker passes, and False if the checker fails.
// All checkers must pass for an Expect to conclude.
// If any checker fails, subsequent checkers will not run. If it also returns an error, the Expect fails.
// If it fails but does not return an error, Expect will continue to run.
type EventChecker func(e handler.Event) (bool, error)
