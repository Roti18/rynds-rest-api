#!/bin/sh
set -e # Exit on error

mkdir -p docs/assets/js
mkdir -p docs/assets/css
mkdir -p docs/assets/fonts/inter
mkdir -p docs/assets/fonts/geist-mono

echo "Building Tailwind..."
# Try explicit path first, usually reliable in CI/Docker
if [ -f "./node_modules/.bin/tailwindcss" ]; then
    ./node_modules/.bin/tailwindcss -i ./src/input.css -o ./docs/assets/css/style-v2.css --minify
else
    echo "Tailwind binary not found in .bin, trying npx..."
    npx tailwindcss -i ./src/input.css -o ./docs/assets/css/style-v2.css --minify
fi

echo "Copying JS libs..."
cp node_modules/axios/dist/axios.min.js docs/assets/js/axios.min.js
cp node_modules/lucide/dist/umd/lucide.min.js docs/assets/js/lucide.min.js

echo "Copying Fonts..."
cp node_modules/@fontsource/inter/files/inter-latin-400-normal.woff2 docs/assets/fonts/inter/inter-regular.woff2
cp node_modules/@fontsource/inter/files/inter-latin-700-normal.woff2 docs/assets/fonts/inter/inter-bold.woff2
cp node_modules/@fontsource/geist-mono/files/geist-mono-latin-400-normal.woff2 docs/assets/fonts/geist-mono/geist-mono-regular.woff2

echo "Assets built successfully!"
