package com.yoti.agescan;

import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.io.*;
import java.net.URISyntaxException;
import java.nio.charset.Charset;
import java.security.*;
import java.util.Properties;

import com.yoti.api.client.spi.remote.call.ResourceException;
import com.yoti.api.client.spi.remote.call.SignedRequest;
import com.yoti.api.client.spi.remote.call.SignedRequestBuilder;
import io.vertx.core.json.JsonObject;
import org.bouncycastle.openssl.PEMKeyPair;
import org.bouncycastle.openssl.PEMParser;
import org.bouncycastle.openssl.jcajce.JcaPEMKeyConverter;

public class Application {

    static Properties prop = new Properties();
    public static void main(String[] args) {
        Security.addProvider(new org.bouncycastle.jce.provider.BouncyCastleProvider());
        try (InputStream input = Application.class.getClassLoader().getResourceAsStream("application.properties")) {

            if (input == null) {
                System.out.println("Unable to find application.properties");
                return;
            }
            prop.load(input);
        } catch (IOException ex) {
            ex.printStackTrace();
        }

        BufferedImage bufferedImage = null;
        ByteArrayOutputStream byteArrayOutputStream = null;
        try {
            bufferedImage = ImageIO.read(new File(Application.class.getClassLoader().
                    getResource(prop.getProperty("TEST_IMAGE_PATH")).
                    getFile()));
            byteArrayOutputStream = new ByteArrayOutputStream();
            ImageIO.write(bufferedImage, "jpg", byteArrayOutputStream );
        } catch (IOException e) {
            e.printStackTrace();
        }

        byte[] image = byteArrayOutputStream.toByteArray();
        JsonObject body = new JsonObject();
        body.put("data", image);


        byte[] payload = body.encode().getBytes();

        try {
            SignedRequest signedRequest = SignedRequestBuilder.newInstance()
                    .withKeyPair(findKeyPair())
                    .withBaseUrl(prop.getProperty("HOST") + "/api/v1/age-verification")
                    .withEndpoint("/checks")
                    .withPayload(payload)
                    .withHttpMethod("POST")
                    .withHeader("X-Yoti-Auth-Id", prop.getProperty("SDK_ID"))
                    .build();

            Result result = signedRequest.execute(Result.class);
            System.out.println(result);
        } catch (GeneralSecurityException | URISyntaxException | IOException | ResourceException ex) {
            ex.printStackTrace();
        }
    }

    /**
     *
     * @return KeyPair
     * @throws IOException
     */
    private static KeyPair findKeyPair() throws IOException {
        InputStream keyStream = new FileInputStream(Application.class.getClassLoader().getResource(prop.getProperty("PEM_FILE_PATH")).getFile());
        PEMParser reader = new PEMParser(new BufferedReader(new InputStreamReader(keyStream, Charset.defaultCharset())));
        KeyPair keyPair = null;
        for (Object o = null; (o = reader.readObject()) != null;) {
            if (o instanceof PEMKeyPair) {
                keyPair = new JcaPEMKeyConverter().getKeyPair((PEMKeyPair) o);
                break;
            }
        }
        return keyPair;
    }
}
