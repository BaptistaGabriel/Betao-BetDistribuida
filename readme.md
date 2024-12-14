# Relatório de Desenvolvimento do Projeto: Betão - Bet Distribuída

---

## Índice

1. [Introdução](#introdução)
2. [Fundamentação Teórica](#fundamentação-teórica)
    1. [Apostas Online e Centralização](#apostas-online-e-centralização)
    2. [Tecnologia de Ledger Distribuído (DLT)](#tecnologia-de-ledger-distribuído-dlt)
    3. [Modelagem Estatística de Odds](#modelagem-estatística-de-odds)
3. [Metodologia](#metodologia)
    1. [Desenvolvimento do Sistema](#desenvolvimento-do-sistema)
    2. [Tecnologias Utilizadas](#tecnologias-utilizadas)
4. [Resultados](#resultados)
    1. [Contas](#contas)
    2. [Eventos](#eventos)
    3. [Apostas](#apostas)
    4. [Simulação](#simulação)
    5. [Odds](#odds)
    6. [Contabilidade](#contabilidade)
    7. [Publicação](#publicação)
5. [Conclusão](#conclusão)

---

## Introdução

O mercado de apostas online tem demonstrado um crescimento exponencial nos últimos anos, impulsionado pela acessibilidade proporcionada pela internet e smartphones. No entanto, as plataformas de apostas tradicionais são frequentemente centralizadas, o que gera preocupações em relação à transparência, segurança e regulação. Este relatório descreve o desenvolvimento de um sistema de apostas online descentralizado, baseado em tecnologia de *ledger* distribuído, com o objetivo de eliminar intermediários, garantir a integridade e confiabilidade das operações, e reduzir o risco de intervenções governamentais. O sistema proposto permite que os usuários cadastrem eventos, realizem apostas e consultem resultados de maneira transparente e segura.

## Fundamentação Teórica

### 1. Apostas Online e Centralização

As plataformas de apostas convencionais operam em um modelo centralizado, onde uma entidade única controla todos os aspectos do sistema, incluindo a definição das *odds* e a distribuição dos lucros. Esse modelo, apesar de eficiente, é vulnerável a manipulações, ataques cibernéticos e restrições legais. A centralização também contribui para a opacidade do sistema, gerando desconfiança entre os usuários.

### 2. Tecnologia de Ledger Distribuído (DLT)

O *Distributed Ledger Technology* (DLT) é uma tecnologia que permite o registro de transações em um banco de dados compartilhado, no qual cada nó da rede possui uma cópia idêntica do *ledger*. No contexto do sistema proposto, o DLT elimina a necessidade de um intermediário centralizado, aumentando a transparência e a segurança do processo. Entre os principais benefícios do DLT, destacam-se:

- **Imutabilidade**: Uma vez confirmados, os registros não podem ser alterados.
- **Transparência**: Todos os eventos e transações são visíveis para os participantes da rede.
- **Resiliência**: A ausência de um ponto central reduz a vulnerabilidade a ataques cibernéticos.

### 3. Modelagem Estatística de Odds

As *odds* funcionam de maneira similar a um bolão, onde os apostadores que acertam o resultado de um evento recebem uma fração do montante total acumulado nas apostas. Essa fração é diretamente proporcional ao valor apostado, garantindo uma distribuição justa baseada no investimento individual de cada participante.

## Metodologia

### 1. Desenvolvimento do Sistema

- **Arquitetura Descentralizada**: O sistema utiliza um modelo *peer-to-peer*, no qual todos os dados e transações são registrados em um *ledger* distribuído.
- **Cadastro de Eventos**: Os usuários podem cadastrar eventos e suas respectivas probabilidades iniciais por meio de uma interface intuitiva.
- **Apostas em Tempo Real**: A implementação de contratos inteligentes permite o gerenciamento das apostas e a distribuição dos prêmios de forma automatizada.
- **Gerenciamento de Saldo**: O sistema realiza depósitos, saques e verifica o saldo antes de confirmar uma aposta.
- **Transparência nos Resultados**: Todos os resultados são armazenados no *ledger* e podem ser auditados por qualquer participante da rede.

### 2. Tecnologias Utilizadas

- **Linguagem de Programação**: Go foi escolhida pela sua eficiência e suporte à programação concorrente.
- **DLT**: A implementação foi feita utilizando o Geth na versão 1.13.15-stable-c5ba367e, configurado com o mecanismo de consenso Proof of Authority (PoA).
- **Configuração Inicial**: O arquivo *genesis.json* define os parâmetros iniciais da blockchain, incluindo *chainId*, dificuldade, *gasLimit* e configuração do PoA.

## Resultados

### 1. Contas

O sistema mantém as contas dos usuários de forma segura, registrando todas as transações no *ledger* distribuído. Isso garante a transparência nas operações, permitindo que os usuários realizem depósitos e saques diretamente por meio de contratos inteligentes. O saldo é automaticamente atualizado após cada transação, oferecendo uma interface intuitiva.

### 2. Eventos

A funcionalidade de cadastro de eventos permite que os administradores criem e listem eventos na blockchain. Isso assegura a integridade e imutabilidade dos dados, com a possibilidade de editar informações antes que os eventos sejam confirmados como ativos.

### 3. Apostas

Os usuários podem realizar apostas em eventos cadastrados, desde que possuam saldo suficiente. As *odds* são calculadas dinamicamente, de acordo com um modelo semelhante a um bolão, onde a fração do prêmio é proporcional ao valor apostado. As transações são registradas de forma transparente no *ledger*, garantindo uma auditoria total do processo.

### 4. Simulação

O sistema permite a simulação de eventos em tempo real, com a execução de algoritmos que processam os resultados de acordo com os critérios definidos no contrato inteligente.

### 5. Odds

As *odds* são atualizadas dinamicamente com base na proporção de apostas realizadas para cada possível resultado. Esse modelo assegura que os prêmios sejam distribuídos de forma justa e proporcional ao valor apostado, eliminando a necessidade de uma "banca" centralizada.

### 6. Contabilidade

Após cada evento, o sistema calcula automaticamente os resultados e distribui os prêmios aos vencedores. Os saldos dos usuários são atualizados de forma precisa, garantindo a consistência entre as transações registradas no *ledger*.

### 7. Publicação

Os resultados dos eventos são armazenados na blockchain e estão acessíveis para consulta pública. A transparência dos resultados permite que qualquer participante audite os eventos e as transações associadas.

## Conclusão

O sistema desenvolvido demonstra a viabilidade de criar uma plataforma de apostas online descentralizada, segura e transparente. O processo de desenvolvimento não foi isento de desafios, como a configuração inicial da blockchain e ajustes no arquivo *genesis.json* para otimizar o desempenho e reduzir os atrasos na sincronização. Além disso, a implementação dos contratos inteligentes exigiu cuidados adicionais para garantir a imutabilidade dos dados e a segurança do sistema. No entanto, os resultados confirmam os benefícios do modelo proposto, que elimina intermediários e garante a integridade das operações. A utilização de tecnologias como DLT e contratos inteligentes assegura a confiabilidade do sistema. O modelo pode ser expandido para incluir mais funcionalidades e explorar novos casos de uso, além de oferecer uma solução robusta para lidar com regulação e restrições no mercado de apostas online.

---

**Autores:**
[Amanda Lima](https://github.com/amandalima)
[Gabriel Baptista](https://github.com/gabrielbaptista)