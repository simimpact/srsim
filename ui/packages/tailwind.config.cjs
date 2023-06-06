/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{html,js,ts,jsx,tsx}"],
  darkMode: "class",
  theme: {
    container: {
      center: true,
      padding: "2rem",
      screens: {
        "2xl": "1400px",
      },
    },
    extend: {
      colors: {
        border: "hsl(var(--border))",
        input: "hsl(var(--input))",
        ring: "hsl(var(--ring))",
        background: "hsl(var(--background))",
        foreground: "hsl(var(--foreground))",
        primary: {
          DEFAULT: "hsl(var(--primary))",
          foreground: "hsl(var(--primary-foreground))",
        },
        secondary: {
          DEFAULT: "hsl(var(--secondary))",
          foreground: "hsl(var(--secondary-foreground))",
        },
        destructive: {
          DEFAULT: "hsl(var(--destructive))",
          foreground: "hsl(var(--destructive-foreground))",
        },
        muted: {
          DEFAULT: "hsl(var(--muted))",
          foreground: "hsl(var(--muted-foreground))",
        },
        accent: {
          DEFAULT: "hsl(var(--accent))",
          foreground: "hsl(var(--accent-foreground))",
        },
        popover: {
          DEFAULT: "hsl(var(--popover))",
          foreground: "hsl(var(--popover-foreground))",
        },
        card: {
          DEFAULT: "hsl(var(--card))",
          foreground: "hsl(var(--card-foreground))",
        },
        fire: {
          DEFAULT: "#CF3C17",
          100: "#D6411A",
          200: "#DE4D26",
          300: "#E56542",
          400: "#ED8B70",
          500: "#F5BAA9",
          foreground: "#FCF1EE",
        },
        ice: {
          DEFAULT: "#0284C7",
          100: "#0991D0",
          200: "#18A1DA",
          300: "#33B3E3",
          400: "#62CAEC",
          500: "#9EE2F6",
          foreground: "#EDFBFF",
        },
        wind: {
          DEFAULT: "#22C55E",
          100: "#26CE60",
          200: "#31D766",
          300: "#47DF75",
          400: "#65E889",
          500: "#91F1AC",
          foreground: "#DBFAE3",
        },
        lightning: {
          DEFAULT: "#C75DE2",
          100: "#CB61E6",
          200: "#D069EB",
          300: "#D674EF",
          400: "#DE89F4",
          500: "#E8ACF8",
          foreground: "#F7E2FD",
        },
        physical: {
          DEFAULT: "#A8A29E",
          100: "#B6B0AC",
          200: "#C4BDB9",
          300: "#D2CBC6",
          400: "#E1D9D3",
          500: "#EFE6E1",
          foreground: "#FDF4EE",
        },
        quantum: {
          DEFAULT: "#6478E1",
          100: "#6B7FE5",
          200: "#7789E9",
          300: "#8A9AEE",
          400: "#A3B0F2",
          500: "#C5CDF6",
          foreground: "#ECEEFA",
        },
        imaginary: {
          DEFAULT: "#EAB308",
          100: "#EEB70D",
          200: "#F1BD1F",
          300: "#F5C63A",
          400: "#F8D25D",
          500: "#FBE08E",
          foreground: "#FFF2CB",
        },
      },
      borderRadius: {
        lg: `var(--radius)`,
        md: `calc(var(--radius) - 2px)`,
        sm: "calc(var(--radius) - 4px)",
      },
      keyframes: {
        "accordion-down": {
          from: { height: 0 },
          to: { height: "var(--radix-accordion-content-height)" },
        },
        "accordion-up": {
          from: { height: "var(--radix-accordion-content-height)" },
          to: { height: 0 },
        },
      },
      animation: {
        "accordion-down": "accordion-down 0.2s ease-out",
        "accordion-up": "accordion-up 0.2s ease-out",
      },
    },
  },
  plugins: [require("tailwindcss-animate")],
};
