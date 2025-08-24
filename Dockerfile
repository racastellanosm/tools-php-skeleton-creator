FROM composer/composer:2.8 AS production

WORKDIR /app

COPY equationlabs-php-cli /app/equationlabs-php-cli

ENTRYPOINT ["/app/equationlabs-php-cli"]
