/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './app/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    extend:{
      colors: {
        'stura-yellow': '#FF7F00',
        'stura-red': '#990000',
        'body-background': '#111928',
      },
    }
  },
  plugins: [],
}
