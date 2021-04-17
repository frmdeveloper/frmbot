#!/usr/bin/bash

apt-get update
apt-get upgrade
apt-get install nodejs -y
apt-get install libwebp -y
apt-get install ffmpeg -y
apt-get install wget -y
apt-get install webp -y
apt-get install tesseract -y
apt-get install imagemagick -y
apt-get install qrencode -y
apt-get install translate-shell -y
apt-get install gconf-service libasound2 libatk1.0-0 libatk-bridge2.0-0 libc6 libcairo2 libcups2 libdbus-1-3 libexpat1 libfontconfig1 libgcc1 libgconf-2-4 libgdk-pixbuf2.0-0 libglib2.0-0 libgtk-3-0 libnspr4 libpango-1.0-0 libpangocairo-1.0-0 libstdc++6 libx11-6 libx11-xcb1 libxcb1 libxcomposite1 libxcursor1 libxdamage1 libxext6 libxfixes3 libxi6 libxrandr2 libxrender1 libxss1 libxtst6 ca-certificates fonts-liberation libappindicator1 libnss3 lsb-release xdg-utils wget
wget -O ~/../usr/share/tessdata/ind.traineddata "https://github.com/tesseract-ocr/tessdata/blob/master/ind.traineddata?raw=true"
npm install n -g
n 14
npm install -g npm
npm install
npm audit fix

echo "[*] All dependencies have been installed, please run the command \"npm start\" to immediately start the script"
