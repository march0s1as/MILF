<p align="center">
  <img src="https://i.imgur.com/Bcu4vGc.jpg" width="250" height="250">
</p>

    >> um script com o intuito de burlar restrições de acesso não autorizado (403 forbidden).
---
```
> para utilizar o módulo de injeção de payloads na URL:

  >> ./403.go -u http://pwnme.com.br/MILF -w url.txt
  
> para utilizar o módulo de injeção de payloads nos headers:

  >> ./403.go -u http://pwnme.com.br/ -w wordlist_headers.txt
```

```
> exemplificação de uso:
  >> ./403.go -u http://127.0.0.1:8000/exemplo/MILF -w url.txt -h "Set-Cookie: Teste" -t 2
    >>> os payloads serão injetados no local do 'MILF', terão 2 de thread e o cookie 'Set-Cookie: Teste'
```
</p>
<p align="center">
  <img src="https://user-images.githubusercontent.com/44043159/155319639-149c303f-c423-4473-9662-2a672c2c3cb2.png">
</p>

```
> exemplificação de uso:
  >> ./403.go -u http://127.0.0.1:8000/ -w wordlist.txt -h "Cookie: Exemplo" -i 192.168.1.5
    >>> os payloads serão injetados nos headers, o cookie 'Cookie: Exemplo' e o IP 192.168.1.5 no payload.
```

<p align="center">
  <img src="https://user-images.githubusercontent.com/44043159/155320417-dcfd2bd9-a50b-4d54-b915-7a0f6d8fd589.png">
</p>

---

<p align="center">
  <img src="https://user-images.githubusercontent.com/44043159/155320727-d97e6756-ec60-4373-9ec9-7b0aa3287b0c.png">
</p>
