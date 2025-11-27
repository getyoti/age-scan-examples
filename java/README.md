# Java Example

- Save your PEM file into the `resources/keys` directory and name it key.pem
- Edit the `application.properties` file, adding client SDK ID
- Optional - Update the ENDPOINT in `application.properties` (default is `age-antispoofing`, can be `age`, `antispoofing`, or `age-antispoofing`)
- Optional - Replace the testimage.jpg with your own.
- Build the project `mvn clean package`
- Run the example `java -jar target/age-scan-1.0.jar`
