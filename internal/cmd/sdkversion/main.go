// Command sdkversion prints the SDK version constant.
//
// The release workflow uses it to refuse a tag that does not match the version
// the SDK reports about itself at runtime.
package main

import (
	"fmt"

	"github.com/AhaSend/ahasend-go/api"
)

func main() {
	fmt.Println(api.Version)
}
