local telescope_setup, telescope = pcall(require, "telescope")
if not telescope_setup then return end

local actions_setup, actions = pcall(require, "telescope.actions")
if not actions_setup then return end

telescope.setup({
    defaults = {
        mappings = {
            i = {
                ["<C-k>"] = actions.move_selection_previous, -- move to prev result
                ["<C-j>"] = actions.move_selection_next, -- move to next result
                ["<C-q>"] = actions.send_selected_to_qflist +
                    actions.open_qflist -- send selected to quickfixlist
            }
        }
    }
})

telescope.load_extension("fzf")

local builtin = require('telescope.builtin')

vim.keymap.set('n', '<leader>ff', builtin.find_files, {}) -- Lists files in your current working directory, respects .gitignore
-- vim.keymap.set('n', '<leader>fg', builtin.live_grep, {}) -- Searches for the string under your cursor or selection in your current working directory
-- vim.keymap.set('n', '<leader>fb', builtin.buffers, {}) -- Lists open buffers in current neovim instance
-- vim.keymap.set('n', '<leader>fh', builtin.help_tags, {}) -- Lists available help tags and opens a new window with the relevant help info on <cr>
vim.keymap.set('n', '<leader>fx', builtin.treesitter, {})
vim.keymap.set('n', '<leader>fr', builtin.lsp_references, {}) -- Lists LSP references for word under the cursor
-- vim.keymap.set('n', '<leader>fi', builtin.lsp_implementations, {}) -- Goto the implementation of the word under the cursor if there's only one, otherwise show all options in Telescope
-- vim.keymap.set('n', '<leader>fd', builtin.lsp_definitions, {}) -- Goto the definition of the word under the cursor, if there's only one, otherwise show all options in Telescope

