/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{html,js,svelte,ts}"
  ],
  theme: {
    extend: {
      colors: {
        // 'red-brown': '#dd6d5e',
        "red-brown": '#9e1e30',
        // 'background': '#ebd8b8',
        'background': '#f5f5f5',        
        'blue-imperial': '#3a4578'
      },
      fontFamily: {
        'system-ui': ['system-ui', '-apple-system', 'BlinkMacSystemFont', '"Segoe UI"', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', '"Open Sans"', '"Helvetica Neue"', 'sans-serif']
      },
    },
  },
  plugins: [],
}

