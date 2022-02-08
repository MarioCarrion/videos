# Localizability (i18n + l10n) using goi18n

[Video](https://youtu.be/l4RmhSbgtfQ)

1. Import the following packages:

```
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
```
1. Install `goi18n` tool: `go install github.com/nicksnyder/go-i18n/v2/goi18n@latest`
1. Add the following `//go:generate goi18n extract -sourceLanguage en`
1. Create an empty file to use for translation, for example `active.es.toml`, using `touch active.es.toml`
1. Run `goi18n merge active.en.toml active.es.toml`

