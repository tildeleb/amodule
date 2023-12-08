// amodule will build as a program with module name of amodule but not github.com/tildeleb/amodule
// amodule woth build as module ONLY with module name github.com/tildeleb/amodule

module github.com/tildeleb/amodule

go 1.21.1

replace two v0.0.0 => ./internal/two

require two v0.0.0
