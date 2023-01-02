#!/bin/sh

confpath=/www/site/config.js

# Recreate config file
rm -rf ${confpath}
touch ${confpath}

# Add assignment
echo "window.API_URL = '$API_URL';" >> ${confpath}
echo "window.ACCOUNT_API_URL = '$ACCOUNT_API_URL';" >> ${confpath}
echo "window.ACCOUNT_AUTH_URL = 'ACCOUNT_AUTH_URL';" >> ${confpath}
echo "window.ACCOUNT_CABINET_URL = 'ACCOUNT_CABINET_URL';" >> ${confpath}
