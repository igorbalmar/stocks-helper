# stocks-helper
Use por sua própria conta e risco! Não é recomendação de compra nem venda de ações!

Stocks helper, baseado na api https://brapi.dev


The main purpose of this software is to notify (right now, only via telegram) when the specified stock is under specific conditions/price.
---
O propósito deste app é executar um serviço Docker que verificar de hora em hora e notifica (via telegram) quando uma ação entra em algum estado/condição que pode ser interessante avaliar.

- Lista das ações a serem validadas. Exemplo:
Arquivo tickers.env
- Token BRAPI
BRAPI_TOKEN='brapi.dev-token'
- ID do grupo ou canal no Telegram
TELEGRAM_CHANNEL='-channel_number'
- Endereço da API do Telegram (notifier) - por padrão escuta na porta 8080
NOTIFIER_ADDR='ip.addr:8080'


# Atualizando tickers

Caso queira atualizar somente os tickersações verificadas, substitua o arquivo tickers.env da imagem

docker cp tickers.env container:/usr/local/etc/stocks-helper/tickers.env

