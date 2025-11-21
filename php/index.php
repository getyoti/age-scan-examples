<?php

require_once './vendor/autoload.php';

use Yoti\Http\RequestBuilder;
use Yoti\Http\Payload;

// Configuration - update these values or set environment variables
$baseUrl = getenv('BASE_URL') ?: 'https://api.yoti.com/ai/v1';
$endpoint = getenv('ENDPOINT') ?: 'age-antispoofing';
$pemFilePath = getenv('PEM_FILE_PATH') ?: 'key.pem';
$sdkId = getenv('SDK_ID');
$imagePath = getenv('IMAGE_PATH') ?: './image.jpeg';

if (!$sdkId) {
    die("Error: SDK_ID environment variable is required. Please set SDK_ID or update the code directly.\n");
}

$image = file_get_contents($imagePath);

$payload = [ "data" => base64_encode($image) ];

$request = (new RequestBuilder())
    ->withBaseUrl($baseUrl)
    ->withPemFilePath($pemFilePath)
    ->withEndpoint('/' . $endpoint)
    ->withMethod('POST')
    ->withPayload(Payload::fromJsonData($payload))
    ->withHeader('X-Yoti-Auth-Id', $sdkId)
    ->build();

$response = $request->execute();

$body = $response->getBody();

echo $body;