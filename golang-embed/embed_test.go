package golang_embed

import (
	_ "embed"
	"fmt"
	"testing"
)

//embed file txt kedalam variable string
//hanya bisa di lakukan di luar function

//go:embed version.txt
var version string

func TestEmbedString(t *testing.T) {
	fmt.Println(version)
}
