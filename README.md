# age-scan-examples

This repository contains examples for integrating with Yoti AI Services API in multiple programming languages.

## API Overview

The examples use the Yoti AI Services API with the following structure:

- **Base URL**: `https://api.yoti.com/ai/v1`
- **Available Endpoints**:
  - `/age` - Age estimation only
  - `/antispoofing` - Liveness/antispoofing check only
  - `/age-antispoofing` - Combined age estimation and antispoofing check (default)

## Available Examples

- [Python](./python/README.md)
- [JavaScript/Node.js](./javascript/README.md)
- [Java](./java/README.md)
- [Go](./go/README.md)
- [.NET Core](./dotnet/CoreExample/README.md)
- [PHP](./php/README.md)

## Getting Started

Each example requires:

1. A Yoti SDK ID and PEM file (obtain from [Yoti Hub](https://hub.yoti.com))
2. Configuration of environment variables or properties file
3. An image file for testing (provided as `testimage.jpg` or `image.jpeg`)

See individual example READMEs for language-specific setup instructions.

## Reference

For more information, see:
- [Yoti AI Services API Documentation](https://developers.yoti.com/ai-services-api)
- [Web FCM Demo](https://github.com/getyoti/web-fcm-demo)