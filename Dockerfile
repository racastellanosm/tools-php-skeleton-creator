FROM php:8.4-cli-alpine AS production

WORKDIR /app

# Runtime dependencies
COPY --from=composer/composer:2.8-bin /composer /usr/bin/composer

# Application Binary
COPY equationlabs-php-cli /app/equationlabs-php-cli

ENTRYPOINT ["/app/equationlabs-php-cli"]
