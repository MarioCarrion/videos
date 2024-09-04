return {
    "ray-x/navigator.lua",
    dependencies = {
        {"hrsh7th/nvim-cmp"}, {"nvim-treesitter/nvim-treesitter"},
        {"ray-x/guihua.lua", run = "cd lua/fzy && make"}, {
            "ray-x/go.nvim",
            event = {"CmdlineEnter"},
            ft = {"go", "gomod"},
            build = ':lua require("go.install").update_all_sync()'
        }, {
            "ray-x/lsp_signature.nvim", -- Show function signature when you type
            event = "VeryLazy",
            config = function() require("lsp_signature").setup() end
        }
    },
    config = function()
        require("go").setup()
        require("navigator").setup({
            lsp_signature_help = true, -- enable ray-x/lsp_signature
            lsp = {
                format_on_save = true,
                gopls = {
                    settings = {
                        gopls = {
                            hints = {
                                assignVariableTypes = false,
                                compositeLiteralFields = true,
                                compositeLiteralTypes = true,
                                constantValues = true,
                                functionTypeParameters = true,
                                parameterNames = true,
                                rangeVariableTypes = false
                            }
                        }
                    }
                }
            }
        })

        vim.api.nvim_create_autocmd("FileType", {
            pattern = {"go"},
            callback = function(ev)
                -- CTRL/control keymaps
                vim.api
                    .nvim_buf_set_keymap(0, "n", "<C-i>", ":GoImports<CR>", {})
                vim.api.nvim_buf_set_keymap(0, "n", "<C-b>", ":GoBuild %:h<CR>",
                                            {})
                vim.api.nvim_buf_set_keymap(0, "n", "<C-t>", ":GoTestPkg<CR>",
                                            {})
                vim.api.nvim_buf_set_keymap(0, "n", "<C-c>",
                                            ":GoCoverage -p<CR>", {})
                vim.api.nvim_buf_set_keymap(0, "n", "<C-r>",
                                            ":GoRun<CR>", {})

                -- Opens test files
                vim.api.nvim_buf_set_keymap(0, "n", "A",
                                            ":lua require('go.alternate').switch(true, '')<CR>",
                                            {}) -- Test
                vim.api.nvim_buf_set_keymap(0, "n", "V",
                                            ":lua require('go.alternate').switch(true, 'vsplit')<CR>",
                                            {}) -- Test Vertical
                vim.api.nvim_buf_set_keymap(0, "n", "S",
                                            ":lua require('go.alternate').switch(true, 'split')<CR>",
                                            {}) -- Test Split
            end,
            group = vim.api.nvim_create_augroup("go_autocommands",
                                                {clear = true})
        })
    end
}
