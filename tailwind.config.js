const defaultTheme = require("tailwindcss/defaultTheme");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.tmpl"],
  //works with tailwind cli
  plugins: [require("@tailwindcss/typography")],
  theme: {
    extend: {
      colors: {
        lightBeige: "#FFFAF2",
        darkBeige: "#FFEDCA",
        mediumGreen: "#264B3C",
        darkGreen: "#123527",
        lightGreen: "#9FD4B6",
        lightBlue: "#B0C2FF",
      },
      fontFamily: {
        sans: ['"Urbanist"', ...defaultTheme.fontFamily.sans],
      },
    },
  },
};
