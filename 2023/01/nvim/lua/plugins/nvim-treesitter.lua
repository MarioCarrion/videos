local setup, treesitter = pcall(require, "nvim-treesitter.configs")
if not setup then return end

vim.opt.foldmethod = "expr"
vim.opt.foldexpr = "nvim_treesitter#foldexpr()"
vim.opt.foldenable = false

local vim = vim
local api = vim.api
local M = {}
function M.nvim_create_augroups(definitions)
    for group_name, definition in pairs(definitions) do
        api.nvim_command('augroup ' .. group_name)
        api.nvim_command('autocmd!')
        for _, def in ipairs(definition) do
            local command = table.concat(vim.tbl_flatten {'autocmd', def}, ' ')
            api.nvim_command(command)
        end
        api.nvim_command('augroup END')
    end
end

local autoCommands = {
    open_folds = {{"BufReadPost,FileReadPost", "*", "normal zR"}}
}

M.nvim_create_augroups(autoCommands)

treesitter.setup {
    ensure_installed = {
        "dockerfile", "gitignore", "go", "gomod", "gowork", "javascript",
        "json", "lua", "markdown", "proto", "python", "rego", "ruby", "sql",
        "svelte", "yaml"
    },
    indent = {enable = true},
    auto_install = true,
    sync_install = false,
    highlight = {
        enable = true,
        disable = {"markdown"}
        --   -- Setting this to true will run `:h syntax` and tree-sitter at the same time.
        --   -- Set this to `true` if you depend on 'syntax' being enabled (like for indentation).
        --   -- Using this option may slow down your editor, and you may see some duplicate highlights.
        --   -- Instead of true it can also be a list of languages
        --   additional_vim_regex_highlighting = false,
    }
}
