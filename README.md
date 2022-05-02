# W&amp;B Local

W&amp;B Local is the self hosted version of Weights &amp; Biases.

## Quickstart

On a machine with [Docker](https://docker.com) and Python installed run the following commands to startup our server:

1. `pip install wandb --upgrade`
2. `wandb local`

This will start our server and forward port 8080 on the host.  To have other machines report metrics to this server run: `wandb login --host=http://X.X.X.X:8080`.  You can also configure other machines with the following environment variables:

```
WANDB_BASE_URL=http://X.X.X.X:8080
WANDB_API_KEY=XXXX
```

## Docker

To run W&amp;B local manually, you can use the following docker command:

```
docker run --rm -d -v wandb:/vol -p 8080:8080 --name wandb-local wandb/local
```

## Production

By default this Docker container is not appropriate for production environments.  You can email `contact@wandb.com` to obtain a license that unlocks production features such as external MySQL, cloud storage, and SSO.  This repository provides [Terraform](https://www.terraform.io/) scripts for provisioning the wandb/local container in a production environment.

## Quickstart 

Run a W&B Server locally on any machine with Docker and Python installed with only two lines of code. For more advanced configuration, see our [Private Hosting documentation](https://docs.wandb.ai/guides/self-hosted/local) and user guide. 

1. On any machine with [Docker](https://www.docker.com/) and [Python](https://www.python.org/) installed, run:
    
    ```
    1 pip install wandb
    2 wandb local
    ```
2. Generate a free license from the [Deployer](https://deploy.wandb.ai/).
3. Add it to your local settings.

****Paste the license in the `/system-admin` page on your localhost**

![2022-02-24 22 13 59](https://user-images.githubusercontent.com/25806817/166264112-d971c240-cef1-437f-b510-4a391072147c.gif)

## Documentation

More documentation about running wandb/local on your own servers can be found [here](https://docs.wandb.com/self-hosted/local).
