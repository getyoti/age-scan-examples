using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using Yoti.Auth.Web;
using System;
using System.Net.Http;
using System.IO;
using Org.BouncyCastle.OpenSsl;
using Org.BouncyCastle.Crypto;
using System.Text;

namespace CoreExample
{

    class Program
    {
        static void Main(string[] args)
        {
            var serviceCollection = new ServiceCollection();
            ConfigureServices(serviceCollection);
            var serviceProvider = serviceCollection.BuildServiceProvider();
            var logger = serviceProvider.GetService<ILogger<Program>>();

            if (File.Exists(".env"))
            {
                logger.LogInformation("using environment variables from .env file");
                DotNetEnv.Env.Load(".env");
            }
       

            //load private RSA key

            string yotiKeyFilePath = Environment.GetEnvironmentVariable("PEM_FILE_PATH");
            logger.LogInformation(
                string.Format(
                    "yotiKeyFilePath='{0}'",
                    yotiKeyFilePath));

            StreamReader privateKeyStream = System.IO.File.OpenText(yotiKeyFilePath);
            var pemReader = new PemReader(privateKeyStream);
            var key = (AsymmetricCipherKeyPair)pemReader.ReadObject();


            //read img file

            byte[] imgBytes = System.IO.File.ReadAllBytes(DotNetEnv.Env.GetString("TEST_IMAGE_PATH"));
           
            string serializedRequest = Newtonsoft.Json.JsonConvert.SerializeObject(new
            {
                data =  Convert.ToBase64String(imgBytes)
            });

            byte[] byteContent = Encoding.UTF8.GetBytes(serializedRequest);

            Request request = new RequestBuilder()
                .WithBaseUri(new Uri(DotNetEnv.Env.GetString("BASE_URL") + "/api/v1/age-verification"))
                .WithEndpoint("/checks")
                .WithHttpMethod(HttpMethod.Post)
                .WithKeyPair(key)
                .WithHeader("X-Yoti-Auth-Id", DotNetEnv.Env.GetString("SDK_ID"))
                .WithContent(byteContent)
                .Build();


            HttpClient httpClient = new HttpClient();
            HttpResponseMessage response = request.Execute(httpClient).Result;
            var respStr = response.Content.ReadAsStringAsync().Result;
            Console.Out.WriteLine(respStr);

        }

        private static void ConfigureServices(IServiceCollection services)
        {
            services.AddLogging(configure =>
            {
                configure.AddConsole();
            });
     

        }


    }
}



