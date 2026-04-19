<img src=".assets/doorman.jpg" alt="Electronics photo" style="width: 400px;" align="right">

# `doorman`

Doorman is a CLI tool for managing secrets. It discovers `doorman-*` plugins on PATH and delegates secret retrieval to the appropriate plugin.

## Plugins

Plugins are standalone CLI binaries following the UNIX plugin protocol. Each plugin handles a specific secret store:

| Plugin | Secret Store |
|--------|-------------|
| `doorman-keepassxc` | KeePassXC password database |
| `doorman-vault` | HashiCorp Vault |
| `doorman-bitwarden` | Bitwarden/Vaultwarden |

## Usage

```bash
# List discovered plugins
doorman plugins

# Get a secret using a specific plugin
doorman get db/password --from vault

# Get a secret (auto-detect plugin from key format)
doorman get secret/data/myapp/password
```

## License

The license for the code and documentation can be found in the [LICENSE](./LICENSE) file.

---

Made in Québec, Canada!
