const cssnano = require('cssnano');
// postcss.config.js
module.exports = {
  plugins: [
    require('tailwindcss'),
    cssnano({
      preset: 'default',
    }),
    require('autoprefixer'),
  ]
};
