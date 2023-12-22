/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.{html,js,templ,go}",
  ],
  theme: {
    extend: {},
    // fontFamily: {
    //   sans: ["Quicksand"],
    // },
  },
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};
