local cmp_setup, cmp = pcall(require, "cmp")
if not cmp_setup then return end

local lspkind_setup, lspkind = pcall(require, "lspkind")
if not lspkind_setup then return end

local lspconfig_setup, lspconfig = pcall(require, "lspconfig")
if not lspconfig_setup then return end

local cmp_nvim_lsp_setup, cmp_nvim_lsp = pcall(require, "cmp_nvim_lsp")
if not cmp_nvim_lsp_setup then return end

local lspconfig_bundler_setup, lspconfig_bundler = pcall(require,
                                                         "lspconfig-bundler")
if not lspconfig_bundler_setup then return end

cmp.setup({
    snippet = {expand = function(args) vim.fn["UltiSnips#Anon"](args.body) end},
    window = {
        completion = cmp.config.window.bordered(),
        documentation = cmp.config.window.bordered()
    },
    mapping = cmp.mapping.preset.insert({
        ['<C-b>'] = cmp.mapping.scroll_docs(-4),
        ['<C-f>'] = cmp.mapping.scroll_docs(4),
        ['<C-Space>'] = cmp.mapping.complete(),
        ['<C-e>'] = cmp.mapping.abort(),
        ['<CR>'] = cmp.mapping.confirm({select = true}) -- Accept currently selected item. Set `select` to `false` to only confirm explicitly selected items.
    }),
    sources = cmp.config.sources({{name = 'nvim_lsp'}, {name = 'ultisnips'}},
                                 {{name = 'buffer'}})
})

local capabilities = cmp_nvim_lsp.default_capabilities()

lspconfig['pyright'].setup {}

lspconfig['solargraph'].setup {}

cmp.setup {
    formatting = {
        format = lspkind.cmp_format({
            mode = "symbol_text",
            maxwidth = 50,

            before = function(entry, vim_item) return vim_item end
        })
    }
}
