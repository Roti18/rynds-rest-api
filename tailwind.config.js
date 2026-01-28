/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./docs/**/*.{html,js}"],
  // darkMode: 'class', // Removed to just use direct utility classes
  theme: {
    extend: {
      fontFamily: {
        sans: ["Inter", "sans-serif"],
        mono: ["Geist Mono", "monospace"],
      },
      colors: {
        background: "#0a0a0a",
        surface: "#121212",
        border: "#262626",
        primary: "#ffffff",
        secondary: "#a1a1aa",
        accent: "#3b82f6",
        black: "#000000",
        dark: "#050505",
      },
      animation: {
        "pulse-slow": "pulse 4s cubic-bezier(0.4, 0, 0.6, 1) infinite",
      },
    },
  },
  plugins: [],
};
