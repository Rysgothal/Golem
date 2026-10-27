# Golem

Bot do Discord, escrito em Go, para gerenciar um servidor de Minecraft que roda em Docker, direto de um canal do Discord.

> **Status: em desenvolvimento.** O bot conecta, registra slash commands e responde ao `/ping`. Os comandos de controle do servidor ainda estão no roadmap.

## Sobre

O Golem nasceu de uma necessidade simples: iniciar, parar e consultar o servidor de Minecraft caseiro sem precisar acessar a máquina por SSH. O servidor já roda em um container Docker, e o bot atua como um painel de controle dentro do Discord.

Princípios do projeto:

- **Simples primeiro.** Um servidor, poucos comandos, estrutura enxuta. Complexidade só entra quando houver necessidade real.
- **O bot não é dono do servidor.** O Minecraft vive no Docker, então reiniciar ou atualizar o bot não derruba o servidor nem expulsa os jogadores.
- **Segurança desde o início.** Todo comando de controle exige cargo de administrador, e entradas de usuário são validadas antes de virar comando.

## Funcionalidades

**Implementado**

- [x] Conexão com o Discord e carregamento de configuração via `.env`
- [x] Registro de slash commands por servidor (guild)
- [x] Roteador central de comandos
- [x] `/ping`

**Planejado**

- [ ] `/running` (em andamento...)
- [ ] Checagem de administrador por ID de cargo
- [ ] `/start`: iniciar o container do servidor
- [ ] `/stop`: parar o container do servidor
- [ ] `/logs`: exibir as últimas linhas do log
- [ ] `/op`: dar permissão de operador a um jogador
- [ ] Ainda em desenvolvimento

## Tecnologias

| Item | Uso |
|---|---|
| [Go](https://go.dev/) | Linguagem do projeto |
| [discordgo](https://github.com/bwmarrin/discordgo) | Biblioteca de integração com a API do Discord |
| [godotenv](https://github.com/joho/godotenv) | Carregamento do arquivo `.env` em desenvolvimento |
| [Docker](https://www.docker.com/) | Execução do servidor de Minecraft |

## Estrutura do projeto

```
golem/
├── cmd/
│   └── golem/
│       └── main.go          # ponto de entrada: liga as peças
├── internal/
│   ├── config/              # leitura e validação da configuração
│   └── discord/             # conexão, roteador e comandos
├── .env.example             # modelo de configuração, sem segredos
├── .gitignore
├── go.mod
└── README.md
```

## Pré-requisitos

- Go (versão compatível com o `go.mod`)
- Uma aplicação criada no [Discord Developer Portal](https://discord.com/developers/applications)
- Docker e um servidor de Minecraft em container (para os comandos de controle, quando implementados)

## Configuração

### 1. Criar o bot no Discord

1. No Developer Portal, crie uma aplicação e abra a aba **Bot**.
2. Gere o token com **Reset Token** e copie-o (ele só é exibido uma vez).
3. Mantenha a opção **Public Bot** desativada.
4. Não é necessário ativar nenhum *privileged intent*.
5. Em **OAuth2 → URL Generator**, selecione os escopos `bot` e `applications.commands`, gere o link e convide o bot para o seu servidor.

### 2. Configurar as variáveis de ambiente

Copie o modelo e preencha:

```bash
cp .env.example .env
```

| Variável | Descrição |
|---|---|
| `BOT_DISCORD_TOKEN` | Token do bot, obtido no Developer Portal |
| `GUILD_ID` (Em analise...)| ID do servidor Discord onde os comandos serão registrados |

Para copiar o ID do servidor, ative o **Modo Desenvolvedor** no Discord (Configurações → Avançado), clique com o botão direito no ícone do servidor e escolha **Copiar ID do servidor**.

## Executando em desenvolvimento

Execute sempre a partir da **raiz do projeto**, pois o `.env` é procurado no diretório de trabalho atual:

```bash
go run ./cmd/golem
```

Ao iniciar com sucesso, o terminal exibe que o bot está online e os comandos foram registrados. Use `Ctrl+C` para encerrar de forma limpa.

## Build para Linux (não testado)

O Go gera um binário único e estático:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -trimpath -ldflags="-s -w" -o dist/golem ./cmd/golem
```

Use `GOARCH=arm64` se o servidor for ARM (confira com `uname -m`: `x86_64` corresponde a `amd64`, `aarch64` a `arm64`).

### Em produção (Não testado...)

Em produção o `.env` não é necessário: as variáveis podem ser injetadas pelo ambiente. A forma recomendada é um serviço **systemd** com `EnvironmentFile=` apontando para um arquivo com permissão restrita (`chmod 600`), executado por um usuário dedicado e sem privilégios. O binário não precisa de modificações para isso.

## Comandos

| Comando | Descrição | Permissão |
|---|---|---|
| `/ping` | Verifica se o bot está respondendo | Todos |
| `/start` | Inicia o servidor de Minecraft | Admin *(planejado)* |
| `/stop` | Para o servidor de Minecraft | Admin *(planejado)* |
| `/logs` | Mostra as últimas linhas do log | Admin *(planejado)* |
| `/op <jogador>` | Concede operador a um jogador | Admin *(planejado)* |
| Em desenvolvimento... | | |

## Roadmap

1. Checagem de administrador no roteador
2. `/logs` sobre uma interface `Controller`, testável sem o Docker
3. Implementação real do `Controller` (start, stop, logs, op)
4. Deploy como serviço systemd no servidor Debian
5. *(Futuro)* Suporte a múltiplos servidores e criação de servidores pelo bot?

## Licença

MIT License — veja [LICENSE](LICENSE) para detalhes.