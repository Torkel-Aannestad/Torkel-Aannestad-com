const typographyStyles = require('./typography')

/** @type {import('tailwindcss').Config} */
const defaultTheme = require('tailwindcss/defaultTheme')
module.exports = {
    content: ["./**/*.html", "./**/*.tmpl"],
    //works with tailwind cli
    plugins: [
      require('@tailwindcss/typography'),
    ],
    theme: {
      extend: {
          colors: {
            lightBeige: '#FFF8E9',
            darkBeige: '#FFEDCA',
            mediumGreen: '#264B3C',
            darkGreen: '#123527',
            lightGreen: '#9FD4B6',
            lightBlue: '#B0C2FF',
          },
          fontFamily: {
            'sans': ['"Urbanist"', ...defaultTheme.fontFamily.sans],
          }
      },
      // typography: typographyStyles,
    },
    
}
