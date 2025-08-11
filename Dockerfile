FROM php:8.4-cli-alpine

COPY --from=composer/composer:2.8-bin /composer /usr/bin/composer

COPY equationlabs-php-cli /usr/bin/

ENTRYPOINT ["/usr/bin/equationlabs-php-cli"]

