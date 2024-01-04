#!/bin/bash
source ./env
for ticker in ${tickers[@]}
  do
    ./stocks-helper ${ticker} ${TELEGRAM_CHANNEL} ${NOTIFIER_ADDR}
done

