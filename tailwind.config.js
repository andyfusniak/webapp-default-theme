const colors = require('tailwindcss/colors');

module.exports = {
  purge: {
    mode: 'layers',
    content: ['./templates/**/*.html'],
  },
  darkMode: 'media', // or 'media' or 'class'
  theme: {
    extend: {
      colors: {
        lightblue: colors.lightBlue,
        gray: colors.trueGray,
        coolGray: colors.coolGray,
        violet: colors.violet,
      }
    },
  },
  variants: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms')
  ],
};
