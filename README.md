<p align="center">
  <img src=".github/wb-logo.png" width="50" alt="Weights & Biases"/>
</p>

<h1 align="center">Weights & Biases Server</h1>
<h4 align="center">Server is the self hosted version of Weights &amp; Biases.</h4>

<p align="center">
  <a href="#Quickstart">Quickstart</a> •
  <a href="https://docs.wandb.ai/guides/self-hosted">Documentation</a> •
  <a href="https://github.com/wandb/terraform-aws-wandb">AWS Terraform</a> •
  <a href="https://github.com/wandb/terraform-google-wandb">Google Terraform</a> •
  <a href="https://github.com/wandb/local/tree/main/terraform/azure">Azure Terraform</a>
</p>

## Quickstart

1. On a machine with [Docker](https://docker.com) and [Python](https://www.python.org/) installed, run:
    ```
    1 pip install wandb --upgrade
    2 wandb server start
    ```
2. Generate a free license from the [Deployer](https://deploy.wandb.ai/).
3. Add it to your W&amp;B Server's localhost's settings.

  **Paste the license in the /system-admin page on your localhost**
  
  ![2022-02-24 22 13 59](https://user-images.githubusercontent.com/25806817/166265834-6a9d1be8-2af5-4c63-872e-8e5b3e4082aa.gif)


## Docker

Running `wandb server start` will start our server and forward port 8080 on the host.
To have other machines report metrics to this server run: `wandb login
--host=http://X.X.X.X:8080`.  You can also configure other machines with the
following environment variables:

```
WANDB_BASE_URL=http://X.X.X.X:8080
WANDB_API_KEY=XXXX
```

To run W&amp;B local manually, you can use the following docker command:

```
docker run --rm -d -v wandb:/vol -p 8080:8080 --name wandb-local wandb/local
```

## Production

By default this Docker container is not appropriate for production environments.
You can email `contact@wandb.com` to obtain a license that unlocks production
features such as external MySQL, cloud storage, and SSO. 

You can find [Terraform](https://www.terraform.io/) scripts for provisioning
wandb/local in a production environment:

 - [Google](https://github.com/wandb/terraform-google-wandb)
 - [AWS](https://github.com/wandb/terraform-aws-wandb)
 - [Azure](https://github.com/wandb/local/tree/main/terraform/azure)

## Documentation

More documentation about running wandb/local on your own servers can be found
[here](https://docs.wandb.com/self-hosted/local).
