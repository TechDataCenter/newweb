#! /bin/sh
kill -9 $(pgrep webserver)
cd ~/go_duke_linux/newweb
git pull https://github.com/TechDataCenter/newweb.git
cd webserver/
./webserver &