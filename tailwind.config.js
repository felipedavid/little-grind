/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.templ", "./**/*.templ"],
  plugins: [require("daisyui")],
  daisyui: {
      themes: ["sunset"],
  }
}

