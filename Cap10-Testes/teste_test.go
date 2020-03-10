package teste

import (
	"fmt"
	"testing"
)

func TestExtrairSite(t *testing.T) {
	site := extrairSite("https://github.com/amaralfelipe1522")
	fmt.Println(<-site)
}
