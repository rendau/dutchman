#!/bin/sh

./refresh_config.sh

nginx -g 'daemon off;'
