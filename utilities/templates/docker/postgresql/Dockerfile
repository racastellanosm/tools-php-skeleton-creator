FROM ghcr.io/racastellanosm/php.roadrunner.pgsql:8.4

# Set the working directory
WORKDIR /app

# Install Composer
ADD composer.json composer.lock /app/
RUN composer install --no-cache --no-scripts -o

# Copy the rest of the application
COPY . .

# Post install scripts
RUN composer run-script post-install-cmd