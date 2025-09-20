FROM ghcr.io/racastellanosm/php.roadrunner.mysql:8.4

# Set the working directory
WORKDIR /app

# Install Composer
ADD composer.json composer.lock /app/
RUN composer install --no-scripts

# Copy the rest of the application
COPY .. /app