
require("dotenv").config();
const { RequestBuilder, Payload } = require('yoti');
const fs = require('fs')

var image = fs.readFileSync(process.env.TEST_IMAGE_PATH);
var imageRequest = {"data": image.toString('base64')}

const request = new RequestBuilder()
    .withBaseUrl(process.env.BASE_URL + '/api/v1/age-verification')
    .withPemFilePath(process.env.PEM_FILE_PATH)
    .withEndpoint('/checks')
    .withPayload(new Payload(imageRequest))
    .withMethod('POST')
    .withHeader('X-Yoti-Auth-Id', process.env.SDK_ID)
    .build();

try {
   const response = request.execute();
   response.then(function(value) {
      console.log(value);
    });
 } catch (error) {
   console.log(error);
 }
