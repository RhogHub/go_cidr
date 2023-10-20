# Calculadora CIDR em Go

Esta é uma simples calculadora CIDR escrita em Go para fornecer informações sobre redes, como o endereço de rede, endereço de broadcast, número de hosts, etc.

## Como Executar o Projeto

1. **Pré-requisitos:**
   - Certifique-se de ter o Go instalado na sua máquina. Você pode baixá-lo em [golang.org](https://golang.org/dl/).

2. **Clone o Repositório:**
```bash
git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio
```

Execute o Projeto:
```bash
go run main.go <CIDR>
```
Substitua <CIDR> pelo endereço CIDR que você deseja calcular.

Exemplo de Uso:
```bash
go run main.go 192.168.1.0/24
```

Saída Esperada:
CIDR: 192.168.1.0/24
Network Address: 192.168.1.0
Broadcast Address: 192.168.1.255
Number of Hosts: 254
Subnet Mask: 255.255.255.0



