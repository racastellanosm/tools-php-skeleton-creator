FROM php:8.4-cli-alpine

COPY --from=mlocati/php-extension-installer:2.8.5 /usr/bin/install-php-extensions /usr/local/bin/
COPY --from=composer/composer:2.8.10-bin /composer /usr/bin/composer

# Install PHP extensions
RUN install-php-extensions pdo_mysql zip opcache 

# Set the working directory
WORKDIR /app

# Install Composer
ADD composer.json composer.lock /app/
RUN composer install --no-scripts

# Copy the rest of the application
COPY . /app