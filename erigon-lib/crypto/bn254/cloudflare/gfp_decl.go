//go:build (amd64 && !generic) || (arm64 && !generic)

package bn254

// This file contains forward declarations for the architecture-specific
// assembly implementations of these functions, provided that they exist.

import (
	"golang.org/x/sys/cpu"
)

var hasBMI2 = cpu.X86.HasBMI2

//go:noescape
func gfpNeg(c, a *gfP)

//go:noescape
func gfpAdd(c, a, b *gfP)

//go:noescape
func gfpSub(c, a, b *gfP)

//go:noescape
func gfpMul(c, a, b *gfP)
