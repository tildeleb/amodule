// will not build with this module name github.com/tildeleb/amodule

module amodule

go 1.21.1

replace two v0.0.0 => ./internal/two

require two v0.0.0
