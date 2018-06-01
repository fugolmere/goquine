package main

import quine "fmt"

var q = "package main\n\nimport quine \"fmt\"\n\nvar q = %s\n\nfunc main() { quine.Printf(q, quine.Sprintf(%q, q), string(37)+string(113)) }\n"

func main() { quine.Printf(q, quine.Sprintf("%q", q), string(37)+string(113)) }
