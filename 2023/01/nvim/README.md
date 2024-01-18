# neovim

* [JetBrains font](https://github.com/ryanoasis/nerd-fonts/blob/master/patched-fonts/JetBrainsMono/Ligatures/Regular/complete/JetBrains%20Mono%20Regular%20Nerd%20Font%20Complete%20Mono.ttf)

## Install

```
brew install neovim

ln -s $PWD ~/.config/nvim
```

### Dependencies

* Python 3: `brew install python` (make sure is a recipe for python3.x) and then `pip3 install neovim`

## Language Server Protocol Implementations

* [Go: gopls](https://github.com/golang/tools/tree/master/gopls)
```
go install golang.org/x/tools/gopls@latest
```

* [Python](https://github.com/microsoft/pyright)

```
brew install pyright
```

## Formatters

* [LUA](https://github.com/Koihik/LuaFormatter): use `lua-format -i *.lua`
