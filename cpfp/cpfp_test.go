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
	//child, err := CalCpfp("d75b094d90e116b3524465a5d1b233d18e188989f0d78edcc7314fa16d6a5297", 176.25, 300)
	child, err := CalCpfp("5a098ff9148d97b669b031b448fe96ac297f9b7307ae4e5d31ac10f8ed96e020", 430.75, 60)
	//child, fee, err := CalCpfp("eb690825e578c69074f12919f8bbd3f05462372189741d88f84187dec0531a82", 0, 9)
	//child, fee, err := CalCpfp("ab6ef6b07507e00903775702ce88ede0343b27d70c1f2236196ec0af3ab71833", 176.5, 9)
	fmt.Println(child)
	//fmt.Println(fee)
	fmt.Println(err)
}
