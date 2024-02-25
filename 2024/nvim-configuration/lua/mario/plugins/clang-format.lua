return {
    "rhysd/vim-clang-format",
    init = function()
        vim.cmd([[
autocmd FileType proto ClangFormatAutoEnable
]])
    end
}
