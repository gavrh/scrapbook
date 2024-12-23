/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./views/**/*.{html,js}"],
    theme: {
        extend: {
            colors: {
                'link': '#00008A',
                'link-active': '#FF0000',
            },
        },
    },
    plugins: [
        function ({ addVariant }) {
            addVariant('child', '& > *');
            addVariant('child-hover', '& > *:hover');
        }
    ],
}
