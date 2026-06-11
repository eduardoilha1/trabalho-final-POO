# trabalho-final-POO

classDiagram
    class Imprimivel {
        <<interface>>
        +ExibirDetalhes()
    }

    class Pessoa {
        +String Nome
        +String CPF
        +ExibirDetalhes()
    }

    class Cliente {
        +String CNH
        +String DataCadastro
        +ExibirDetalhes()
    }

    class Veiculo {
        +String Placa
        +String Modelo
        +Float ValorDiaria
        +Boolean Disponivel
    }

    class Carro {
        +String Categoria
        +Int PortaMalas
    }

    class Locacao {
        +Int ID
        +Int Dias
        +Float ValorPago
        +Boolean Ativa
        +CalcularValorTotal() Float
        +RealizarDevolucao(Int diasAtraso)
    }

    %% Relacionamentos de Herança (Composição no Go)
    Pessoa <|-- Cliente : Herança (Embedding)
    Veiculo <|-- Carro : Herança (Embedding)

    %% Relacionamentos de Associação
    Locacao "1" --> "1" Cliente : Associa
    Locacao "1" --> "1" Carro : Associa
    
    %% Realização de Interface
    Imprimivel <|.. Pessoa
    Imprimivel <|.. Cliente
