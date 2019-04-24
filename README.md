<h1>Ponteiros em Go</h1>
<p>Cada variavel ocupa um espaço na memória, e esta possui um endereço, chamado de endereço de memória.</p>
<p>Ponteiros são referências diretas para endereço de memória.</p>
<p>Em Go é utilizado o operador & para visualizar o endereço de memória de uma variável.</p>
<p> &nomeDaVariavel ===  retorna o endereço de memória da variável.</p>
<p>var ponteiro *string ===  decalaração de um ponteiro</p>
<p>ponteiro = &nomeDaVariavel === armazenando o endereço que sera referenciado pelo ponteiro</p>
<p> *ponteiro === retorna o valor do endereço de memória que o ponteiro faz referência</p>