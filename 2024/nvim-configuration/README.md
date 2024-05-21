# Neovim configuration for Go development using lazy.nvim and go.nvim

* [YouTube Video](https://youtu.be/n5_WLgxwkU8)
* [Blog post](https://mariocarrion.com/2024/05/20/neovim-migrating-to-lazy-and-go-nvim.html)

❗️❗️This configuration works with Neovim **v0.10.0**❗️❗️

## Requirements

* Install a patched font, my favorite is [JetBrains patched font](https://github.com/ryanoasis/nerd-fonts/tree/master/patched-fonts/JetBrainsMono).

If you're using [iTerm](https://www.iterm2.com/downloads.html), I recommend the following preferences:

* General -> Working Directory -> Advanced Configuration -> Working Directory for New Tabs -> Reuse previous session's directory
* Text -> Font -> JetBrainsMono Nerd Font
* Text -> Font -> Use ligatures
* Text -> Font -> Anti-aliased
* Keys -> General -> Left Option key -> Esc+
* Keys -> Key Mappings -> Presets -> Natural Text Editing

## Install

You can copy the content of this folder to your Neovim config path or create a
symbolic link, to know your Neovim config path use the following command:

```
:echo stdpath("config")
```

[lazy.nvim](https://github.com/folke/lazy.nvim) is the plugin manager, use `:Lazy` to install all plugins.

### Dependencies

Make sure the following are installed:

* Python 3: `brew install python`. Make sure is a recipe for python3.x, and then `pip3 install neovim`.
* `fd` alternative to `find`: `brew install fd`. Used by "telescope" plugins.

## Tricks

* Beautify JSON buffer: `:%!jq`

## Language Server Protocol Implementations

* [Go: gopls](https://github.com/golang/tools/tree/master/gopls)
```
go install golang.org/x/tools/gopls@latest
```

* [Python](https://github.com/microsoft/pyright)

```
brew install pyright
```

* [Ruby](https://solargraph.org/)

```
gem install solargraph
```

* [PHP](https://github.com/phpactor/phpactor)

```
brew install php
```
