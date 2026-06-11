package main

import (
	"fmt"
)

type Imprimivel interface {
	ExibirDetalhes()
}

type Pessoa struct {
	Nome string
	CPF  string
}

func (p *Pessoa) ExibirDetalhes() {
	fmt.Printf("[PESSOA] Nome: %s | CPF: %s\n", p.Nome, p.CPF)
}

type Cliente struct {
	Pessoa
	CNH          string
	DataCadastro string
}

func NovoCliente(nome, cpf, cnh, dataCadastro string) *Cliente {
	return &Cliente{
		Pessoa:       Pessoa{Nome: nome, CPF: cpf},
		CNH:          cnh,
		DataCadastro: dataCadastro,
	}
}

func (c *Cliente) ExibirDetalhes() {
	fmt.Printf("[CLIENTE] Nome: %s | CPF: %s | CNH: %s | Cadastrado em: %s\n", c.Nome, c.CPF, c.CNH, c.DataCadastro)
}

type Veiculo struct {
	Placa       string
	Modelo      string
	ValorDiaria float64
	Disponivel  bool
}

type Carro struct {
	Veiculo
	Categoria  string
	PortaMalas int
}

func NovoCarro(placa, modelo string, valorDiaria float64, categoria string, portaMalas int) *Carro {
	return &Carro{
		Veiculo:    Veiculo{Placa: placa, Modelo: modelo, ValorDiaria: valorDiaria, Disponivel: true},
		Categoria:  categoria,
		PortaMalas: portaMalas,
	}
}

type Locacao struct {
	ID        int
	Cliente   *Cliente
	Carro     *Carro
	Dias      int
	ValorPago float64
	Ativa     bool
}

func NovaLocacao(id int, cliente *Cliente, carro *Carro, dias int) *Locacao {
	return &Locacao{
		ID:      id,
		Cliente: cliente,
		Carro:   carro,
		Dias:    dias,
		Ativa:   true,
	}
}

func (l *Locacao) CalcularValorTotal() float64 {
	totalBruto := float64(l.Dias) * l.Carro.ValorDiaria
	
	if l.Dias > 7 {
		desconto := totalBruto * 0.10
		fmt.Printf("[REGRA 1 - DESCONTO] %d dias qualifica para 10%% OFF! Economia de: R$ %.2f\n", l.Dias, desconto)
		return totalBruto - desconto
	}
	return totalBruto
}

func (l *Locacao) RealizarDevolucao(diasAtraso int) {
	fmt.Println("\n--- Executando Regra de Negócio: Devolução ---")
	if !l.Ativa {
		fmt.Println("Erro: Esta locação já foi encerrada anteriormente.")
		return
	}

	valorFinal := l.CalcularValorTotal()

	if diasAtraso > 0 {
		multa := float64(diasAtraso) * (l.Carro.ValorDiaria * 2)
		valorFinal += multa
		fmt.Printf("[REGRA 2 - MULTA] Atraso de %d dias detectado. Multa aplicada: R$ %.2f\n", diasAtraso, multa)
	}

	l.ValorPago = valorFinal
	l.Ativa = false
	l.Carro.Disponivel = true

	fmt.Printf(">> Devolução da Locação #%d concluída!\n", l.ID)
	fmt.Printf(">> Total Final Pago: R$ %.2f\n", l.ValorPago)
}

func main() {
	fmt.Println("=====================================================")
	fmt.Println("   SISTEMA DE LOCADORA UPF - TRABALHO PRÁTICO POO    ")
	fmt.Println("=====================================================")
	fmt.Println()

	cliente1 := NovoCliente("Eduardo Silva", "111.222.333-44", "123456789", "11/06/2026")
	carro1 := NovoCarro("UPF-2026", "Volkswagen Polo", 150.00, "Hatch", 300)

	fmt.Println("--- [TESTE 1] Polimorfismo & Sobrescrita de Métodos ---")
	
	var listaExibicao []Imprimivel
	
	pessoaComum := &Pessoa{Nome: "Carlos Souza", CPF: "000.111.222-33"}
	listaExibicao = append(listaExibicao, pessoaComum)
	listaExibicao = append(listaExibicao, cliente1)

	for _, item := range listaExibicao {
		item.ExibirDetalhes()
	}
	fmt.Println()

	fmt.Println("--- [TESTE 2] Fluxo de Locação e Regras de Negócio ---")
	if carro1.Disponivel {
		locacao1 := NovaLocacao(101, cliente1, carro1, 10) 
		carro1.Disponivel = false 
		
		fmt.Printf("Locação #%d iniciada para o cliente: %s\n", locacao1.ID, locacao1.Cliente.Nome)
		fmt.Printf("Veículo Alugado: %s | Valor da Diária: R$ %.2f\n", locacao1.Carro.Modelo, locacao1.Carro.ValorDiaria)
		
		locacao1.RealizarDevolucao(2)
	}

	fmt.Println("\n--- [TESTE 3] Validação do Estado dos Objetos em Memória ---")
	fmt.Printf("Verificando se o veículo %s voltou a ficar disponível: %t\n", carro1.Modelo, carro1.Disponivel)
	if carro1.Disponivel {
		fmt.Println("Sucesso: O objeto Carro foi liberado corretamente em memória pela classe Locacao!")
	}
	fmt.Println("=====================================================")
}