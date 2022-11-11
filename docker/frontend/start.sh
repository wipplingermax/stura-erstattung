#!/bin/sh
echo 'app started'

cd app

## install dependencies (nextJS, typescript, tailwind...)
yarn install

## insert additional app start-up routine before yarn start
yarn run dev