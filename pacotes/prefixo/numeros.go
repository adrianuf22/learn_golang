package prefixo

import (
	// @v5 Cycle import
	// "github.com/adrianuf22/learn_golang/pacotes/operadora"
)

//Capital prefixo da capital de Sao Paulo
var Capital int = 11

// NomeOperadoraPrefixoInterior nome da operadora e prefixo do interior
// @v5 var NomeOperadoraPrefixoInterior = operadora.NomeOperadora + " - 13"
// The code above does not compiles. => import cycle not allowed
// The current package already been imported in operadora package and now,
// the package was imported inside a package that already has them
// This is not allowed by Go compiler
