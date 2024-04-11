package cpfp

import (
	"fmt"
	"github.com/tristan-club/kit/log"
	"testing"
)

func init() {
	log.SetConsoleWrite()
}

func TestCPFP(t *testing.T) {
	child, fee, err := CalCpfp("d75b094d90e116b3524465a5d1b233d18e188989f0d78edcc7314fa16d6a5297", 176.25, 300)
	fmt.Println(child)
	fmt.Println(fee)
	fmt.Println(err)
}
