# Pacote OS

## O que é package os?

Permite executar e usar recursos do sistema operacional

### Setenv(key, value string)

Modifica o valor de uma variável de ambiente, recebendo o nome e o valor a ser atribuído.

Retorna um erro, caso encontre algum incoveniente.

```
err := os.Setenv("Name", "Gopher")
```

### Getenv(key string)

Permite acessar as variáveis de ambiente do sistema.

Passando a variável a ser acessada, será retornado seu valor.

```
value := os.Getenv("NAME")
```

### LookupEnv(key string)

Semelhante a função Getenv(), com a diferença de que retorna 2 valores:
- O nome da variável de ambiente
- Um valor booleano mostrando se a variável existe ou não

```
value, ok := os.LookupEnv("NAME")
```

### ReadDir(name string) ([]DirEntry, error)

Lê o diretório nomeado, retornando as entradas de diretório ordenadas pelo nome do arquivo.

Caso haja erro ao ler o diretório, será retornado as entradas que foram lidas antes do erro, junto com o erro.

```
files, err := os.ReadDir(".")
```

### ReadFile(filename string)

Recebe como parâmetro o endereço e o nome do arquivo em formato texto, e retorna o conteúdo do arquivo em bytes ou um erro se houver.

```
data, err := os.ReadFile("./myFile.txt")
```

### WriteFile(filename string, data []byte, perm FileMode)
