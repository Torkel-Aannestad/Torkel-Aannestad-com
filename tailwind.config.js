/** @type {import('tailwindcss').Config} */
const defaultTheme = require('tailwindcss/defaultTheme')
module.exports = {
    content: ["./**/*.html", "./**/*.tmpl"],
    theme: {
        extend: {
          colors: {
            lightBeige: '#FFF8E9'
          },
          fontFamily: {
            'sans': ['"Urbanist"', ...defaultTheme.fontFamily.sans],
        }
        }
    }
}
