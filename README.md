# Sobre o CODEPIX

## Descrição

- É uma solução para simular transferências de valores entre bancos fictícios através de chaves (e-mail, CPF).
- Simularemos diversos bancos e contas bancárias que possuem uma chave PIX atribuída.
- Cada conta bancária poderá cadastrar suas chaves Pix.
- Uma conta bancária modera realizar uma transferência para outra conta em outro banco utilizando a chave Pix da conta de destino.
- Uma transação não pode ser perdida mesmo que: o CodePix esteja fora do ar.
- Uma transação não pode ser perdida mesmo que: o Banco de destino esteja fora do ar.

# Sobre os Bancos

- O banco será um microsserviço com funções limitadas a cadastro de contas e chave Pix, bem como transferências de valores.
- Utilizaremos a mesma aplicação para simular diversos bancos, mudando apenas as cores, nome e código.
- Nest.js no Backend.
- Next.js no frontend.

# Sobre o Microserviço CODEPIX

- Será responsável por intermediar as transferências bancárias.
- Receberá a transação de transferência.
- Encaminhará a transação para o banco de destino (Status: PENDING).
- Receberá a confirmação do banco de destino (Status: CONFIRMED).
- Envia confirmação para o banco de origem informando quando o banco de destino processou.
- Recebe a confirmação do banco de origem de que ele processou (Status: COMPLETED ).
- Marca a transação como completa (Status: COMPLETED).

## Cadastro e consulta de chaves Pix


