vim.g.mapleader = ","

vim.opt.encoding = "utf-8"

vim.opt.compatible = false
vim.opt.hlsearch = true
vim.opt.relativenumber = true
vim.opt.laststatus = 2
vim.opt.vb = true
vim.opt.ruler = true
vim.opt.spell = true
vim.opt.spelllang = "en_us"
vim.opt.autoindent = true
vim.opt.colorcolumn = "120"
vim.opt.textwidth = 120
vim.opt.mouse = "a"
vim.opt.clipboard = "unnamed"
vim.opt.scrollbind = false
vim.opt.wildmenu = true

vim.opt.tabstop = 3
vim.opt.shiftwidth = 3

-- filetype related 
vim.api.nvim_create_autocmd("FileType", {
    pattern = {"gitcommit"},
    callback = function(ev)
        vim.api.nvim_set_option_value("textwidth", 72, {scope = "local"})
    end
})

vim.api.nvim_create_autocmd("FileType", {
    pattern = {"markdown"},
    callback = function(ev)
        vim.api.nvim_set_option_value("textwidth", 0, {scope = "local"})
        vim.api.nvim_set_option_value("wrapmargin", 0, {scope = "local"})
        vim.api.nvim_set_option_value("linebreak", true, {scope = "local"})
    end
})
