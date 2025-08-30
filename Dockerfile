FROM php:8.4-cli-alpine AS production

WORKDIR /app

# Runtime dependencies
COPY --from=composer/composer:2.8-bin /composer /usr/bin/composer

# Application Binary
COPY php-skeleton-creator-cli /app/php-skeleton-creator-cli

ENTRYPOINT ["/app/php-skeleton-creator-cli"]
