# Normalização de respostas — demonstração experimental offline

**Estágio: experimental / em teste. Demonstração parcial sanitizada, não pronta para produção.**

## Objetivo
Um recorte de normalização de respostas aninhadas usado em consulta de auditoria.

## Recursos
Localiza listas de eventos/pedidos e metadados de paginação em objetos JSON.

## Instalação
Go 1.18+; apenas biblioteca padrão. Baixe somente os arquivos desta demonstração para uma pasta local.

## Uso e testes
**Status experimental: esta versão não foi compilada nem executada na preparação para publicação. Os quatro testes incluídos não foram executados, e o funcionamento não foi verificado.** Não havia um compilador Go válido no ambiente de preparação. Os comandos abaixo são instruções para validação futura, não evidência de testes aprovados. Pode haver erros de compilação ou execução; não use este recorte em produção.

```sh
go run demo.go normalizer.go
go test demo.go normalizer.go normalizer_test.go
```

## Estágio e limites
Demonstração parcial de lógica extraída e adaptada, não uma cópia integral do sistema original. Fixtures e entradas são inteiramente fictícias.

Não é um auditor financeiro completo. Não calcula divergências nem prova conciliação. O exemplo demonstra somente extração de listas/paginação. Nenhuma conexão, sessão, servidor ou credencial é utilizada; não executa requisições.

## Dados e segurança
Não inclui configurações operacionais, contas, histórico de conversas, bancos de dados, capturas de rede, segredos ou endpoints de produção. Nenhum envio externo ocorre.

## Licença
Nenhuma licença de redistribuição é atribuída automaticamente. A definição de licença depende de revisão dos direitos aplicáveis antes da publicação.
