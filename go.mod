module github.com/AmadlaOrg/doorman

go 1.23.5

replace github.com/AmadlaOrg/LibraryUtils => ../LibraryUtils

replace github.com/AmadlaOrg/LibraryFramework => ../LibraryFramework

require (
	github.com/AmadlaOrg/LibraryFramework v0.0.0
	github.com/AmadlaOrg/LibraryUtils v0.0.0
	github.com/spf13/cobra v1.8.1
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
)
