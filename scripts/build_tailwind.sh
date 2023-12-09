#!/bin/bash
cwd=$(pwd)

cd $cwd/web
echo "Build tailwind"
npm run build-tailwind
echo "Done build tailwind"
