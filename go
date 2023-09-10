rm -r dist
npm run build
mv build dist
npx cap sync
adb connect 192.168.1.101
#npx cap open android
npx cap run android
