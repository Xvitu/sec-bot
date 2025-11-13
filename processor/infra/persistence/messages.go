package persistence

import "github.com/xvitu/sec-bot/processor/domain"

var Messages = map[domain.Step]map[string]string{
	domain.MainMenu: {
		"greetings": "*Oi! Sou o SecBot 🤖*\nComo posso te ajudar hoje?\n\n*1* - Dúvidas\n*2* - Quiz\n*3* - Dicas\n*4* - Infos sobre golpes",
	},

	domain.Error: {
		"invalid_option": "❌ Não entendi, pode repetir?",
	},

	domain.Faq: {
		"faq_menu": "*Aqui está uma lista de assuntos sobre os quais posso tirar dúvidas:*\n\n" +
			"1️⃣ O que é phishing?\n" +
			"2️⃣ Para que serve um firewall?\n" +
			"3️⃣ Diferença entre HTTP e HTTPS\n" +
			"4️⃣ O que é um malware?\n" +
			"5️⃣ O que é autenticação de dois fatores (2FA)?\n" +
			"6️⃣ Criptografia simétrica vs assimétrica\n" +
			"7️⃣ O que é um ataque DDoS?\n" +
			"8️⃣ Para que serve uma VPN?\n" +
			"9️⃣ O que é SQL Injection e como prevenir?\n" +
			"🔟 O que é engenharia social?\n\n" +
			"⬅️ 11 - Voltar",
		"faq_1":  "É um ataque de engenharia social usado para roubar dados, como senhas e informações bancárias.",
		"faq_2":  "Protege a rede filtrando tráfego e bloqueando acessos não autorizados.",
		"faq_3":  "HTTPS é a versão segura do HTTP, usando criptografia TLS/SSL.",
		"faq_4":  "Software malicioso criado para danificar ou explorar sistemas.",
		"faq_5":  "Método que exige duas verificações para acessar uma conta.",
		"faq_6":  "Simétrica usa a *mesma chave*; assimétrica usa *chave pública + privada*.",
		"faq_7":  "Ataque distribuído que sobrecarrega serviços, tornando-os indisponíveis.",
		"faq_8":  "Garante uma conexão segura e criptografada entre usuário e rede.",
		"faq_9":  "Ataque que explora falhas em consultas SQL; previne-se com validação e queries parametrizadas.",
		"faq_10": "Manipulação psicológica usada para obter informações confidenciais.",
	},

	domain.Tips: {
		"tip_1":    "🔐 Crie senhas longas e complexas. Evite repetir senhas.",
		"tip_2":    "📲 Ative 2FA sempre que possível.",
		"tip_3":    "🔄 Mantenha sistemas e apps atualizados.",
		"tip_5":    "⚠️ Não clique em links suspeitos.",
		"tip_6":    "🌐 Redes públicas são arriscadas — use VPN.",
		"tip_8":    "💾 Faça backups periódicos.",
		"tip_9":    "🔒 Verifique se a URL usa HTTPS.",
		"tip_10":   "📎 Não abra anexos desconhecidos.",
		"tip_11":   "🛡️ Use antivírus atualizado.",
		"tip_12":   "🙅 Evite divulgar dados pessoais sensíveis.",
		"tip_menu": "*1* - Dica aleatória\n*2* - Voltar",
	},

	domain.Scams: {
		"scam_menu": "*Aqui estão alguns golpes comuns:*\n\n" +
			"1️⃣ Phishing (e-mail/mensagem falsa)\n" +
			"2️⃣ Voice phishing (telefone)\n" +
			"3️⃣ Smishing (SMS)\n" +
			"4️⃣ Suporte técnico falso\n" +
			"5️⃣ Boleto/cobrança falsa\n" +
			"6️⃣ Golpe do código do WhatsApp\n" +
			"7️⃣ Comprovante falso / PIX falso\n" +
			"8️⃣ Loja ou anúncio falso\n" +
			"9️⃣ Prêmio falso\n" +
			"🔟 Romance scam\n\n" +
			"⬅️ 11 - Voltar",
		"scam_1":  "Engana a vítima para obter senhas ou dados financeiros.",
		"scam_2":  "Golpista liga fingindo ser empresa/autoridade para roubar dados.",
		"scam_3":  "SMS fraudulento com links maliciosos.",
		"scam_4":  "Golpista finge ser suporte técnico para instalar malware.",
		"scam_5":  "Boletos falsos enviados para pagamento indevido.",
		"scam_6":  "Criminoso tenta obter o código do WhatsApp.",
		"scam_7":  "Comprovantes falsos ou mensagens pedindo PIX.",
		"scam_8":  "Loja falsa que cobra mas não entrega.",
		"scam_9":  "Promessa falsa de prêmio para obter dados ou dinheiro.",
		"scam_10": "Criminosos criam perfis falsos para enganar emocionalmente e pedir dinheiro.",
	},

	domain.Quiz: {
		"quiz_menu": "*Escolha uma opção:*\n1️⃣ Enviar pergunta\n2️⃣ Voltar",
	},

	domain.QuizQuestion: {
		"quiz_1":  "*O que é phishing?*\n1. Malware\n2. Engenharia social\n3. Criptografia",
		"quiz_2":  "*Para que serve um firewall?*\n1. Protege redes\n2. Acelera conexão\n3. Cria senhas",
		"quiz_3":  "*HTTP vs HTTPS*?\n1. HTTPS é seguro\n2. HTTP é seguro\n3. Iguais",
		"quiz_4":  "*O que é malware?*\n1. Software malicioso\n2. Atualização\n3. Firewall",
		"quiz_5":  "*O que é 2FA?*\n1. Senha longa\n2. Duas verificações\n3. VPN",
		"quiz_6":  "*Criptografia simétrica vs assimétrica*\n1. Só pública\n2. Só privada\n3. Mesma chave vs par de chaves",
		"quiz_7":  "*O que é DDoS?*\n1. Sobrecarrega\n2. Criptografa\n3. Instala malware",
		"quiz_8":  "*Para que serve VPN?*\n1. Atualização\n2. Conexão segura\n3. Velocidade",
		"quiz_9":  "*O que é SQL Injection?*\n1. Firewall\n2. Criptografia\n3. Ataque a consultas SQL",
		"quiz_10": "*Engenharia social é?*\n1. Manipulação\n2. Criptografia\n3. Malware",
		"quiz_11": "*O que é hash?*\n1. Firewall\n2. Resumo de dados\n3. Senha",
		"quiz_12": "*Zero-day?*\n1. Atualizações\n2. Ataques antigos\n3. Falhas sem correção",
		"quiz_13": "*XSS vs CSRF*\n1. Scripts vs ações forçadas\n2. Ambos SQL\n3. Malware",
		"quiz_14": "*O que é certificado digital?*\n1. Senha\n2. Identidade online\n3. VPN",
		"quiz_15": "*Defense in Depth?*\n1. Criptografia\n2. Um firewall\n3. Camadas de segurança",
		"quiz_16": "*Vishing?*\n1. Golpe por telefone\n2. Email\n3. Cartão",
		"quiz_17": "*Smishing?*\n1. Criptografia\n2. SMS malicioso\n3. Firewall móvel",
		"quiz_18": "*Golpe do falso suporte técnico?*\n1. Proteção\n2. Antivírus\n3. Engana usuário para instalar malware",
		"quiz_19": "*Boleto falso?*\n1. Fraudulento\n2. Legítimo\n3. VPN",
		"quiz_20": "*Prêmio falso?*\n1. Real\n2. Promessa falsa\n3. Update",
	},

	domain.QuizAnswer: {
		"quiz_1":  "2",
		"quiz_2":  "1",
		"quiz_3":  "1",
		"quiz_4":  "1",
		"quiz_5":  "2",
		"quiz_6":  "3",
		"quiz_7":  "1",
		"quiz_8":  "2",
		"quiz_9":  "3",
		"quiz_10": "1",
		"quiz_11": "2",
		"quiz_12": "3",
		"quiz_13": "1",
		"quiz_14": "2",
		"quiz_15": "3",
		"quiz_16": "1",
		"quiz_17": "2",
		"quiz_18": "3",
		"quiz_19": "1",
		"quiz_20": "2",
	},

	domain.QuizFeedback: {
		"quiz_error":   "❌ *Resposta incorreta!*\nAqui vai uma breve explicação:",
		"quiz_success": "✅ *Certa resposta! Muito bem!*",
	},

	domain.QuizExplanation: {
		"quiz_1":  "Phishing é um ataque de engenharia social usado para roubar dados sensíveis.",
		"quiz_2":  "Um firewall protege a rede filtrando tráfego malicioso.",
		"quiz_3":  "HTTPS usa criptografia para garantir segurança na comunicação.",
		"quiz_4":  "Malware é software criado para prejudicar sistemas.",
		"quiz_5":  "2FA usa duas verificações para aumentar segurança.",
		"quiz_6":  "Simétrica usa uma chave; assimétrica usa chave pública e privada.",
		"quiz_7":  "DDoS sobrecarrega serviços com tráfego excessivo.",
		"quiz_8":  "VPN cria conexão segura e criptografada.",
		"quiz_9":  "SQL Injection insere comandos maliciosos em consultas.",
		"quiz_10": "Engenharia social manipula pessoas para obter dados.",
		"quiz_11": "Hash é um resumo único usado para verificação.",
		"quiz_12": "Zero-day é uma falha sem correção disponível.",
		"quiz_13": "XSS injeta scripts; CSRF força ações indesejadas.",
		"quiz_14": "Certificado digital valida identidades online.",
		"quiz_15": "Defense in Depth usa múltiplas camadas de segurança.",
		"quiz_16": "Vishing é golpe por telefone.",
		"quiz_17": "Smishing é golpe via SMS.",
		"quiz_18": "Falso suporte técnico instala malware ou obtém acesso.",
		"quiz_19": "Boleto falso engana a vítima para pagar valores indevidos.",
		"quiz_20": "Prêmio falso promete recompensas inexistentes.",
	},
}
