# PHP Example

- Run `composer install`
- Save your PEM file into the same directory and name it key.pem
- Set environment variables or edit the `index.php` file:
  - `SDK_ID`: Your client SDK ID (required)
  - `BASE_URL`: API base URL (default: `https://api.yoti.com/ai/v1`)
  - `ENDPOINT`: API endpoint (default: `age-antispoofing`, can be `age`, `antispoofing`, or `age-antispoofing`)
  - `PEM_FILE_PATH`: Path to PEM file (default: `key.pem`)
  - `IMAGE_PATH`: Path to test image (default: `./image.jpeg`)
- Optional - Replace the `image.jpeg` with your own.
- Run the project with `php index.php`
