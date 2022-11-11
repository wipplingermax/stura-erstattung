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
        'stura-red-hover': '#993333',
        'body-background': '#111928',
      },
    }
  },
  plugins: [],
}
