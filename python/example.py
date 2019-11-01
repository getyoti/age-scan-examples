#!/usr/bin/env python

from yoti_python_sdk.http import SignedRequest, RequestHandler
import json
import requests
import base64
from dotenv import load_dotenv
import os

def execute(request):
    response = requests.request(
        url=request.url, data=request.data, headers=request.headers, method=request.method)
    return response.content

def generate_session():
    with open(os.getenv('TEST_IMAGE_PATH'), "rb") as image_file:
        encoded_string = base64.b64encode(image_file.read())
    data = {"data" : encoded_string.decode("utf-8")}

    payload_string = json.dumps(data).encode()

    signed_request = (
        SignedRequest
        .builder()
        .with_pem_file(os.getenv('PEM_FILE_PATH'))
        .with_base_url(os.getenv('HOST') + "/api/v1/age-verification")
        .with_endpoint("/checks")
        .with_http_method("POST")
        .with_header("X-Yoti-Auth-Id", os.getenv('SDK_ID'))
        .with_payload(payload_string)
        .build()

    )

	# get Yoti response
    response = signed_request.execute()
    response_payload = json.loads(response.text)
    print(response_payload)

def main():
    load_dotenv()
    generate_session()

if __name__ == '__main__':
    main()
