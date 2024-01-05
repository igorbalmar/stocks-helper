#!/bin/bash
source /usr/local/etc/stocks-helper/tickers.env
sleep_time=1h
if [[ -z ${tickers} || -z ${BRAPI_TOKEN} || -z ${TELEGRAM_CHANNEL} || -z ${NOTIFIER_ADDR} ]]
 then
 echo "Check if the environment variables are set accordingly!
 Mandatory: tickers(array), BRAPI_TOKEN, TELEGRAM_CHANNEL, NOTIFIER_ADDR"
 exit 1
fi
while true
  do
    for ticker in ${tickers[@]}
      do
        stocks-helper ${ticker} ${TELEGRAM_CHANNEL} ${NOTIFIER_ADDR}
    done
    echo "Sleeping for ${sleep_time}"
    sleep ${sleep_time}
done

