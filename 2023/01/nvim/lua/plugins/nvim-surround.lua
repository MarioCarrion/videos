local setup, surround = pcall(require, "nvim-surround")
if not setup then return end

surround.setup()
