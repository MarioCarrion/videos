return {
    "numToStr/Comment.nvim",
    dependencies = {"nvim-treesitter/nvim-treesitter"},
    config = function() require("Comment").setup() end
}
