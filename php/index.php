<?php

require_once './vendor/autoload.php';

use Yoti\Http\RequestBuilder;
use Yoti\Http\Payload;

$image = file_get_contents('./image.jpeg');

$payload = [ "data" => base64_encode($image) ];

$request = (new RequestBuilder())
    ->withBaseUrl('<YOTI_BASE_URL>/api/v1/age-verification')
    ->withPemFilePath('key.pem')
    ->withEndpoint('/checks')
    ->withMethod('POST')
    ->withPayload(Payload::fromJsonData($payload))
    ->withHeader('X-Yoti-Auth-Id', '<SDK_ID>')
    ->build();

$response = $request->execute();

$body = $response->getBody();
