/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors:{
        color1: "#C1DBF7",
        color2: "#88E4AD",
        color3: "#4997EB",
        writeColor: "#222222"
      },
      boxShadow:{
        "innerShadow": 'inset 0 2px 4px rgba(0, 0, 0, 0.6)',
      },
      fontSize:{
        sm: '0.6rem',
        "dash": '0.8rem'
      }
    },
  },
  plugins: [],
}