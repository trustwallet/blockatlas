release: bin/setup -c $HOME/config.yml
web: bin/api -c $HOME/config.yml -p $PORT
consumer_transactions: CONSUMER_SERVICE=transactions bin/consumer -c $HOME/config.yml
consumer_subscriptions: CONSUMER_SERVICE=subscriptions bin/consumer -c $HOME/config.yml
consumer_subscriptions_tokens: CONSUMER_SERVICE=subscriptions_tokens bin/consumer -c $HOME/config.yml
consumer_tokens: CONSUMER_SERVICE=tokens bin/consumer -c $HOME/config.yml
parser: bin/parser -c $HOME/config.yml
