## Api Gateway Kong

Neste projeto, apresento um exemplo prático que utiliza o Docker para ilustrar o funcionamento dos upstreams no Kong, um API Gateway poderoso. A utilização de upstreams permite criar balanceadores de carga entre três serviços distintos e configurar o balanceamento de carga de forma eficaz. Essas funcionalidades são cruciais para garantir que seus serviços estejam sempre disponíveis e funcionando de maneira otimizada, garantindo assim uma experiência confiável para os usuários finais.

## Tecnologias Utilizadas

- Kong API Gateway
- Docker

## Pré-requisitos

Antes de começar, certifique-se de ter o Docker instalado na sua máquina.

## 🛠 Configuração

Siga os passos abaixo para configurar e executar o projeto:

1. Clone este repositório em sua máquina local:

```bash
git clone https://github.com/seu-usuario/api-gateway-kong.git
```
Navegue até o diretório do projeto:
```bash
cd api-gateway-kong
```
Inicie o ambiente usando Docker Compose:
```bash
docker-compose up -d
```
Você pode acessar a interface de administração do Kong em:
```bash
http://localhost:8001
```
A partir daqui, você pode configurar os upstreams, balanceadores de carga e verificar o status da saúde dos serviços.
