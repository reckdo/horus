## horus - scanner tcp/ip em golang.

<p align="justify">Um das ferramentas interessantes a ser utilizada em manutenção e criações de servidores é o scanner map. Esse tipo de ferramenta nos possibilitam ter informações que são necessárias, como por exemplo: as portas que estão abertas em nossos servidores de forma interna ou externa.</p> 

<p align="justify">Uma ferramenta que é bem conhecido que também faz isso é o NMAP, que é um software que realiza validações de portas em redes e servidores. Com tudo, o NMAP é uma ferramenta digamos que bem mais complexa, porém iremos fazer uma versão que se aproxima, mas de uma forma mais simples e em golang.</p> 

<p align="justify">Antes de começarmos a criar essa ferramenta, precisamos entender alguns conceitos que são bem importantes sobre alguns conceitos de protocolo de tcp/ip. O protocolo de comunicação TCP (Protocolo de Controle de Tramissão) tem a função de ser a camada de transporte em rede de computadores no modelo OSI. Sendo conhecido também como protocolo TCP/IP, que é um complemento do protocolo de internet (IP). Essa junção dos TCP+IP traz a possibilidade dos aplicativos se comunicarem entre si.</p> 

- O cliente envia um pacote chamado SYN para sincronização que contém um número sequencial gerado aleatoriamente;
- O servidor envia um pacote SYN-ACK. (O ACK significa reconhecimento do SYN, que envia o numero recebido e complementa com um novo numero para retorno);
- O cliente finaliza com um ACK informando que foi fechado a comunicação entre as partes. E essa comunicação é chamada de full-duplex.

<p align="justify">Pensando em um cenário onde o cliente gostaria de se conectar em um servidor na porta 80, caso essa a porta esteja fechada para comunicação TCP o cliente envia o seu sinalizador SYN, e o servidor responder um RST em vez de SYNK-ACK, informando que essa comunicação não teve sucesso.</p> 

<p align="justify">O RST é um tipo de retorno onde o servidor informa para o cliente que a conexão para ação em porta não poderá ser concluída, e com isso retorna o reset da comunicação. Esse tipo de retorno pode ser avaliado por meio de bloqueios em Firewalls ou porta fechada. Existem outros tipos de varreduras que podem ser utilizadas para verificação, tais como: XMAS, FIN, TCP ACK e etc. Para esses tipos de scanners podemos utilizar pacotes Gopacket.</p> 

