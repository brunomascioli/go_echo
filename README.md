go_echo — Instruções de execução
================================

Este repositório contém uma pequena aplicação de exemplo em Go (servidores echo), um cliente de teste em Python e uma configuração de HAProxy para balanceamento. Abaixo estão instruções simples para:

- rodar os servidores e o HAProxy com Docker Compose;
- criar e ativar uma venv Python;
- instalar dependências do `requirements.txt`.

Pré-requisitos
--------------

- Docker e Docker Compose instalados no sistema.
- Python 3.8+ (ou compatível) instalado para executar o cliente de teste e criar a venv.

Verifique as versões:

```bash
docker --version
docker compose version
python3 --version
```

Rodando os servidores e o HAProxy com Docker Compose
---------------------------------------------------

O compose já está configurado para construir as imagens e subir os containers necessários (servidores Go e HAProxy). Para construir as imagens e iniciar os serviços em foreground, execute:

```bash
docker compose up --build
```

Isso irá:

- construir a imagem do servidor Go (a partir do `Dockerfile`);
- criar e iniciar os containers definidos (servidores e HAProxy);
- mostrar os logs no terminal.

Para rodar em background (detached), adicione `-d`:

```bash
docker compose up --build -d
```

Para parar e remover os containers criados pelo compose:

```bash
docker compose down
```

Criar e ativar uma venv Python
-----------------------------

No diretório do projeto, crie uma venv chamada `venv` com:

```bash
python3 -m venv venv
```

Ative a venv:

Linux/macOS (zsh/bash):

```bash
source venv/bin/activate
```

Windows (PowerShell):

```powershell
venv\\Scripts\\Activate.ps1
```

Ao ativar a venv, o prompt do shell deve mostrar o nome `(venv)`.

Instalar dependências do `requirements.txt`
-----------------------------------------

Com a venv ativada (ou se preferir instalar globalmente), instale as dependências:

```bash
pip install -r requirements.txt
```

Observação: no enunciado foi usado `pip i -r requirements.txt`; o comando correto é `pip install -r requirements.txt`.

Testando o cliente
------------------

Há um cliente Python de exemplo em `client/` (por exemplo `locustfile.py`). Dependendo do que você quer testar:

- Para executar scripts Python de teste (por exemplo, `locustfile.py`), ative a venv e rode os comandos apropriados (por ex. `locust` se estiver usando Locust).

Exemplo rápido (após ativar venv e instalar requirements):

```bash
# rodar Locust para testes de carga (execute a partir da raiz do projeto)
locust -f client/locustfile.py
```

Notas finais
-----------

- Se você usar `docker compose up --build`, o Docker fará a construção usando o `Dockerfile` presente na raiz.
- Se algum comando falhar por falta de permissões, tente prefixar com `sudo` ou ajuste o seu ambiente (recomenda-se adicionar o usuário ao grupo `docker` no Linux para evitar `sudo`).
- Em caso de dúvidas sobre os serviços que foram levantados, confira os arquivos `docker-compose.yml`, `Dockerfile` e `haproxy.cfg` para detalhes de portas e configuração.

---

README atualizado com instruções básicas para executar o ambiente de desenvolvimento e testes.

