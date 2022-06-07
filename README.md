#       MILF (man i laugh forbidden)
<p align="center">
  <img src="https://camo.githubusercontent.com/13160423f50201640ce87be60be8327e5da49070f6ffab2df422b30d323f47ac/68747470733a2f2f696d616765732d7769786d702d6564333061383662386334636138383737373335393463322e7769786d702e636f6d2f662f39383937623065632d323937372d343733362d383563662d6161343836656431323431302f64316b703667702d61396536343433372d343662352d343135622d623838652d6362393830386338363262632e6769663f746f6b656e3d65794a30655841694f694a4b563151694c434a68624763694f694a49557a49314e694a392e65794a7a645749694f694a31636d3436595842774f6a646c4d4751784f4467354f4449794e6a517a4e7a4e684e5759775a4451784e5756684d4751794e6d55774969776961584e7a496a6f6964584a754f6d467763446f335a54426b4d5467344f5467794d6a59304d7a637a5954566d4d4751304d54566c5954426b4d6a5a6c4d434973496d39696169493657317437496e4268644767694f694a634c325a634c7a6b344f5464694d47566a4c5449354e7a63744e44637a4e6930344e574e6d4c5746684e4467325a5751784d6a51784d4677765a44467263445a6e634331684f5755324e44517a4e7930304e6d49314c5451784e574974596a67345a53316a596a6b344d44686a4f445979596d4d755a326c6d496e3164585377695958566b496a7062496e5679626a707a5a584a3261574e6c4f6d5a70624755755a473933626d7876595751695858302e44384736415f6b776466696f697972374574434d5f744d33626a786e365871555a4b3278334d447132634d" width="250" height="250">
</p>

    >> um script com o intuito de burlar restrições de acesso não autorizado (403 forbidden).
---
```bash
> para utilizar o módulo de injeção de payloads na URL:

  >> ./403 -u http://pwnme.com.br/MILF -w wordlist.txt
  
> para utilizar o módulo de injeção de payloads nos headers:

  >> ./403 -u http://pwnme.com.br/ -w wordlist_headers.txt
```

```bash
> exemplificação de uso:
  >> ./403 -u http://127.0.0.1:8000/exemplo/MILF -w wordlist.txt -h "Set-Cookie: Teste" -t 2
    >> os payloads serão injetados na posição no 'MILF', terão 2 de thread e o cookie 'Set-Cookie: Teste'
```
</p>
<p align="center">
  <img src="https://user-images.githubusercontent.com/44043159/155319639-149c303f-c423-4473-9662-2a672c2c3cb2.png">
</p>

```bash
> exemplificação de uso:
  >> ./403 -u http://127.0.0.1:8000/ -w wordlist.txt -h "Cookie: Exemplo" -i 192.168.1.5
    >> os payloads serão injetados nos headers, o cookie 'Cookie: Exemplo' e o IP 192.168.1.5 no payload.
```

<p align="center">
  <img src="https://user-images.githubusercontent.com/44043159/155320417-dcfd2bd9-a50b-4d54-b915-7a0f6d8fd589.png">
</p>

---

<p align="center">
  <img src="https://user-images.githubusercontent.com/44043159/155320727-d97e6756-ec60-4373-9ec9-7b0aa3287b0c.png">
</p>

```bash
> instalação
  :: git clone https://github.com/march0s1as/MILF
  :: cd MILF
  :: go get -v github.com/fatih/color
  :: go build 403.go
  :: ./403 -h
```

