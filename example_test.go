package embed_test

import (
	"fmt"
	"github.com/cratonica/embed"
	"os"
)

func Example() {
	// Note: This serialization part is implemented in a tool that can be found at http://github.com/cratonica/embedder

	// Create a resource map from files in a directory
	resourceMap, _ := embed.CreateFromFiles("/www/js")

	// Pack the files into a serializeable byte buffer
	packed, _ := embed.Pack(resourceMap)

	// Generate a .go file that we can include in our project
	goCode := embed.GenerateGoCode("main", "Scripts", packed)
	fout, _ := os.Open("scripts.go")
	fout.Write([]byte(goCode))
	fout.Close()

	// ...

	// Now in our consumer program, we have a variable called "Scripts"
	var Scripts embed.PackedResourceMap

	// Unpack the resource map so we can use it
	scriptMap, _ := embed.Unpack(Scripts)

	// Pull out the resource we want
	jquery := scriptMap["jquery/jquery.min.js"]

	// Use the embedded file data
	fmt.Print(jquery)
}
