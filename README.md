# stocks-helper

The main purpose of this software is to notify (right now, only via telegram) when the specified stock is under specific conditions/price.

---

Use por sua própria conta e risco! Não é recomendação de compra nem venda de ações!

Stocks helper, baseado na api https://brapi.dev


O propósito deste app é executar um serviço Docker que verificar de hora em hora e notifica (via telegram) quando uma ação entra em algum estado/condição que pode ser interessante avaliar.

# Variáveis de ambiente a serem configuradas:

- STOCKS_DB_STRING
url exemplo = "postgres://username:password@localhost:5432/database_name"

- BRAPI_TOKEN='brapi.dev-token'
Token BRAPI

- TELEGRAM_GROUP_ID
ID do grupo ou canal no Telegram

- NOTIFIER_ADDR='ip.addr:8080'
Endereço da API do Telegram (notifier) - por padrão escuta na porta 8080

# Atualizando/Cadastrando tickers

Ver db.sql em tickers/

