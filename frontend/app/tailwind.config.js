const plugin = require('tailwindcss/plugin')

module.exports = {
  purge: ['./src/**/*.svelte', './public/*.html'],
  variants: {},
  plugins: [
    require('@tailwindcss/typography'),
    require('@tailwindcss/forms'),
    require('@tailwindcss/aspect-ratio'),
    plugin(function({ addComponents }) {
      const widths = {
        '.min-w-half': {
          minWidth: '50%',
        },
        '.w-vw-90': {
          width: '90vw'
        },
        '.w-vw-80': {
          width: '80vw'
        },
        '.w-vw-85': {
          width: '85vw'
        },
        '.min-w-96': {
          minWidth: '24rem'
        },
      }
      addComponents(widths)
    })
  ]
}
