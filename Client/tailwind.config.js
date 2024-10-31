/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        marigold: {
        '100': '#FF5733',
        '200': '#FF7F33',
        '300': '#FFA633',
        '400': '#FFD333',
        '500': '#FFE633'
        },
        dusk: {
          '100': '#000000',
          '200': '#241e21',
          '300': '#66535b',
          '400': '#b09f71'
        }
      }
    }
  },
  plugins: [],
}