local lazypath = vim.fn.stdpath("data") .. "/lazy/lazy.nvim"
if not vim.loop.fs_stat(lazypath) then
    vim.fn.system({
        "git", "clone", "--filter=blob:none",
        "https://github.com/folke/lazy.nvim.git", "--branch=stable", -- latest stable release
        lazypath
    })
end
vim.opt.rtp:prepend(lazypath)

-- LuaFormatter off
require("lazy").setup({
    -- Look and Feel
    {"catppuccin/nvim", name = "catppuccin", priority = 1000},
    -- File Types
    {"chrisbra/csv.vim"},
    -- Productivity
    {
        "romgrk/barbar.nvim",
        dependencies = {
            "nvim-tree/nvim-web-devicons", "lewis6991/gitsigns.nvim"
        },
        init = function() vim.g.barbar_auto_setup = false end,
    },
    {"nvim-lualine/lualine.nvim"},
    {
        "nvim-tree/nvim-tree.lua",
        dependencies = {"nvim-tree/nvim-web-devicons"},
        tag = "v0.99",
    },
    {"nvim-telescope/telescope-fzf-native.nvim", build = "make"},
    {
        "nvim-telescope/telescope.nvim",
        dependencies = {"nvim-lua/plenary.nvim"},
        branch = "0.1.x"
    },
    {"lewis6991/gitsigns.nvim"},
    -- Development
    {
        "nvim-treesitter/nvim-treesitter",
        tag = "v0.9.2",
        build = ":TSUpdate",
    },
    {"nvim-treesitter/nvim-treesitter-textobjects"},
    {"rhysd/vim-clang-format"},
    {
        "ray-x/go.nvim",
        dependencies = { -- optional packages
            "ray-x/guihua.lua",
            "neovim/nvim-lspconfig",
            "nvim-treesitter/nvim-treesitter",
        },
        event = {"CmdlineEnter"},
        ft = {"go", 'gomod'},
    },
    {
        "ray-x/navigator.lua",
        requires = {
            { "ray-x/guihua.lua", run = "cd lua/fzy && make" },
            { "neovim/nvim-lspconfig" },
        },
    },
    {
        "ray-x/lsp_signature.nvim",
        event = "VeryLazy",
    },
    {"SirVer/ultisnips"},
    {"hrsh7th/cmp-nvim-lsp"},
    {"hrsh7th/nvim-cmp"},
    {"neovim/nvim-lspconfig"},
    {"onsails/lspkind-nvim"},
    {"quangnguyen30192/cmp-nvim-ultisnips"},
    {"williamboman/nvim-lsp-installer"},
    {"numToStr/Comment.nvim"},
    {"kylechui/nvim-surround", version = "*"},
    {
        "mihyaeru21/nvim-lspconfig-bundler",
        dependencies = "neovim/nvim-lspconfig"
    }
})
-- LuaFormatter on
