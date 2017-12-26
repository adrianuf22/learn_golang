package operadora

import (
	"strconv"

	"github.com/adrianuf22/learn_golang/pacotes/prefixo"
)

//NomeOperadora é o nome da operadora
var NomeOperadora = "XPTO Telecom"

//PrefixoDaCapitalOperadora prefixo da capital e nome da operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora

// Using a "private" or not exported variable from another scope
//OperadoraComPrefixoTest nome da operadora e prefixo test
// @v5 var OperadoraComPrefixoTest = NomeOperadora + " " + prefixo.test
// The compiler brokes and returns "cannot refer to unexported name prefixo.test"
