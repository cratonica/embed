embed
=====

Package embed allows for storing data resources in a virtual filesystem that gets compiled directly into the output program, eliminating the need to distribute data files with the application. This is especially useful for small web servers that need to deliver content files.

The data is gzipped to save space.

An external tool for generating output files can be found [here](http://github.com/cratonica/embedder)

View the full documentation [here](http://godoc.org/github.com/cratonica/embed)

Example
-------
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

