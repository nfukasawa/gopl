#!/bin/bash
set -eux

go run main.go \
    http://google.com \
    http://youtube.com \
    http://facebook.com \
    http://baidu.com \
    http://wikipedia.org \
    http://yahoo.com \
    http://reddit.com \
    http://google.co.in \
    http://qq.com \
    http://taobao.com \
    http://twitter.com \
    http://amazon.com \
    http://sohu.com \
    http://google.co.jp \
    http://live.com \
    http://tmall.com \
    http://vk.com \
    http://instagram.com \
    http://sina.com.cn \
    http://360.cn

rm results_*