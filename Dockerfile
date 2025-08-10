FROM scratch AS build

COPY --from=composer/composer:2.8.10-bin /composer /usr/bin/composer

COPY equationlabs-php-cli /usr/local/bin/

FROM scratch AS production

COPY --from=build . .

ENTRYPOINT ["/usr/local/bin/equationlabs-php-cli"]
