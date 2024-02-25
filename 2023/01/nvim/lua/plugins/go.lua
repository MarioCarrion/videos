local setup, go = pcall(require, "go")
if not setup then return end

local lsp_setup, go_lsp = pcall(require, "go.lsp")
if not lsp_setup then return end

local lspconfig_setup, lspconfig = pcall(require, "lspconfig")
if not lspconfig_setup then return end

local lsp_signature_setup, lsp_signature = pcall(require, "lsp_signature")
if not lsp_signature_setup then return end

go.setup({
    gofmt = "gofmt"
    -- run_in_floaterm = true,
    -- floaterm = {   -- position
    --   posititon = 'auto', -- one of {`top`, `bottom`, `left`, `right`, `center`, `auto`}
    --   width = 0.45, -- width of float window if not auto
    --   height = 0.98, -- height of float window if not auto
    --   title_colors = 'nord', -- default to nord, one of {'nord', 'tokyo', 'dracula', 'rainbow', 'solarized ', 'monokai'}
    --                             -- can also set to a list of colors to define colors to choose from
    --                             -- e.g {'#D8DEE9', '#5E81AC', '#88C0D0', '#EBCB8B', '#A3BE8C', '#B48EAD'}
    -- },
})

lspconfig.gopls.setup(go_lsp.config())

vim.cmd("augroup go")
vim.cmd("autocmd!")

-- vim.cmd("autocmd FileType go nmap <leader>Gb :GoBuild %:.:h<CR>")
vim.cmd("autocmd FileType go nmap <leader>Gb :GoBuild %:p:h<CR>")
vim.cmd("autocmd FileType go nmap <leader>Gt :GoTestPkg<CR>")
vim.cmd("autocmd FileType go nmap <leader>Gc :GoCoverage -p<CR>")

-- vim.cmd("autocmd FileType go nmap <Leader>gc :lua require('go.comment').gen()")

-- vim.fn.getcwd()

-- vim.cmd("autocmd CursorHold <buffer> lua vim.lsp.buf.document_highlight()")
--

vim.cmd(
    "autocmd BufWritePre (InsertLeave?) <buffer> lua vim.lsp.buf.formatting_sync(nil,500)")

vim.cmd(
    "autocmd Filetype go command! -bang A :lua require('go.alternate').switch(true, '')")
vim.cmd(
    "autocmd Filetype go command! -bang AV :lua require('go.alternate').switch(true, 'vsplit')")
vim.cmd(
    "autocmd Filetype go command! -bang AS :lua require('go.alternate').switch(true, 'split')")

vim.cmd("augroup END")

-- {

lsp_signature.setup({
    auto_close_after = 1,
    handler_opts = {
        border = "rounded" -- double, rounded, single, shadow, none, or a table of borders
    },
    floating_window = true, -- show hint in a floating window, set to false for virtual text only mode
    floating_window_off_x = 5, -- adjust float windows x position.
    floating_window_off_y = function() -- adjust float windows y position. e.g. set to -2 can make floating window move up 2 lines
        local linenr = vim.api.nvim_win_get_cursor(0)[1] -- buf line number
        local pumheight = vim.o.pumheight
        local winline = vim.fn.winline() -- line number in the window
        local winheight = vim.fn.winheight(0)

        -- window top
        if winline - 1 < pumheight then return pumheight end

        -- window bottom
        if winheight - winline < pumheight then return -pumheight end
        return 0
    end
})
