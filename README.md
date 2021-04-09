# Scrapi

## TL;DR
Scrapi (scr-API) is a microservice that exposes results from the Scrapy library (https://github.com/Vorstenbosch/scrapy) via a REST interface.

To deploy a Scrapi instance to your kubernetes cluster run:
```bash
helm upgrade --install scrapi ./helm/scrapi --values  ./helm/scrapi/values.yaml
```