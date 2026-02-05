module github.com/AmadlaOrg/doorman

go 1.24.0

toolchain go1.24.7

replace github.com/AmadlaOrg/LibraryUtils => ../LibraryUtils

replace github.com/AmadlaOrg/LibraryFramework => ../LibraryFramework

require (
	github.com/AmadlaOrg/LibraryFramework v0.0.0
	github.com/dgraph-io/ristretto v0.2.0
	github.com/spf13/cobra v1.9.1
	golang.org/x/sys v0.30.0
)

require (
	github.com/AmadlaOrg/LibraryUtils v0.0.0 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
)
