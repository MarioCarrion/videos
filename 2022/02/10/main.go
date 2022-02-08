package main

import (
	"flag"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:generate goi18n extract -sourceLanguage en

func main() {
	var count int
	var lang string

	flag.IntVar(&count, "count", 0, "number of items to buy")
	flag.StringVar(&lang, "lang", "en", "language to use")

	flag.Parse()

	//-

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("active.en.toml")
	bundle.LoadMessageFile("active.es.toml")

	localizer := i18n.NewLocalizer(bundle, lang)

	buying := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "BuyingCookies",
			One:   "You're buying 1 cookie.",
			Other: "You're buying {{.PluralCount}} cookies.",
		},
		PluralCount: count,
	})

	// buying := "" // "You're buying %d cookies\n", count

	fmt.Printf("%s\n", buying)
}
