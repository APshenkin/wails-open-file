# README

## About

This is the official Wails React-TS template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## Ubuntu debian package with file association and app icon
To build linux debian package with file association and app icon I used [nfpm](https://nfpm.goreleaser.com/)
To reproduce locally first install nfpm using one of described methods [here](https://nfpm.goreleaser.com/install/)
Note that svg icons are used for Ubuntu as this is preferred format to have good quality on all screen resolutions.
Then run following command from project root directory:
```sh
$ wails build
$ cd ./build/linux
$ nfpm pkg --packager deb --target .
```
