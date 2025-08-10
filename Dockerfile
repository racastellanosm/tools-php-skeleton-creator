FROM scratch AS build

COPY --from=composer/composer:2.8.10-bin /composer /usr/bin/composer

COPY dist/equationlabs-php-cli /usr/local/bin/

FROM scratch AS production

ENTRYPOINT ["/usr/local/bin/equationlabs-php-cli"]

COPY --from=build . .
