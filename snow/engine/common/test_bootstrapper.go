// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

package common

var (
	_ Engine        = (*BootstrapperTest)(nil)
	_ Bootstrapable = (*BootstrapperTest)(nil)
)

// EngineTest is a test engine
type BootstrapperTest struct {
	BootstrapableTest
	EngineTest
}

func (b *BootstrapperTest) Default(cant bool) {
	b.BootstrapableTest.Default(cant)
	b.EngineTest.Default(cant)
}
